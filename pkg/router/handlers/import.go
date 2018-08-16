package handlers

import (
	"net/http"

	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"git.containerum.net/ch/kube-api/pkg/model"
	kubtypes "github.com/containerum/kube-client/pkg/model"
	importerModel "github.com/containerum/kube-importer/pkg/model"
	m "github.com/containerum/kube-importer/pkg/router/middleware"

	"github.com/gin-gonic/gin"

	"git.containerum.net/ch/kube-api/pkg/kubeerrors"
	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/kube-importer/pkg/clients"
	"github.com/containerum/kube-importer/pkg/kierrors"
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

func importNamespaceList(ctx *gin.Context, kube *kubernetes.Kube, perm clients.Permissions) (*kubtypes.ImportResponse, error) {
	quotas, err := kube.GetNamespaceQuotaList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	nsList, err := model.ParseKubeResourceQuotaList(quotas)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	nss, err := kube.GetNamespaceList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	nsList = importerModel.AddNamespacesWithoutRQ(nsList, nss)

	resp, err := perm.ImportNamespaces(ctx, *nsList)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
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

func importDeploymentsList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	depl, err := kube.GetDeploymentList("", "")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	ret, err := model.ParseKubeDeploymentList(depl, false)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := res.ImportDeployments(ctx, *ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
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

func importServicesList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {

	quotas, err := kube.GetServiceList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	ret, err := model.ParseKubeServiceList(quotas, false)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := res.ImportServices(ctx, *ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
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

func importConfigMapsList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	quotas, err := kube.GetConfigMapList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	ret, err := model.ParseKubeConfigMapList(quotas, false)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	for i := range ret.ConfigMaps {
		ret.ConfigMaps[i].Data = nil
	}

	resp, err := res.ImportConfigMaps(ctx, *ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
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

func importIngressesList(ctx *gin.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	quotas, err := kube.GetIngressList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	ret, err := model.ParseKubeIngressList(quotas, false)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := res.ImportIngresses(ctx, *ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}

	return resp, nil
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

func importStoragesList(ctx *gin.Context, kube *kubernetes.Kube, vol clients.Volumes) (*kubtypes.ImportResponse, error) {
	storageList, err := kube.GetStorageClassesList()
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	ret, err := model.ParseStoragesList(storageList)
	if err != nil {
		gonic.Gonic(kubeerrors.ErrUnableGetResourcesList(), ctx)
		return nil, err
	}

	resp, err := vol.ImportStorages(ctx, *ret)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableCreateResource(), ctx)
		return nil, err
	}
	return resp, nil
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
