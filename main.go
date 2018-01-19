package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	//URL to st2 webhook
	URL = flag.String("url", "https://ubuntu-st2/api/v1/webhooks/netpro", "st2 webhook")
	//FILE is object group where IP list is included
	FILE    = flag.String("f", "./targets/", "file to read targets from")
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
}
