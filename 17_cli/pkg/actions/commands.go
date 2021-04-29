package actions

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func showHostNameAndIP() error {

	hostName, err := os.Hostname()
	if err == nil {
		fmt.Printf("Running on hostname %s\n", hostName)

	} else {
		return err
	}
	ip, err := net.LookupIP(hostName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, eachIP := range ip {
		if ipv4 := eachIP.To4(); ipv4 != nil {
			fmt.Printf("IP Address: %s\n", ipv4)
		} else if ipv6 := eachIP.To16(); ipv6 != nil {
			fmt.Printf("IPv6 Address: %s\n", ipv6)

		}
	}
	return nil
}

func Commands() {
	// cli.VersionFlag = &cli.BoolFlag{
	// 	Name:    "version",
	// 	Aliases: []string{"V"},
	// 	Usage:   "print only the version",
	// }
	app := &cli.App{
		Name:     "Website Lookup CLI",
		Usage:    "Let's you query Name Servers, IPs, CNAMEs and MX records",
		Version:  "v1.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Nui Nattapon",
				Email: "nattapon.s@pm.me",
			},
		},
		Copyright: "(c) 2021 NATTAPON.me",
	}

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
				showHostNameAndIP()
				fmt.Printf("\nLooking up the name servers for '%s':\n", c.String("host"))

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
				showHostNameAndIP()
				fmt.Printf("\nLooking up the IP addresses for '%s':\n", c.String("host"))

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
				showHostNameAndIP()
				fmt.Printf("\nLooking up the CNAME for '%s':\n", c.String("host"))

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
				showHostNameAndIP()
				fmt.Printf("\nLooking up for mx records for '%s':\n", c.String("host"))

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
