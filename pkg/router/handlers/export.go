package handlers

import (
	"net/http"

	"git.containerum.net/ch/kube-api/pkg/kubeerrors"
	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"git.containerum.net/ch/kube-api/pkg/model"
	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/kube-importer/pkg/kierrors"
	importerModel "github.com/containerum/kube-importer/pkg/model"
	m "github.com/containerum/kube-importer/pkg/router/middleware"
	"github.com/gin-gonic/gin"
)

func ExportNamespaceList(ctx *gin.Context) {

	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)

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

	ctx.JSON(http.StatusOK, nsList)
}

func ExportDeploymentsList(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)

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

	ctx.JSON(http.StatusOK, ret)
}

func ExportServicesList(ctx *gin.Context) {

	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)

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

	ctx.JSON(http.StatusOK, ret)
}

func ExportIngressesList(ctx *gin.Context) {

	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)

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

	ctx.JSON(http.StatusOK, ret)
}

func ExportConfigMapsList(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)

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

	ctx.JSON(http.StatusOK, ret)
}

func ExportVolumesList(ctx *gin.Context) {

	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)

	quotas, err := kube.GetPersistentVolumeClaimsList("")
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	ret, err := model.ParseKubePersistentVolumeClaimList(quotas, false)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func ExportStoragesList(ctx *gin.Context) {

	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)

	storageList, err := kube.GetStorageClassesList()
	if err != nil {
		gonic.Gonic(kierrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	ret, err := model.ParseStoragesList(storageList)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kubeerrors.ErrUnableGetResourcesList(), ctx)
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
