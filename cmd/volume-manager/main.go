package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"text/tabwriter"
	"time"

	"git.containerum.net/ch/volume-manager/pkg/errors"
	"git.containerum.net/ch/volume-manager/pkg/router"
	"git.containerum.net/ch/volume-manager/pkg/server"
	"git.containerum.net/ch/volume-manager/pkg/utils/validation"
	"github.com/containerum/cherry/adaptors/cherrylog"
	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/utils/httputil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

//go:generate swagger generate spec -m -i ../../swagger-basic.yml -o ../../swagger.json
//go:generate swagger flatten ../../swagger.json -o ../../swagger.json
//go:generate swagger validate ../../swagger.json

func exitOnError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func errFuture(f func() error) <-chan error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()
	return ch
}

func prettyPrintFlags(ctx *cli.Context) {
	fmt.Printf("Starting %v %v\n", ctx.App.Name, ctx.App.Version)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent|tabwriter.Debug)
	for _, f := range ctx.App.VisibleFlags() {
		fmt.Fprintf(w, "Flag: %s\t Value: %v\n", f.Names()[0], ctx.Generic(f.Names()[0]))
	}
	w.Flush()
}

const httpServerContextKey = "httpsrv"

var version string

func main() {
	app := cli.App{
		Name:        "volume",
		Description: "Volume-Manager volumes management service for Container hosting",
		Version:     version,
		Flags: []cli.Flag{
			&ModeFlag,
			&LogLevelFlag,
			&DBUserFlag,
			&DBPassFlag,
			&DBHostFlag,
			&DBBaseFlag,
			&DBSSLModeFlag,
			&ListenAddrFlag,
			&BillingAddrFlag,
			&KubeAPIAddrFlag,
			&CORSFlag,
		},
		Before: func(ctx *cli.Context) error {
			prettyPrintFlags(ctx)

			if err := setupLogger(ctx); err != nil {
				return err
			}

			listenAddr := getListenAddr(ctx)

			translate := setupTranslator()
			validate := validation.StandardPermissionsValidator(translate)

			db, err := setupDB(ctx)
			if err != nil {
				return err
			}

			clients, err := setupServiceClients(ctx)
			if err != nil {
				return err
			}

			srv := server.NewServer(db, clients)

			g := gin.New()
			g.Use(gonic.Recovery(errors.ErrInternal, cherrylog.NewLogrusAdapter(logrus.WithField("component", "gin_recovery"))))
			g.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
			binding.Validator = &validation.GinValidatorV9{Validate: validate} // gin has no local validator

			if ctx.Bool(CORSFlag.Name) {
				corsCfg := cors.DefaultConfig()
				corsCfg.AllowAllOrigins = true
				corsCfg.AddAllowHeaders(
					httputil.UserIDXHeader,
					httputil.UserRoleXHeader,
				)
				g.Use(cors.New(corsCfg))
			}

			status := model.ServiceStatus{
				Name:     ctx.App.Name,
				Version:  ctx.App.Version,
				StatusOK: true,
			}

			r := router.NewRouter(g, &status, &router.TranslateValidate{UniversalTranslator: translate, Validate: validate})
			r.SetupVolumeHandlers(srv)
			r.SetupStorageHandlers(srv)

			// for graceful shutdown
			httpsrv := &http.Server{
				Addr:    listenAddr,
				Handler: g,
			}

			ctx.App.Metadata[httpServerContextKey] = httpsrv

			return nil
		},
		Action: func(ctx *cli.Context) error {
			httpsrv := ctx.App.Metadata[httpServerContextKey].(*http.Server)
			errCh := errFuture(func() error {
				return httpsrv.ListenAndServe()
			})

			// Wait for interrupt signal to gracefully shutdown the server with
			// a timeout of 5 seconds.
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt) // subscribe on interrupt event

			select {
			case err := <-errCh:
				return err
			case <-quit:
				logrus.Infoln("shutting down server...")
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				return httpsrv.Shutdown(ctx)
			}
		},
	}

	exitOnError(app.Run(os.Args))
}
