package cloudping

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// RequestType describes a type for a request type
type RequestType int

const (
	// RequestTypeHTTP is HTTP type of request
	RequestTypeHTTP RequestType = iota
	// RequestTypeTCP is TCP type of request
	RequestTypeTCP
)

// Requester is an interface to do a network request
type Requester interface {
	Do(ua, url string, reqType RequestType) (time.Duration, error)
}

// HTTPRequester is an interface for HTTP requests
type HTTPRequester interface {
	Do(req *http.Request) (*http.Response, error)
}

// TCPRequester is an interface for TCP requests
type TCPRequester interface {
	Dial(network, address string) (net.Conn, error)
}

// Request implements Requester interface
type Request struct {
	httpClient HTTPRequester
	tcpClient  TCPRequester
}

func ParseRequestType(s string) (RequestType, error) {
	switch strings.ToLower(s) {
	case "http":
		return RequestTypeHTTP, nil
	case "tcp":
		return RequestTypeTCP, nil
	}
	return 0, fmt.Errorf("request type not recognized: %s", s)
}

// NewRequest creates a new instance of Request
func NewRequest() *Request {
	return &Request{
		httpClient: &http.Client{},
		tcpClient:  &net.Dialer{},
	}
}

// DoHTTP does HTTP request for a URL by User-Agent (ua)
func (r *Request) DoHTTP(ua, url string) (time.Duration, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", ua)

	start := time.Now()
	resp, err := r.httpClient.Do(req)
	latency := time.Since(start)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return latency, nil
}

// DoTCP does TCP request to the Addr
func (r *Request) DoTCP(_, addr string) (time.Duration, error) {
	start := time.Now()
	conn, err := r.tcpClient.Dial("tcp", addr)
	if err != nil {
		return 0, err
	}
	l := time.Since(start)
	defer conn.Close()

	return l, nil
}

// Do does a request. Type of request depends on reqType
func (r *Request) Do(ua, url string, reqType RequestType) (time.Duration, error) {
	if reqType == RequestTypeHTTP {
		return r.DoHTTP(ua, url)
	}
	return r.DoTCP(ua, url)
}
