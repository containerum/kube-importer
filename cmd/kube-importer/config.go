package main

import (
	"errors"
	"net/url"

	"github.com/containerum/kube-importer/pkg/clients"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "DEBUG",
		Name:   "debug",
		Usage:  "start the server in debug mode",
	},
	cli.StringFlag{
		EnvVar: "PORT",
		Name:   "port",
		Value:  "1666",
		Usage:  "port for kube-importer server",
	},
	cli.StringFlag{
		EnvVar: "KUBE_CONF",
		Name:   "kubeconf",
		Usage:  "config file for kubernetes apiserver client",
	},
	cli.BoolFlag{
		EnvVar: "TEXTLOG",
		Name:   "textlog",
		Usage:  "output log in text format",
	},
	cli.BoolFlag{
		EnvVar: "CORS",
		Name:   "cors",
		Usage:  "enable CORS",
	},
	cli.StringSliceFlag{
		EnvVar: "EXCLUDED_NS",
		Name:   "excluded_ns",
		Usage:  "namespaces excluded from import",
	},
	cli.StringFlag{
		EnvVar: "RESOURCE",
		Name:   "resource",
		Value:  "http",
		Usage:  "resource server client type",
	},
	cli.StringFlag{
		EnvVar: "RESOURCE_URL",
		Name:   "resource_url",
		Usage:  "resource server url",
	},
	cli.StringFlag{
		EnvVar: "PERMISSIONS",
		Name:   "permissions",
		Value:  "http",
		Usage:  "permissions client type",
	},
	cli.StringFlag{
		EnvVar: "PERMISSIONS_URL",
		Name:   "permissions_url",
		Usage:  "permissions url",
	},
	cli.StringFlag{
		EnvVar: "VOLUMES",
		Name:   "volumes",
		Value:  "http",
		Usage:  "volumes client type",
	},
	cli.StringFlag{
		EnvVar: "VOLUMES_URL",
		Name:   "volumes_url",
		Usage:  "volumes url",
	},
}

func setupLogs(c *cli.Context) {
	if c.Bool("debug") {
		gin.SetMode(gin.DebugMode)
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
		logrus.SetLevel(logrus.InfoLevel)
	}

	if c.Bool("textlog") {
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

func setupResource(c *cli.Context) (*clients.Resource, error) {
	switch c.String("resource") {
	case "http":
		resurl, err := url.Parse(c.String("resource_url"))
		if err != nil {
			return nil, err
		}
		client := clients.NewResourceHTTP(resurl)
		return &client, nil
	case "dummy":
		client := clients.NewDummyResource()
		return &client, nil
	default:
		return nil, errors.New("invalid resource-service client type")
	}
}

func setupPermissions(c *cli.Context) (*clients.Permissions, error) {
	switch c.String("permissions") {
	case "http":
		permurl, err := url.Parse(c.String("permissions_url"))
		if err != nil {
			return nil, err
		}
		client := clients.NewPermissionsHTTP(permurl)
		return &client, nil
	case "dummy":
		client := clients.NewDummyPermissions()
		return &client, nil
	default:
		return nil, errors.New("invalid permissions client type")
	}
}

func setupVolumes(c *cli.Context) (*clients.Volumes, error) {
	switch c.String("volumes") {
	case "http":
		permurl, err := url.Parse(c.String("volumes_url"))
		if err != nil {
			return nil, err
		}
		client := clients.NewVolumesHTTP(permurl)
		return &client, nil
	case "dummy":
		client := clients.NewDummyVolumes()
		return &client, nil
	default:
		return nil, errors.New("invalid permissions client type")
	}
}
