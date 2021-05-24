package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

var Version string

var au aurora.Aurora

func parseArgs() (string, string, string, bool) {
	var query, net, ip string
	var compact, color bool

	au = aurora.NewAurora(true)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", au.Bold(os.Args[0]))
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n  Version: %s\n", au.Bold(Version))
	}

	flag.StringVar(&query, "q", "", "query ['!http']")
	flag.StringVar(&net, "n", "", "net [192.168.0.0/24]")
	flag.StringVar(&ip, "i", "", "ip [192.168.0.1]")
	flag.BoolVar(&compact, "c", false, "compact, no detail")
	flag.BoolVar(&color, "b", false, "black & white, no color")

	flag.Parse()

	au = aurora.NewAurora(!color)

	return query, net, ip, compact
}
