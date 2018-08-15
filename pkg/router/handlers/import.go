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

func ImportNamespaceListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	perm := ctx.MustGet(m.PermClient).(clients.Permissions)

	if err := importNamespaceList(ctx, kube, perm); err != nil {
		ctx.Error(err)
	} else {
		ctx.Status(http.StatusAccepted)
	}
}

func importNamespaceList(ctx *gin.Context, kube *kubernetes.Kube, perm clients.Permissions) error {
	quotas, err := kube.GetNamespaceQuotaList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	nsList, err := model.ParseKubeResourceQuotaList(quotas)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	nss, err := kube.GetNamespaceList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	nsList = importerModel.AddNamespacesWithoutRQ(nsList, nss)

	if err := perm.ImportNamespaces(ctx, *nsList); err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return err
	}

	return nil
}

func ImportDeploymentsListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)
	if err := importDeploymentsList(ctx, kube, res); err != nil {
		ctx.Error(err)
	} else {
		ctx.Status(http.StatusAccepted)
	}
}

func importDeploymentsList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) error {
	depl, err := kube.GetDeploymentList("", "")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	ret, err := model.ParseKubeDeploymentList(depl, false)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	if err := res.ImportDeployments(ctx, *ret); err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return err
	}

	return nil
}

func ImportServicesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	if err := importServicesList(ctx, kube, res); err != nil {
		ctx.Error(err)
	} else {
		ctx.Status(http.StatusAccepted)
	}
}

func importServicesList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) error {

	quotas, err := kube.GetServiceList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	ret, err := model.ParseKubeServiceList(quotas, false)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	if err := res.ImportServices(ctx, *ret); err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return err
	}

	return nil
}

func ImportConfigMapsListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	if err := importConfigMapsList(ctx, kube, res); err != nil {
		ctx.Error(err)
	} else {
		ctx.Status(http.StatusAccepted)
	}
}

func importConfigMapsList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) error {
	quotas, err := kube.GetConfigMapList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	ret, err := model.ParseKubeConfigMapList(quotas, false)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	for i := range ret.ConfigMaps {
		ret.ConfigMaps[i].Data = nil
	}

	if err := res.ImportConfigMaps(ctx, *ret); err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return err
	}

	return nil
}

func ImportIngressesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	if err := importIngressesList(ctx, kube, res); err != nil {
		ctx.Error(err)
	} else {
		ctx.Status(http.StatusAccepted)
	}
}

func importIngressesList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) error {
	quotas, err := kube.GetIngressList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	ret, err := model.ParseKubeIngressList(quotas, false)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return err
	}

	if err := res.ImportIngresses(ctx, *ret); err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return err
	}

	return nil
}

func ImportAllHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	perm := ctx.MustGet(m.PermClient).(clients.Permissions)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	if err := importNamespaceList(ctx, kube, perm); err != nil {
		ctx.Error(err)
		return
	}

	if err := importDeploymentsList(ctx, kube, res); err != nil {
		ctx.Error(err)
		return
	}

	if err := importServicesList(ctx, kube, res); err != nil {
		ctx.Error(err)
		return
	}

	if err := importIngressesList(ctx, kube, res); err != nil {
		ctx.Error(err)
		return
	}

	if err := importConfigMapsList(ctx, kube, res); err != nil {
		ctx.Error(err)
		return
	}
}
