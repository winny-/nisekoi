package main

import (
	"nisekoi/utils"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:        "calc",
			Usage:       "nisekoi calc [<github-org> | <github-org/repo>]",
			Description: "Calculate average landing PR times",
			Action: func(c *cli.Context) error {
				lookup := c.Args().First()
				if !utils.ValidateSearchTerm(lookup) {
					return cli.NewExitError("The search term doesn't conform to [<github-org> | <github-org/repo>]", 1)
				}

				username := c.String("username")
				if len(username) > 0 {
					if !utils.ValidateIdentifier(username) {
						return cli.NewExitError("The username provided is invalid", 2)
					}
				}
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "username, u",
					Usage:  "If set, average times for `USERNAME` will be displayed",
					EnvVar: "NISEKOI_USERNAME",
				},
			},
		},
	}
	app.Run(os.Args)
}
