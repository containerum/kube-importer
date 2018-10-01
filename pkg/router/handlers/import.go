package handlers

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"git.containerum.net/ch/kube-api/pkg/model"
	"github.com/containerum/cherry/adaptors/gonic"
	kubtypes "github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-importer/pkg/clients"
	"github.com/containerum/kube-importer/pkg/kierrors"
	m "github.com/containerum/kube-importer/pkg/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// swagger:operation POST /namespaces Import ImportNamespacesList
// Import namespaces.
//
// ---
// x-method-visibility: public
// responses:
//  '202':
//    description: import result
//    schema:
//      $ref: '#/definitions/ImportResponse'
//  default:
//    $ref: '#/responses/error'
func ImportNamespacesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	perm := ctx.MustGet(m.PermClient).(clients.Permissions)

	resp, err := importNamespacesList(ctx, kube, perm)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("namespaces"), ctx)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

// swagger:operation POST /limits Import CreateMissingLimits
// Create missing LimitRange and ResourceQuota in Kubernetes namespaces.
//
// ---
// x-method-visibility: public
// responses:
//  '202':
//    description: creation result
//    schema:
//      $ref: '#/definitions/ImportResponse'
//  default:
//    $ref: '#/responses/error'
func CreateMissingLimitsHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)

	resp, err := createMissingLimits(kube)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("namespaces"), ctx)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

// swagger:operation POST /deployments Import ImportDeploymentsList
// Import deployments.
//
// ---
// x-method-visibility: public
// responses:
//  '202':
//    description: import result
//    schema:
//      $ref: '#/definitions/ImportResponse'
//  default:
//    $ref: '#/responses/error'
func ImportDeploymentsListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)
	resp, err := importDeploymentsList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("deployments"), ctx)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

// swagger:operation POST /services Import ImportServicesList
// Import services.
//
// ---
// x-method-visibility: public
// responses:
//  '202':
//    description: import result
//    schema:
//      $ref: '#/definitions/ImportResponse'
//  default:
//    $ref: '#/responses/error'
func ImportServicesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	resp, err := importServicesList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("services"), ctx)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

// swagger:operation POST /ingresses Import ImportIngressesList
// Import ingresses.
//
// ---
// x-method-visibility: public
// responses:
//  '202':
//    description: import result
//    schema:
//      $ref: '#/definitions/ImportResponse'
//  default:
//    $ref: '#/responses/error'
func ImportIngressesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	resp, err := importIngressesList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("ingresses"), ctx)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

// swagger:operation POST /configmaps Import ImportConfigMapsList
// Import configmaps.
//
// ---
// x-method-visibility: public
// responses:
//  '202':
//    description: import result
//    schema:
//      $ref: '#/definitions/ImportResponse'
//  default:
//    $ref: '#/responses/error'
func ImportConfigMapsListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	res := ctx.MustGet(m.ResClient).(clients.Resource)

	resp, err := importConfigMapsList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("configmaps"), ctx)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

// swagger:operation POST /storages Import ImportStoragesList
// Import storages.
//
// ---
// x-method-visibility: public
// responses:
//  '202':
//    description: import result
//    schema:
//      $ref: '#/definitions/ImportResponse'
//  default:
//    $ref: '#/responses/error'
func ImportStoragesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	vol := ctx.MustGet(m.VolClient).(clients.Volumes)

	resp, err := importStoragesList(ctx, kube, vol)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("storages"), ctx)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

// swagger:operation POST /volumes Import ImportVolumesList
// Import volumes.
//
// ---
// x-method-visibility: public
// responses:
//  '202':
//    description: import result
//    schema:
//      $ref: '#/definitions/ImportResponse'
//  default:
//    $ref: '#/responses/error'
func ImportVolumesListHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	vol := ctx.MustGet(m.VolClient).(clients.Volumes)

	resp, err := importVolumesList(ctx, kube, vol)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("volumes"), ctx)
	} else {
		ctx.JSON(http.StatusAccepted, resp)
	}
}

// swagger:operation POST /all Import ImportAllHandler
// Import all resources.
//
// ---
// x-method-visibility: public
// responses:
//  '202':
//    description: import result
//    schema:
//      $ref: '#/definitions/ImportResponseTotal'
//  default:
//    $ref: '#/responses/error'
func ImportAllHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	perm := ctx.MustGet(m.PermClient).(clients.Permissions)
	res := ctx.MustGet(m.ResClient).(clients.Resource)
	vol := ctx.MustGet(m.VolClient).(clients.Volumes)

	ret := make(kubtypes.ImportResponseTotal)

	respLimits, err := createMissingLimits(kube)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableCreateLimits().AddDetails("kubernetes-limits"), ctx)
		return
	}
	ret["kubernetes-limits"] = *respLimits

	respNs, err := importNamespacesList(ctx, kube, perm)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("namespaces"), ctx)
		return
	}
	ret["namespaces"] = *respNs

	respDepl, err := importDeploymentsList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("deployments"), ctx)
		return
	}
	ret["deployments"] = *respDepl

	respSvc, err := importServicesList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("services"), ctx)
		return
	}
	ret["services"] = *respSvc

	respIngr, err := importIngressesList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("ingresses"), ctx)
		return
	}
	ret["ingresses"] = *respIngr

	respCM, err := importConfigMapsList(ctx, kube, res)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("configmaps"), ctx)
		return
	}
	ret["configmaps"] = *respCM

	respStorages, err := importStoragesList(ctx, kube, vol)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("storages"), ctx)
		return
	}
	ret["storages"] = *respStorages

	respVolumes, err := importVolumesList(ctx, kube, vol)
	if err != nil {
		ctx.Error(err)
		gonic.Gonic(kierrors.ErrUnableImportResources().AddDetails("volumes"), ctx)
		return
	}
	ret["volumes"] = *respVolumes

	ctx.JSON(http.StatusAccepted, ret)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// swagger:operation GET /all/ws Import ImportAllWS
// Import all resources with websockets response.
//
// ---
// x-method-visibility: public
// parameters:
//  - $ref: '#/parameters/UpgradeHeader'
//  - $ref: '#/parameters/ConnectionHeader'
//  - $ref: '#/parameters/SecWebSocketKeyHeader'
//  - $ref: '#/parameters/SecWebsocketVersionHeader'
// responses:
//  '101':
//    description: import response
//  default:
//    $ref: '#/responses/error'
func ImportAllWSHandler(ctx *gin.Context) {
	kube := ctx.MustGet(m.KubeClient).(*kubernetes.Kube)
	perm := ctx.MustGet(m.PermClient).(clients.Permissions)
	res := ctx.MustGet(m.ResClient).(clients.Resource)
	vol := ctx.MustGet(m.VolClient).(clients.Volumes)

	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Errorln("upgrade error:", err)
		return
	}
	defer c.Close()
	messages := make(chan kubtypes.ImportResponseTotal)
	errorch := make(chan error)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		respLimits, err := createMissingLimits(kube)
		if err != nil {
			errorch <- err
			return
		}
		messages <- kubtypes.ImportResponseTotal{"kubernetes-limits": *respLimits}

		respNs, err := importNamespacesList(ctx, kube, perm)
		if err != nil {
			errorch <- err
			return
		}
		messages <- kubtypes.ImportResponseTotal{"namespaces": *respNs}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		respDepl, err := importDeploymentsList(ctx, kube, res)
		if err != nil {
			errorch <- err
			return
		}
		messages <- kubtypes.ImportResponseTotal{"deployments": *respDepl}

		respSvc, err := importServicesList(ctx, kube, res)
		if err != nil {
			errorch <- err
			return
		}
		messages <- kubtypes.ImportResponseTotal{"services": *respSvc}

		respIngr, err := importIngressesList(ctx, kube, res)
		if err != nil {
			errorch <- err
			return
		}
		messages <- kubtypes.ImportResponseTotal{"ingresses": *respIngr}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		respCM, err := importConfigMapsList(ctx, kube, res)
		if err != nil {
			errorch <- err
			return
		}
		messages <- kubtypes.ImportResponseTotal{"configmaps": *respCM}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		respStorages, err := importStoragesList(ctx, kube, vol)
		if err != nil {
			errorch <- err
			return
		}
		messages <- kubtypes.ImportResponseTotal{"storages": *respStorages}

		respVolumes, err := importVolumesList(ctx, kube, vol)
		if err != nil {
			errorch <- err
			return
		}
		messages <- kubtypes.ImportResponseTotal{"volumes": *respVolumes}
	}()

	go func() {
		wg.Wait()
		close(messages)
	}()

	timer := time.NewTimer(600 * time.Second)
	defer timer.Stop()
	for {
		select {
		case resp, ok := <-messages:
			if !ok {
				return
			}
			if err := c.WriteJSON(resp); err != nil {
				ctx.Abort()
				return
			}
		case err := <-errorch:
			logrus.Error(err)
			if err := c.WriteJSON(err); err != nil {
				logrus.Error(err)
				ctx.Abort()
				return
			}
		case <-timer.C:
			//Timeout
			ctx.Abort()
			return
		}
	}
}

func importNamespacesList(ctx context.Context, kube *kubernetes.Kube, perm clients.Permissions) (*kubtypes.ImportResponse, error) {
	ret, err := exportNamespaces(kube)
	if err != nil {
		return nil, err
	}
	return perm.ImportNamespaces(ctx, ret)
}

func createMissingLimits(kube *kubernetes.Kube) (*kubtypes.ImportResponse, error) {
	nswq, err := getNamespacesWithoutQuota(kube)
	if err != nil {
		return nil, err
	}
	resp := kubtypes.ImportResponse{
		Imported: []kubtypes.ImportResult{},
		Failed:   []kubtypes.ImportResult{},
	}
	for _, v := range nswq.Namespaces {
		newQuota, errs := model.MakeResourceQuota(v.ID, nil, v.Resources.Hard)
		if errs != nil {
			resp.ImportFailed("quota", v.ID, fmt.Sprintf("%v", errs))
			continue
		}
		if _, err := kube.CreateNamespaceQuota(v.ID, newQuota); err != nil {
			resp.ImportFailed("quota", v.ID, err.Error())
			continue
		}
		if err := kube.CreateLimitRange(v.ID); err != nil {
			resp.ImportFailed("quota", v.ID, err.Error())
			continue
		}
		resp.ImportSuccessful("quota", v.ID)
	}
	return &resp, nil
}

func importDeploymentsList(ctx context.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	ret, err := exportDeployments(kube)
	if err != nil {
		return nil, err
	}
	return res.ImportDeployments(ctx, ret)
}

func importServicesList(ctx context.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	ret, err := exportServices(kube)
	if err != nil {
		return nil, err
	}
	return res.ImportServices(ctx, ret)
}

func importIngressesList(ctx context.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	ret, err := exportIngresses(kube)
	if err != nil {
		return nil, err
	}
	return res.ImportIngresses(ctx, ret)
}

func importConfigMapsList(ctx context.Context, kube *kubernetes.Kube, res clients.Resource) (*kubtypes.ImportResponse, error) {
	ret, err := exportConfigMaps(kube)
	if err != nil {
		return nil, err
	}
	return res.ImportConfigMaps(ctx, ret)
}

func importStoragesList(ctx context.Context, kube *kubernetes.Kube, vol clients.Volumes) (*kubtypes.ImportResponse, error) {
	ret, err := exportStorages(kube)
	if err != nil {
		return nil, err
	}
	return vol.ImportStorages(ctx, ret)
}

func importVolumesList(ctx context.Context, kube *kubernetes.Kube, vol clients.Volumes) (*kubtypes.ImportResponse, error) {
	ret, err := exportVolumes(kube)
	if err != nil {
		return nil, err
	}
	return vol.ImportVolumes(ctx, ret)
}
