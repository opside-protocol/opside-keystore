package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const appName = "keystore"

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Commands = []*cli.Command{
		{
			Name:    "encryptKey",
			Aliases: []string{},
			Usage:   "Encrypts the privatekey with a password and create a keystore file",
			Action:  encryptKey,
			Flags:   encryptKeyFlags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
