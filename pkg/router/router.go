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

func CreateRouter(kube *kubernetes.Kube, res *clients.Resource, perm *clients.Permissions, enableCORS bool) http.Handler {
	e := gin.New()
	initMiddlewares(e, kube, res, perm, enableCORS)
	initRoutes(e)
	return e
}

func initMiddlewares(e gin.IRouter, kube *kubernetes.Kube, res *clients.Resource, perm *clients.Permissions, enableCORS bool) {
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
	}
	e.Group("/static").
		StaticFS("/", static.HTTP)
	/* System */
	e.Use(ginrus.Ginrus(logrus.WithField("component", "gin"), time.RFC3339, true))
	e.Use(gonic.Recovery(kierrors.ErrInternalError, cherrylog.NewLogrusAdapter(logrus.WithField("component", "gin"))))
}

func initRoutes(e gin.IRouter) {
	e.GET("/namespaces", h.ExportNamespaceList)
	e.POST("/namespaces", h.ImportNamespaceList)

	e.GET("/deployments", h.ExportDeploymentsList)
	e.POST("/deployments", h.ImportDeploymentsList)

	e.GET("/services", h.ExportServicesList)
	e.POST("/services", h.ImportServicesList)

	e.GET("/configmaps", h.ExportConfigMapsList)
	e.POST("/configmaps", h.ImportConfigMapsList)

	e.GET("/ingresses", h.ExportIngressesList)
	e.POST("/ingresses", h.ImportIngressesList)

	e.GET("/volumes", h.ExportVolumesList)
	//TODO Import volumes

	e.GET("/storages", h.ExportStoragesList)
	//TODO Import storages

	//TODO Import all
}
