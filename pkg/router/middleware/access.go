package middleware

import (
	volErrors "git.containerum.net/ch/volume-manager/pkg/errors"
	"github.com/containerum/cherry/adaptors/gonic"
	kubeModel "github.com/containerum/kube-client/pkg/model"
	headers "github.com/containerum/utils/httputil"
	"github.com/gin-gonic/gin"
)

var (
	readLevels = []kubeModel.AccessLevel{
		kubeModel.Owner,
		kubeModel.Write,
		kubeModel.ReadDelete,
		kubeModel.Read,
	}

	deleteLevels = []kubeModel.AccessLevel{
		kubeModel.Owner,
		kubeModel.Write,
		kubeModel.ReadDelete,
	}

	writeLevels = []kubeModel.AccessLevel{
		kubeModel.Owner,
		kubeModel.Write,
	}
)

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

func IsAdmin(ctx *gin.Context) {
	if role := GetHeader(ctx, headers.UserRoleXHeader); role != RoleAdmin {
		gonic.Gonic(volErrors.ErrAdminRequired(), ctx)
		return
	}
}

func ReadAccess(ctx *gin.Context) {
	CheckAccess(ctx, readLevels)
}

func DeleteAccess(ctx *gin.Context) {
	CheckAccess(ctx, deleteLevels)
}

func WriteAccess(ctx *gin.Context) {
	CheckAccess(ctx, writeLevels)
}

func CheckAccess(ctx *gin.Context, level []kubeModel.AccessLevel) {
	ns := ctx.Param("ns_id")
	if GetHeader(ctx, headers.UserRoleXHeader) == RoleUser {
		var userNsData *kubeModel.UserHeaderData
		nsList := ctx.MustGet(UserNamespaces).(UserHeaderDataMap)
		for _, n := range nsList {
			if ns == n.ID {
				userNsData = &n
				break
			}
		}
		if userNsData != nil {
			if ok := containsAccess(userNsData.Access, level...); ok {
				return
			}
			gonic.Gonic(volErrors.ErrRequestValidationFailed().AddDetailF("access error"), ctx)
			return
		}
		gonic.Gonic(volErrors.ErrResourceNotExists().AddDetails("namespace is not found"), ctx)
		return
	}
}

func containsAccess(access kubeModel.AccessLevel, in ...kubeModel.AccessLevel) bool {
	for _, acc := range in {
		if acc == access {
			return true
		}
	}
	return false
}
