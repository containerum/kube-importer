package router

import (
	"net/http"

	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"github.com/containerum/cherry/adaptors/cherrylog"
	"github.com/containerum/cherry/adaptors/gonic"
	h "github.com/containerum/kube-importer/pkg/router/handlers"
	"github.com/containerum/kube-importer/pkg/router/middleware"
	"github.com/containerum/kube-importer/static"
	"github.com/gin-contrib/cors"
	"github.com/sirupsen/logrus"

	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-importer/pkg/clients"
	"github.com/containerum/kube-importer/pkg/kierrors"
	"github.com/containerum/utils/httputil"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
)

func CreateRouter(kube *kubernetes.Kube, res *clients.Resource, perm *clients.Permissions, vol *clients.Volumes, excluded []string, status *model.ServiceStatus, enableCORS bool) http.Handler {
	e := gin.New()
	e.GET("/status", httputil.ServiceStatus(status))
	initMiddlewares(e, kube, res, perm, vol, excluded, enableCORS)
	initRoutes(e)
	return e
}

func initMiddlewares(e gin.IRouter, kube *kubernetes.Kube, res *clients.Resource, perm *clients.Permissions, vol *clients.Volumes, excluded []string, enableCORS bool) {
	middleware.SaveBlacklist(excluded)
	/* CORS */
	if enableCORS {
		cfg := cors.DefaultConfig()
		cfg.AllowAllOrigins = true
		e.Use(cors.New(cfg))
	}
	e.Use(middleware.RegisterKubeClient(kube))
	e.Use(middleware.RegisterResourceClient(res))
	e.Use(middleware.RegisterPermissionsClient(perm))
	e.Use(middleware.RegisterVolumesClient(vol))
	e.Group("/static").
		StaticFS("/", static.HTTP)
	/* System */
	e.Use(ginrus.Ginrus(logrus.WithField("component", "gin"), time.RFC3339, true))
	e.Use(gonic.Recovery(kierrors.ErrInternalError, cherrylog.NewLogrusAdapter(logrus.WithField("component", "gin"))))
}

func initRoutes(e gin.IRouter) {
	group := e.Group("/", httputil.RequireAdminRole(kierrors.ErrAdminRequired))
	group.GET("/namespaces", h.ExportNamespacesListHandler)
	group.POST("/namespaces", h.ImportNamespacesListHandler)

	group.GET("/deployments", h.ExportDeploymentsListHandler)
	group.POST("/deployments", h.ImportDeploymentsListHandler)

	group.GET("/services", h.ExportServicesListHandler)
	group.POST("/services", h.ImportServicesListHandler)

	group.GET("/configmaps", h.ExportConfigMapsListHandler)
	group.POST("/configmaps", h.ImportConfigMapsListHandler)

	group.GET("/ingresses", h.ExportIngressesListHandler)
	group.POST("/ingresses", h.ImportIngressesListHandler)

	group.GET("/storages", h.ExportStoragesListHandler)
	group.POST("/storages", h.ImportStoragesListHandler)

	group.GET("/volumes", h.ExportVolumesListHandler)
	group.POST("/volumes", h.ImportVolumesListHandler)

	group.POST("/all", h.ImportAllHandler)
	group.GET("/all/ws", h.ImportAllWSHandler)
}
