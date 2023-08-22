package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redlightconsole/cloudping"
	"github.com/redlightconsole/cloudping/internal"
	"github.com/redlightconsole/cloudping/internal/build"
	"os"
)

var (
	count       = flag.Int("c", 3, "Number of count")
	reqType     = flag.String("t", "tcp", "Ping transport client: tcp,http (default is tcp)")
	showVer     = flag.Bool("v", false, "Show version")
	listRegions = flag.Bool("list-regions", false, "Show list of regions")
	provider    = flag.String("provider", "", "Cloud provider to ping (aws, gcp, azure, ...), leave empty to select all")
)

func main() {
	flag.Parse()

	if *showVer {
		fmt.Println("cloudping", build.String())
		os.Exit(0)
	}

	targets := make([]*cloudping.RegionTarget, 0)
	for _, t := range cloudping.GetAllTargets() {
		if *provider != "" && *provider != t.Provider {
			continue
		}
		targets = append(targets, t)
	}

	if *listRegions {
		for _, t := range targets {
			fmt.Println(t.CodeName)
		}
		os.Exit(0)
	}

	fmt.Printf("Running latency checks with options: request type: %s, repeat: %d\n", *reqType, *count)

	p := cloudping.NewPinger(*reqType, *count)
	p.AddTarget(targets...)
	err := p.Run(context.Background())
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(0)
	}

	a := internal.NewAnalysis()
	a.AddResult(p.Results()...)
	err = a.WriteOutput(os.Stdout)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(0)
	}
}
