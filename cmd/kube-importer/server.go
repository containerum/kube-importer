package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"text/tabwriter"
	"time"

	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-importer/pkg/router"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func initServer(c *cli.Context) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent|tabwriter.Debug)
	for _, f := range c.GlobalFlagNames() {
		fmt.Fprintf(w, "Flag: %s\t Value: %s\n", f, c.String(f))
	}
	w.Flush()

	setupLogs(c)

	kube := kubernetes.Kube{}
	go exitOnErr(kube.RegisterClient(c.String("kubeconf")))

	res, err := setupResource(c)
	exitOnErr(err)

	perm, err := setupPermissions(c)
	exitOnErr(err)

	vol, err := setupVolumes(c)
	exitOnErr(err)

	status := model.ServiceStatus{
		Name:     c.App.Name,
		Version:  c.App.Version,
		StatusOK: true,
	}

	app := router.CreateRouter(&kube, res, perm, vol, c.StringSlice("excluded_ns"), &status, c.Bool("cors"))

	srv := &http.Server{
		Addr:    ":" + c.String("port"),
		Handler: app,
	}

	go exitOnErr(srv.ListenAndServe())

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // subscribe on interrupt event
	<-quit                            // wait for event
	logrus.Infoln("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
