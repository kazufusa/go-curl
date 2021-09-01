package main

import (
	"io"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"
)

var (
	httpString = regexp.MustCompile("^https?://")
)

type Curl struct {
	res         *http.Response
	req         *http.Request
	alreadyRead bool
	body        string
}

func NewCurl(method, url, body string) (curl *Curl, err error) {
	curl = new(Curl)
	if !httpString.MatchString(url) {
		url = "http://" + url
	}
	curl.req, err = http.NewRequest(method, url, strings.NewReader(body))
	return
}

func (c *Curl) Do() (err error) {
	c.res, err = new(http.Client).Do(c.req)
	return
}

func (c *Curl) ResponseBody() (string, error) {
	if c.alreadyRead {
		return c.body, nil
	}
	defer c.res.Body.Close()
	b, err := io.ReadAll(c.res.Body)
	if err != nil {
		return "", err
	}

	c.body = string(b)
	return c.body, nil
}

func (c *Curl) AddHeader(s string) {
	_s := strings.Split(s, "=")
	if len(_s) < 2 {
		return
	}
	key := _s[0]
	value := strings.Join(_s[1:], "")
	c.req.Header.Add(key, value)
}

func (c *Curl) DumpRequest() string {
	b, _ := httputil.DumpRequest(c.req, true)
	return string(b)
}
