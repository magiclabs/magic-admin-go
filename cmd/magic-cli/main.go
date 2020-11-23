package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/magiclabs/magic-admin-go"
	"github.com/magiclabs/magic-admin-go/client"
	"github.com/magiclabs/magic-admin-go/token"
)

func main() {
	app := &cli.App{
		Name: "magic-cli",
		Usage: "command line utility to make requests to api and validate tokens",
		Compiled: time.Now(),
		Commands: []*cli.Command{
			{
				Name: "decode",
				Aliases: []string{"d"},
				Usage: "magic-cli decode --did <DID token>",
				Flags: []cli.Flag{
					&cli.StringFlag {
						Name:  "did",
						Usage: "Did token which must be decoded and validated",
					},
				},
				Action: decodeDIDToken,
			},
			{
				Name: "user",
				Aliases: []string{"u"},
				Usage: "magic-cli -s <secret> user --did <DID token>",
				Flags: []cli.Flag{
					&cli.StringFlag {
						Name:  "did",
						Usage: "Did token used for user metadata receiving",
					},
				},
				Action: userMetadata,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag {
				Name:    "secret",
				Usage:   "Secret token which will be used for making request to backend api",
				Aliases: []string{"s"},
				EnvVars: []string{"MAGIC_API_SECRET_KEY"},
			},
		},
	}


	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func userMetadata(ctx *cli.Context) error {
	m := client.New(ctx.String("secret"), magic.NewDefaultClient())

	meta, err := m.User.GetMetadataByToken(ctx.String("did"))
	if err != nil {
		return err
	}

	fmt.Println(meta)

	return nil
}

func decodeDIDToken(ctx *cli.Context) error {
	tk, err := token.NewToken(ctx.String("did"))
	if err != nil {
		return err
	}

	fmt.Println(tk.GetClaim())

	return nil
}