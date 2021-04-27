package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	// cli.VersionFlag = &cli.BoolFlag{
	// 	Name:    "version",
	// 	Aliases: []string{"V"},
	// 	Usage:   "print only the version",
	// }

	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query Name Servers, IPs, CNAMEs and MX records"
	app.Version = "v1.0"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Nui Nattapon",
			Email: "nattapon.s@pm.me",
		},
	}
	app.Copyright = "(c) 2021 NATTAPON.me"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "host",
			Aliases:  []string{"H"},
			Value:    "oracle.com",
			Required: false,
			Usage:    "provide a hostname",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for _, eachNS := range ns {
					fmt.Println(eachNS.Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks Up the IP addresses for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for _, eachIP := range ip {
					fmt.Println(eachIP)
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks Up CNAME for a particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))

				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Look Up for mx records for a particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))

				if err != nil {
					fmt.Println(err)
					return err
				}
				for _, eachMX := range mx {
					fmt.Println(eachMX.Host, eachMX.Pref)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)

	}
}
