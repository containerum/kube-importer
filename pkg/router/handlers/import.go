package handlers

import (
	"net/http"

	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"git.containerum.net/ch/kube-api/pkg/model"
	kubtypes "github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"

	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/kube-importer/pkg/clients"
	"github.com/containerum/kube-importer/pkg/kierrors"
	m "github.com/containerum/kube-importer/pkg/router/middleware"
)

func ImportNamespaceListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	perm := ctx.MustGet(m.PermClient).(clients.Permissions)

	resp, err := importNamespaceList(ctx, kube, perm)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

func ImportDeploymentsListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)
	resp, err := importDeploymentsList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

func ImportServicesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	resp, err := importServicesList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

func ImportIngressesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	resp, err := importIngressesList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

func ImportConfigMapsListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	resp, err := importConfigMapsList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

func ImportStoragesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	vol := ctx.MustGet(m.VolClient).(clients.Volumes)

	resp, err := importStoragesList(ctx, kube, vol)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

func ImportVolumesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	vol := ctx.MustGet(m.VolClient).(clients.Volumes)

	resp, err := importVolumesList(ctx, kube, vol)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

func importNamespaceList(ctx *gin.Context, kube *kubernetes.Kube, perm clients.Permissions) (*kubtypes.ImportResponse, error) {
	ret, err := exportNamespaces(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := perm.ImportNamespaces(ctx, ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
}

func importDeploymentsList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	ret, err := exportDeployments(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := res.ImportDeployments(ctx, ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
}

func importServicesList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	ret, err := exportServices(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := res.ImportServices(ctx, ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
}

func importIngressesList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	ret, err := exportIngresses(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := res.ImportIngresses(ctx, ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
}

func importConfigMapsList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	ret, err := exportConfigMaps(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := res.ImportConfigMaps(ctx, ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
}

func importStoragesList(ctx *gin.Context, kube *kubernetes.Kube, vol clients.Volumes) (*kubtypes.ImportResponse, error) {
	ret, err := exportStorages(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := vol.ImportStorages(ctx, ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}
	return resp, nil
}

func importVolumesList(ctx *gin.Context, kube *kubernetes.Kube, vol clients.Volumes) (*kubtypes.ImportResponse, error) {
	quotas, err := kube.GetPersistentVolumeClaimsList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	ret, err := model.ParseKubePersistentVolumeClaimList(quotas, false)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := vol.ImportVolumes(ctx, *ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}
	return resp, nil
}

func ImportAllHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	perm := ctx.MustGet(m.PermClient).(clients.Permissions)
	res := ctx.MustGet(m.ResClient).(clients.Resource)
	vol := ctx.MustGet(m.VolClient).(clients.Volumes)

	ret := make(kubtypes.ImportResponseTotal)

	respNs, err := importNamespaceList(ctx, kube, perm)
	if err != nil {
		ctx.Error(err)
		return
	}
	ret["namespaces"] = *respNs

	respDepl, err := importDeploymentsList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		return
	}
	ret["deployments"] = *respDepl

	respSvc, err := importServicesList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		return
	}
	ret["services"] = *respSvc

	respIngr, err := importIngressesList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		return
	}
	ret["ingresses"] = *respIngr

	respCM, err := importConfigMapsList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		return
	}
	ret["configmaps"] = *respCM

	respStorages, err := importStoragesList(ctx, kube, vol)
	if err != nil {
		ctx.Error(err)
		return
	}
	ret["storages"] = *respStorages

	respVolumes, err := importVolumesList(ctx, kube, vol)
	if err != nil {
		ctx.Error(err)
		return
	}
	ret["volumes"] = *respVolumes

	ctx.JSON(http.StatusAccepted, ret)
}
