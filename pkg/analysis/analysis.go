package analysis

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/redlightconsole/cloudping"
	"github.com/redlightconsole/cloudping/pkg/stripansi"
	"github.com/rodaine/table"
	"io"
	"sort"
	"strings"
	"unicode/utf8"
)

type Analysis struct {
	results []cloudping.PingResult
}

func New() *Analysis {
	return &Analysis{results: make([]cloudping.PingResult, 0)}
}

func (a *Analysis) AddResult(results ...cloudping.PingResult) *Analysis {
	a.results = append(a.results, results...)
	return a
}

func (a *Analysis) WriteOutput(w io.Writer) error {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Provider", "Region name", "Code", "Min", "Max", "Avg")
	tbl.WithWriter(w)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.WithWidthFunc(func(s string) int {
		return utf8.RuneCountInString(stripansi.Strip(s))
	})

	// Sort results
	sort.SliceStable(a.results, func(i, j int) bool {
		return a.results[i].Average() < a.results[j].Average()
	})

	for _, r := range a.results {
		if r.Err != nil {
			_, err := fmt.Fprintf(w, "%s %s error: %s\n", strings.ToUpper(r.Target.Provider), r.Target.CodeName, r.Err)
			if err != nil {
				return err
			}
			continue
		}

		tbl.AddRow(strings.ToUpper(r.Target.Provider), r.Target.Name, r.Target.CodeName, a.latencyValueColor(r.Min()), a.latencyValueColor(r.Max()), a.latencyValueColor(r.Average()))
	}
	_, _ = fmt.Fprintf(w, "\n")
	tbl.Print()
	return nil
}

func (a *Analysis) latencyValueColor(lat int) string {
	if lat < 100 {
		return color.GreenString("%d", lat)
	}
	if lat < 190 {
		return color.YellowString("%d", lat)
	}
	return color.RedString("%d", lat)
}
