package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/ns3777k/go-shodan/v4/shodan"
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

func printHost(j int, h *shodan.HostData) {
	udp := ""
	if h.Transport == "udp" {
		udp = "/udp"
	}

	t, _ := time.Parse(
		"2006-01-02T15:04:05.000000", h.Timestamp)

	fmt.Printf("%d> %s\t[%d%s]\t%s\t\t%s\n",
		j, au.Bold(h.IP), h.Port, udp,
		au.Green(h.Product), t.Format("02/01/2006 15h04"))

	if h.SSL != nil {
		sslv := strings.Join(h.SSL.Versions, " ")
		ssld := h.SSL.Certificate.Expires
		te, _ := time.Parse("20060102150405Z", ssld) //20191127120000Z
		fmt.Printf("  SSL: %s %s %s\n", au.Brown(sslv), te.Format("02-Jan-2006"), au.Brown(h.SSL.Certificate.Subject.CommonName))
	}

	cpe := strings.Join(h.CPE, ",")
	cpe = strings.Replace(cpe, "cpe:/", "", -1)
	if len(cpe) > 0 {
		fmt.Printf("  %s ", au.Brown(cpe))
	}
	if len(h.OS) > 0 {
		fmt.Printf("  (%s) ", au.Magenta(h.OS))
	}
	if len(h.Hostnames) > 0 {
		fmt.Printf(" ")
		for a := range h.Hostnames {
			fmt.Printf(" %s", au.Cyan(h.Hostnames[a]))
		}
	}
	if len(h.OS) > 0 || len(h.Hostnames) > 0 || len(cpe) > 0 {
		fmt.Println()
	}
}
