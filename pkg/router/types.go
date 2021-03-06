package router

import (
	"net/textproto"

	"git.containerum.net/ch/volume-manager/pkg/errors"
	"git.containerum.net/ch/volume-manager/pkg/router/middleware"
	"git.containerum.net/ch/volume-manager/static"
	"github.com/containerum/cherry"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/utils/httputil"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

type TranslateValidate struct {
	*ut.UniversalTranslator
	*validator.Validate
}

func (tv *TranslateValidate) HandleError(err error) (int, *cherry.Err) {
	switch err.(type) {
	case *cherry.Err:
		e := err.(*cherry.Err)
		return e.StatusHTTP, e
	default:
		return errors.ErrInternal().StatusHTTP, errors.ErrInternal().AddDetailsErr(err)
	}
}

func (tv *TranslateValidate) BadRequest(ctx *gin.Context, err error) (int, *cherry.Err) {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		ret := errors.ErrRequestValidationFailed()
		for _, fieldErr := range validationErr {
			if fieldErr == nil {
				continue
			}
			t, _ := tv.FindTranslator(httputil.GetAcceptedLanguages(ctx.Request.Context())...)
			ret.AddDetailF("Field %s: %s", fieldErr.Namespace(), fieldErr.Translate(t))
		}
		return ret.StatusHTTP, ret
	}
	return errors.ErrRequestValidationFailed().StatusHTTP, errors.ErrRequestValidationFailed().AddDetailsErr(err)
}

func (tv *TranslateValidate) ValidateHeaders(headerTagMap map[string]string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerErr := make(map[string]validator.ValidationErrors)
		for header, tag := range headerTagMap {
			ferr := tv.VarCtx(ctx.Request.Context(), ctx.GetHeader(textproto.CanonicalMIMEHeaderKey(header)), tag)
			if ferr != nil {
				headerErr[header] = ferr.(validator.ValidationErrors)
			}
		}
		if len(headerErr) > 0 {
			ret := errors.ErrRequestValidationFailed()
			for header, fieldErrs := range headerErr {
				for _, fieldErr := range fieldErrs {
					if fieldErr == nil {
						continue
					}
					t, _ := tv.FindTranslator(httputil.GetAcceptedLanguages(ctx.Request.Context())...)
					ret.AddDetailF("Header %s: %s", header, fieldErr.Translate(t))
				}
			}
			ctx.AbortWithStatusJSON(ret.StatusHTTP, ret)
			return
		}
	}
}

func (tv *TranslateValidate) ValidateURLParams(paramTagMap map[string]string) gin.HandlerFunc {
	return httputil.ValidateURLParamsMiddleware(paramTagMap, tv.Validate, tv.UniversalTranslator, errors.ErrRequestValidationFailed)
}

type Router struct {
	engine gin.IRouter
	tv     *TranslateValidate
}

func NewRouter(engine gin.IRouter, status *model.ServiceStatus, tv *TranslateValidate) *Router {
	engine.StaticFS("/static", static.HTTP)

	engine.GET("/status", httputil.ServiceStatus(status))

	ret := &Router{
		engine: engine,
		tv:     tv,
	}
	ret.engine.Use(httputil.SaveHeaders)
	ret.engine.Use(httputil.PrepareContext)
	ret.engine.Use(httputil.RequireHeaders(errors.ErrRequiredHeadersNotProvided, httputil.UserIDXHeader, httputil.UserRoleXHeader))
	ret.engine.Use(tv.ValidateHeaders(map[string]string{
		httputil.UserIDXHeader:   "uuid",
		httputil.UserRoleXHeader: "eq=admin|eq=user",
	}))
	ret.engine.Use(httputil.SubstituteUserMiddleware(tv.Validate, tv.UniversalTranslator, errors.ErrRequestValidationFailed))
	ret.engine.Use(middleware.RequiredUserHeaders())
	return ret
}
