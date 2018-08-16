package handlers

import (
	"net/http"

	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"git.containerum.net/ch/kube-api/pkg/model"
	"github.com/containerum/cherry/adaptors/gonic"
	kubtypes "github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-importer/pkg/kierrors"
	importerModel "github.com/containerum/kube-importer/pkg/model"
	m "github.com/containerum/kube-importer/pkg/router/middleware"
	"github.com/gin-gonic/gin"
)

func ExportNamespaceListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	ret, err := exportNamespaces(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

func ExportDeploymentsListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	ret, err := exportDeployments(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

func ExportServicesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	ret, err := exportServices(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

func ExportIngressesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	ret, err := exportIngresses(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

func ExportConfigMapsListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	ret, err := exportConfigMaps(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

func ExportStoragesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	ret, err := exportStorages(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

func ExportVolumesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	ret, err := exportVolumes(kube)
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}
	ctx.JSON(http.StatusOK, ret)
}

func exportNamespaces(kube *kubernetes.Kube) (filteredNsList kubtypes.NamespacesList, err error) {
	quotas, err := kube.GetNamespaceQuotaList("")
	if err != nil {
		return
	}

	nsList, err := model.ParseKubeResourceQuotaList(quotas)
	if err != nil {
		return
	}

	nss, err := kube.GetNamespaceList("")
	if err != nil {
		return
	}

	nsList = importerModel.AddNamespacesWithoutRQ(nsList, nss)

	for _, ns := range nsList.Namespaces {
		if !m.IsExcluded(ns.ID) {
			filteredNsList.Namespaces = append(filteredNsList.Namespaces, ns)
		}
	}
	return
}

func exportDeployments(kube *kubernetes.Kube) (filteredDeplList kubtypes.DeploymentsList, err error) {
	quotas, err := kube.GetDeploymentList("", "")
	if err != nil {
		return
	}

	ret, err := model.ParseKubeDeploymentList(quotas, false)
	if err != nil {
		return
	}

	for _, depl := range ret.Deployments {
		if !m.IsExcluded(depl.Namespace) {
			filteredDeplList.Deployments = append(filteredDeplList.Deployments, depl)
		}
	}
	return
}

func exportServices(kube *kubernetes.Kube) (filteredSvcList model.ServiceWithParamList, err error) {
	quotas, err := kube.GetServiceList("")
	if err != nil {
		return
	}

	ret, err := model.ParseKubeServiceList(quotas, false)
	if err != nil {
		return
	}

	for _, svc := range ret.Services {
		if !m.IsExcluded(svc.Namespace) {
			filteredSvcList.Services = append(filteredSvcList.Services, svc)
		}
	}
	return
}

func exportIngresses(kube *kubernetes.Kube) (filteredIngrList kubtypes.IngressesList, err error) {
	ingresses, err := kube.GetIngressList("")
	if err != nil {
		return
	}

	ret, err := model.ParseKubeIngressList(ingresses, false)
	if err != nil {
		return
	}

	for _, ingr := range ret.Ingress {
		if !m.IsExcluded(ingr.Namespace) {
			filteredIngrList.Ingress = append(filteredIngrList.Ingress, ingr)
		}
	}
	return
}

func exportConfigMaps(kube *kubernetes.Kube) (filteredCMList kubtypes.ConfigMapsList, err error) {
	quotas, err := kube.GetConfigMapList("")
	if err != nil {
		return
	}

	ret, err := model.ParseKubeConfigMapList(quotas, false)
	if err != nil {
		return
	}

	for _, cm := range ret.ConfigMaps {
		if !m.IsExcluded(cm.Namespace) {
			filteredCMList.ConfigMaps = append(filteredCMList.ConfigMaps, cm)
		}
	}
	return
}

func exportStorages(kube *kubernetes.Kube) (storagesList model.StorageList, err error) {
	storageList, err := kube.GetStorageClassesList()
	if err != nil {
		return
	}

	ret, err := model.ParseStoragesList(storageList)
	if err != nil {
		return
	}
	return *ret, nil
}

func exportVolumes(kube *kubernetes.Kube) (filteredVolList kubtypes.VolumesList, err error) {
	pvc, err := kube.GetPersistentVolumeClaimsList("")
	if err != nil {
		return
	}

	ret, err := model.ParseKubePersistentVolumeClaimList(pvc, false)
	if err != nil {
		return
	}

	for _, vol := range ret.Volumes {
		if !m.IsExcluded(vol.Namespace) {
			filteredVolList.Volumes = append(filteredVolList.Volumes, vol)
		}
	}
	return
}