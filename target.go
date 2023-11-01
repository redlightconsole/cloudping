package cloudping

import (
	"fmt"
	"net"
	"sync"
)

var (
	tm      = &sync.Mutex{}
	targets = make([]*RegionTarget, 0)
)

type RegionTarget struct {
	Provider string
	Name     string
	CodeName string
	Uri      string
}

func NewRegionTarget(provider, regionName, regionCodeName, uri string) *RegionTarget {
	return &RegionTarget{
		Provider: provider,
		Name:     regionName,
		CodeName: regionCodeName,
		Uri:      uri,
	}
}

func AddTarget(t ...*RegionTarget) {
	tm.Lock()
	defer tm.Unlock()
	targets = append(targets, t...)
}

func GetAllTargets() []*RegionTarget {
	return targets
}

// GetURL return HTTP URL for target
func (t *RegionTarget) GetURL() string {
	return fmt.Sprintf("http://%s?x=%s", t.Uri, mkRandomString(13))
}

// GetIP return IP for target
func (t *RegionTarget) GetIP() (*net.TCPAddr, error) {
	return net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:80", t.Uri))
}
