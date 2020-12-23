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
				Name: "token",
				Aliases: []string{"t"},
				Usage: "magic-cli token [decode|validate] --did <DID token>",
				Subcommands: []*cli.Command{
					{
						Name:  "decode",
						Usage: "magic-cli token decode --did <DID token>",
						Flags: []cli.Flag{
							&cli.StringFlag {
								Name:  "did",
								Usage: "Did token which must be decoded",
							},
						},
						Action: decodeDIDToken,
					},
					{
						Name:  "validate",
						Usage: "magic-cli token validate --did <DID token>",
						Flags: []cli.Flag{
							&cli.StringFlag {
								Name:  "did",
								Usage: "Did token which must be validated",
							},
						},
						Action: validateDIDToken,
					},
				},
			},
			{
				Name: "user",
				Aliases: []string{"u"},
				Usage: "magic-cli -s <secret> user --did <DID token>",
				Flags: []cli.Flag{
					&cli.StringFlag {
						Name:  "did",
						Usage: "Did token used for user info receiving",
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

	userInfo, err := m.User.GetMetadataByToken(ctx.String("did"))
	if err != nil {
		return err
	}

	fmt.Println(userInfo.String())

	return nil
}

func decodeDIDToken(ctx *cli.Context) error {
	tk, err := token.NewToken(ctx.String("did"))
	if err != nil {
		return err
	}

	claim := tk.GetClaim()
	fmt.Println(claim.String())

	return nil
}

func validateDIDToken(ctx *cli.Context) error {
	tk, err := token.NewToken(ctx.String("did"))
	if err != nil {
		return err
	}

	if err := tk.Validate(); err != nil {
		return err
	}

	return nil
}