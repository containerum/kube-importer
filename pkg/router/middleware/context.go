package middleware

import (
	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"github.com/containerum/kube-importer/pkg/clients"
	"github.com/gin-gonic/gin"
)

const (
	UserNamespaces = "user-namespaces"
	UserRole       = "user-role"
	UserID         = "user-id"

	KubeClient = "kubernetes-client"
	ResClient  = "resource-client"
	PermClient = "perm-client"
)

func RegisterKubeClient(kube *kubernetes.Kube) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(KubeClient, kube)
	}
}

func RegisterResourceClient(cli *clients.Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ResClient, *cli)
	}
}

func RegisterPermissionsClient(cli *clients.Permissions) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(PermClient, *cli)
	}
}
