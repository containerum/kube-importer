package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

//go:generate swagger generate spec -m -i ../../swagger-basic.yml -o ../../swagger.json
//go:generate swagger flatten ../../swagger.json -o ../../swagger.json
//go:generate swagger validate ../../swagger.json

var version string

func main() {
	app := cli.NewApp()
	app.Name = "kube-importer"
	app.Version = version
	app.Usage = "Kube importer utility"
	app.Flags = flags

	fmt.Printf("Starting %v %v\n", app.Name, app.Version)

	app.Action = initServer

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
