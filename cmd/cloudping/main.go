package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redlightconsole/cloudping"
	"github.com/redlightconsole/cloudping/pkg/analysis"
	"github.com/redlightconsole/cloudping/pkg/build"
	"os"
)

var (
	count       = flag.Int("c", 3, "Number of requests to send for each region")
	reqType     = flag.String("t", "tcp", "Ping transport client: tcp,http (default is tcp)")
	showVer     = flag.Bool("v", false, "Show version")
	listRegions = flag.Bool("list-regions", false, "Show list of regions")
	provider    = flag.String("p", "all", "Cloud provider to ping (aws, azure, ...), leave empty to select all")
)

func main() {
	flag.Parse()

	if *showVer {
		fmt.Println("cloudping", build.String())
		os.Exit(0)
	}

	targets := make([]*cloudping.RegionTarget, 0)
	for _, t := range cloudping.GetAllTargets() {
		if *provider != "all" && *provider != t.Provider {
			continue
		}
		targets = append(targets, t)
	}

	if *listRegions {
		for _, t := range targets {
			fmt.Printf("%s-%s\n", t.Provider, t.CodeName)
		}
		os.Exit(0)
	}

	requestType, err := cloudping.ParseRequestType(*reqType)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(0)
	}

	fmt.Printf("Running latency checks with options: request type: %s, repeat: %d, provider(s): %s\n", *reqType, *count, *provider)

	p := cloudping.NewPinger(*count, 10, requestType)
	p.AddTarget(targets...)
	err = p.Run(context.Background())
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(0)
	}

	err = analysis.New().AddResult(p.Results()...).WriteOutput(os.Stdout)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(0)
	}
}
