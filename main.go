package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/IrekRomaniuk/gost2netpro/utils"
)

var (
	//URL to st2 webhook
	URL = flag.String("url", "https://ubuntu-st2/api/v1/webhooks/netpro", "st2 webhook")
	//API of st2
	API = flag.String("api", "", "st2 api key")
	//FILE with list of remote targets
	HOSTS = flag.String("a", "./targets/pinglist.txt", "file to read targets from")
	//USER
	USER = flag.String("u", "user", "target username")
	//PASS
	PASS = flag.String("p", "pass", "target password")
	//CMD to run
	CMD     = flag.String("c", "uptime", "cmd to run on targets")
	version = flag.Bool("v", false, "Prints current version")
	// Version : Program version
	Version = "No Version Provided"
	// BuildTime : Program build time
	BuildTime = ""
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2018 @IrekRomaniuk. All rights reversed.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}
}

func main() {
	list := &utils.Netpro{Hosts: "", User: *USER, Pass: *PASS, Cmd: *CMD}
	err := list.ReadFile(*HOSTS, *USER, *PASS, *CMD)
	if err != nil {
		log.Fatal(err)
	}
	cmd, err := utils.PostPage(*URL, *API, list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cmd: %v\n", cmd)
}
