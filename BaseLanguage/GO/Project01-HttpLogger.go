package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{
		Transport: NewLoggedRoundTripper(http.DefaultTransport, NewDefalutHttpLogger()),
	}
	client.Get("http://www.baidu.com")
}

type HttpLogger interface {
	Request(*http.Request)
	Response(*http.Request, *http.Response, time.Duration, error)
}

type LoggedRoundTripper struct {
	tr  http.RoundTripper
	log HttpLogger
}

func NewLoggedRoundTripper(tr http.RoundTripper, log HttpLogger) *LoggedRoundTripper {
	return &LoggedRoundTripper{tr: tr, log: log}
}

func (c *LoggedRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	c.log.Request(req)
	start_time := time.Now()
	resp, err := c.tr.RoundTrip(req)
	duration := time.Since(start_time)
	c.log.Response(req, resp, duration, err)
	return resp, err
}

//-----------------------------Default HttpLogger Interface---------------------------------
func NewDefalutHttpLogger() *DefalutHttpLogger {
	return &DefalutHttpLogger{
		log: log.New(os.Stderr, "Log - ", log.LstdFlags),
	}
}

type DefalutHttpLogger struct {
	log *log.Logger
}

func (c *DefalutHttpLogger) Request(req *http.Request) {
	c.log.Printf("Request--------------------")
}

func (c *DefalutHttpLogger) Response(req *http.Request, resp *http.Response, duration time.Duration, err error) {
	duration /= time.Millisecond
	c.log.Printf("Response-------------------- Cost [%d]", duration)
}
