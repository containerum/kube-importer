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

	"github.com/containerum/kube-importer/pkg/clients"
	"github.com/containerum/kube-importer/pkg/kierrors"
	headers "github.com/containerum/utils/httputil"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
)

func CreateRouter(kube *kubernetes.Kube, res *clients.Resource, perm *clients.Permissions, vol *clients.Volumes, excluded []string, enableCORS bool) http.Handler {
	e := gin.New()
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
		cfg.AddAllowMethods(http.MethodDelete)
		cfg.AddAllowHeaders(headers.UserRoleXHeader, headers.UserIDXHeader, headers.UserNamespacesXHeader)
		e.Use(cors.New(cfg))
		e.Use(middleware.RegisterKubeClient(kube))
		e.Use(middleware.RegisterResourceClient(res))
		e.Use(middleware.RegisterPermissionsClient(perm))
		e.Use(middleware.RegisterVolumesClient(vol))
	}
	e.Group("/static").
		StaticFS("/", static.HTTP)
	/* System */
	e.Use(ginrus.Ginrus(logrus.WithField("component", "gin"), time.RFC3339, true))
	e.Use(gonic.Recovery(kierrors.ErrInternalError, cherrylog.NewLogrusAdapter(logrus.WithField("component", "gin"))))
}

func initRoutes(e gin.IRouter) {
	e.GET("/namespaces", h.ExportNamespacesListHandler)
	e.POST("/namespaces", h.ImportNamespacesListHandler)

	e.GET("/deployments", h.ExportDeploymentsListHandler)
	e.POST("/deployments", h.ImportDeploymentsListHandler)

	e.GET("/services", h.ExportServicesListHandler)
	e.POST("/services", h.ImportServicesListHandler)

	e.GET("/configmaps", h.ExportConfigMapsListHandler)
	e.POST("/configmaps", h.ImportConfigMapsListHandler)

	e.GET("/ingresses", h.ExportIngressesListHandler)
	e.POST("/ingresses", h.ImportIngressesListHandler)

	e.GET("/storages", h.ExportStoragesListHandler)
	e.POST("/storages", h.ImportStoragesListHandler)

	e.GET("/volumes", h.ExportVolumesListHandler)
	e.POST("/volumes", h.ImportVolumesListHandler)

	e.POST("/all", h.ImportAllHandler)
}
