package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/deis/object-storage-cli/actions"
	"github.com/deis/object-storage-cli/config"
)

func main() {
	app := cli.NewApp()
	app.Name = "objstorage"
	app.Usage = "Use a variety of different object storage systems with a single tool"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  config.StorageTypeFlag,
			Value: config.S3StorageType.String(),
			Usage: "Specify the type of the object storage system",
		},
	}
	app.Commands = []cli.Command{
		actions.DownloadCommand,
		actions.UploadCommand,
	}
	app.Run(os.Args)
}
