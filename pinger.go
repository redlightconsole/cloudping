package cloudping

import (
	"context"
	"fmt"
	"github.com/redlightconsole/cloudping/pkg/build"
	"sort"
	"sync"
	"time"
)

type PingResult struct {
	Err    error
	Target RegionTarget
	Pings  []int
}

func (r *PingResult) Min() int {
	return r.Pings[0]
}

func (r *PingResult) Max() int {
	return r.Pings[len(r.Pings)-1]
}

func (r *PingResult) Average() int {
	var avg int
	for _, lat := range r.Pings {
		avg += lat
	}
	return avg / len(r.Pings)
}

type Pinger struct {
	m       *sync.Mutex
	reqType string
	repeat  int
	targets []*RegionTarget
	results []PingResult
}

func NewPinger(reqType string, c int) *Pinger {
	return &Pinger{
		m:       &sync.Mutex{},
		reqType: reqType,
		repeat:  c,
		targets: make([]*RegionTarget, 0),
		results: make([]PingResult, 0),
	}
}

func (p *Pinger) AddTarget(targets ...*RegionTarget) {
	p.targets = append(p.targets, targets...)
}

func (p *Pinger) Run(ctx context.Context) error {
	runner := NewRunner(10)
	var reqType RequestType

	switch p.reqType {
	case "http":
		reqType = RequestTypeHTTP
	case "tcp":
		reqType = RequestTypeTCP
	default:
		return fmt.Errorf("unexpected request type %s", p.reqType)
	}

	for _, t := range p.targets {
		runner.Add(1)
		go func(t *RegionTarget) {
			defer runner.Done()
			var pings []int
			var reqerr error

			for i := 1; i <= p.repeat; i++ {
				// Add a 100ms timeout between requests
				if i != 1 {
					time.Sleep(200 * time.Millisecond)
				}
				addr, err := p.formatHost(reqType, t)
				if err != nil {
					reqerr = err
					return
				}

				req := NewRequest()
				d, err := req.Do(fmt.Sprintf("cloudping/%s", build.String()), addr, reqType)
				pings = append(pings, int(d.Milliseconds()))
				reqerr = err
			}

			p.m.Lock()
			defer p.m.Unlock()

			result := PingResult{
				Err:    reqerr,
				Target: *t,
				Pings:  pings,
			}
			// Sort values in ascending order
			sort.Ints(result.Pings)

			p.results = append(p.results, result)
		}(t)
	}
	runner.Wait()
	return nil
}

func (p *Pinger) Results() []PingResult {
	return p.results
}

func (p *Pinger) formatHost(reqType RequestType, t *RegionTarget) (string, error) {
	var addr string

	switch reqType {
	case RequestTypeHTTP:
		addr = t.GetURL()
	case RequestTypeTCP:
		ip, err := t.GetIP()
		if err != nil {
			return "", err
		}
		addr = ip.String()
	default:
		return "", fmt.Errorf("unexpected request type")
	}
	return addr, nil
}
