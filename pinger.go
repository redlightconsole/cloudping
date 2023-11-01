package cloudping

import (
	"context"
	"fmt"
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
	m                *sync.Mutex
	repeat           int
	concurrencyLimit int
	targets          []*RegionTarget
	results          []PingResult
	reqType          RequestType
	UserAgent        string
}

func NewPinger(c, concurrencyLimit int, reqType RequestType) *Pinger {
	return &Pinger{
		m:                &sync.Mutex{},
		repeat:           c,
		concurrencyLimit: concurrencyLimit,
		targets:          make([]*RegionTarget, 0),
		results:          make([]PingResult, 0),
		reqType:          reqType,
		UserAgent:        "cloudping",
	}
}

func (p *Pinger) SetUserAgent(ua string) {
	p.UserAgent = ua
}

func (p *Pinger) AddTarget(targets ...*RegionTarget) {
	p.targets = append(p.targets, targets...)
}

func (p *Pinger) Run(ctx context.Context) error {
	runner := NewWaitGroupLimit(p.concurrencyLimit)

	for _, t := range p.targets {
		runner.Add(1)
		go func(t *RegionTarget) {
			defer runner.Done()
			var pings []int
			var reqerr error

			for i := 1; i <= p.repeat; i++ {
				// Add a 200ms timeout between requests
				if i != 1 {
					time.Sleep(200 * time.Millisecond)
				}
				addr, err := p.formatHost(p.reqType, t)
				if err != nil {
					reqerr = err
					continue
				}

				req := NewRequest()
				d, err := req.Do(p.UserAgent, addr, p.reqType)
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
