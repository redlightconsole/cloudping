package cloudping

import (
	"context"
	"fmt"
	"github.com/redlightconsole/cloudping/internal/build"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type PingResult struct {
	Err    error
	Target RegionTarget
	Pings  []int64
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
	var wg sync.WaitGroup
	var reqType RequestType

	switch p.reqType {
	case "http":
		reqType = RequestTypeHTTP
	case "https":
		reqType = RequestTypeHTTP
	case "tcp":
		reqType = RequestTypeTCP
	default:
		return fmt.Errorf("unexpected request type %s", p.reqType)
	}

	for _, t := range p.targets {
		wg.Add(1)
		go func(t *RegionTarget) {
			defer wg.Done()
			var pings []int64
			var reqerr error

			for i := 1; i <= p.repeat; i++ {
				// Add a 100ms timeout between requests
				if i != 1 {
					time.Sleep(100 * time.Millisecond)
				}
				addr, err := p.formatHost(reqType, t)
				if err != nil {
					reqerr = err
					return
				}

				req := NewRequest()
				d, err := req.Do(fmt.Sprintf("cloudping/%s", build.String()), addr, reqType)
				pings = append(pings, d.Milliseconds())
				reqerr = err
			}

			p.m.Lock()
			defer p.m.Unlock()

			p.results = append(p.results, PingResult{
				Err:    reqerr,
				Target: *t,
				Pings:  pings,
			})
		}(t)
	}
	wg.Wait()
	return nil
}

func (p *Pinger) Results() []PingResult {
	return p.results
}

func (p *Pinger) formatHost(reqType RequestType, t *RegionTarget) (string, error) {
	var addr string

	switch reqType {
	case RequestTypeHTTP:
		addr = fmt.Sprintf("https://%s?x=%s", t.Host, mkRandomString(13))
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
