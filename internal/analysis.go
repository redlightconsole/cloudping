package internal

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/redlightconsole/cloudping"
	"github.com/rodaine/table"
	"io"
	"sort"
	"strings"
)

type Analysis struct {
	results []cloudping.PingResult
}

func NewAnalysis() *Analysis {
	return &Analysis{results: make([]cloudping.PingResult, 0)}
}

func (a *Analysis) AddResult(results ...cloudping.PingResult) {
	a.results = append(a.results, results...)
}

func (a *Analysis) WriteOutput(w io.Writer) error {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Provider", "Region", "Min", "Max", "Avg")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.WithWriter(w)

	// Sort results
	sort.SliceStable(a.results, func(i, j int) bool {
		sort.Ints(a.results[i].Pings)
		sort.Ints(a.results[j].Pings)
		return a.results[i].Pings[0] < a.results[j].Pings[0]
	})

	for _, r := range a.results {
		if r.Err != nil {
			_, err := fmt.Fprintf(w, "%s %s error: %s\n", strings.ToUpper(r.Target.Provider), r.Target.CodeName, r.Err)
			if err != nil {
				return err
			}
			continue
		}

		var (
			min int
			max int
			avg int
		)
		sort.Ints(r.Pings)
		min = r.Pings[0]
		max = r.Pings[len(r.Pings)-1]
		for _, lat := range r.Pings {
			avg += lat
		}
		avg = avg / len(r.Pings)
		tbl.AddRow(strings.ToUpper(r.Target.Provider), r.Target.CodeName, a.latencyValueColor(min), a.latencyValueColor(max), a.latencyValueColor(avg))
	}
	_, _ = fmt.Fprintf(w, "\n")
	tbl.Print()
	return nil
}

func (a *Analysis) latencyValueColor(lat int) string {
	if lat < 50 {
		return color.GreenString("%d", lat)
	}
	if lat < 100 {
		return color.YellowString("%d", lat)
	}
	return color.RedString("%d", lat)
}
