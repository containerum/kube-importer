package handlers

import (
	"net/http"

	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"git.containerum.net/ch/kube-api/pkg/model"
	importerModel "github.com/containerum/kube-importer/pkg/model"
	m "github.com/containerum/kube-importer/pkg/router/middleware"

	"github.com/gin-gonic/gin"

	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/kube-importer/pkg/clients"
	"github.com/containerum/kube-importer/pkg/kierrors"
)

func ImportNamespaceList(ctx *gin.Context) {

	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	perm := ctx.MustGet(m.PermClient).(clients.Permissions)

	quotas, err := kube.GetNamespaceQuotaList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	nsList, err := model.ParseKubeResourceQuotaList(quotas)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	nss, err := kube.GetNamespaceList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	nsList = importerModel.AddNamespacesWithoutRQ(nsList, nss)

	if err := perm.ImportNamespaces(ctx, *nsList); err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return
	}
	ctx.Status(http.StatusAccepted)
}

func ImportDeploymentsList(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	quotas, err := kube.GetDeploymentList("", "")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	ret, err := model.ParseKubeDeploymentList(quotas, false)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	if err := res.ImportDeployments(ctx, *ret); err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return
	}

	ctx.Status(http.StatusAccepted)
}

func ImportServicesList(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	quotas, err := kube.GetServiceList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	ret, err := model.ParseKubeServiceList(quotas, false)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	if err := res.ImportServices(ctx, *ret); err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return
	}

	ctx.Status(http.StatusAccepted)
}

func ImportConfigMapsList(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	quotas, err := kube.GetConfigMapList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	ret, err := model.ParseKubeConfigMapList(quotas, false)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	for i := range ret.ConfigMaps {
		ret.ConfigMaps[i].Data = nil
	}

	if err := res.ImportConfigMaps(ctx, *ret); err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return
	}

	ctx.Status(http.StatusAccepted)
}

func ImportIngressesList(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	quotas, err := kube.GetIngressList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	ret, err := model.ParseKubeIngressList(quotas, false)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	if err := res.ImportIngresses(ctx, *ret); err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return
	}

	ctx.Status(http.StatusAccepted)
}
