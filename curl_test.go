package main

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	c, err := NewCurl("POST", "example.com/api/v1/user", "{}")
	if err != nil {
		t.Errorf("expected %v, actual %v", nil, err)
	}
	c.AddHeader("Access-Control-Allow-Origin: *")
	c.AddHeader("Age: 2318192")
	c.AddHeader("Cache-Control: public, max-age=315360000")
	c.AddHeader("Connection: keep-alive")
	c.AddHeader("Date: Mon, 18 Jul 2016 16:06:00 GMT")
	c.AddHeader("Server: Apache")
	c.AddHeader("Vary: Accept-Encoding")
	c.AddHeader("Via: 1.1 3dc30c7222755f86e824b93feb8b5b8c.cloudfront.net (CloudFront)")
	c.AddHeader("X-Amz-Cf-Id: TOl0FEm6uI4fgLdrKJx0Vao5hpkKGZULYN2TWD2gAWLtr7vlNjTvZw==")
	c.AddHeader("X-Backend-Server: developer6.webapp.scl3.mozilla.com")
	c.AddHeader("X-Cache: Hit from cloudfront")
	c.AddHeader("X-Cache-Info: cached")

	dump := strings.Split(strings.Replace(c.DumpRequest(), "\r", "", -1), "\n")
	expectedDump := strings.Split(`POST /api/v1/user HTTP/1.1
Host: example.com
Access-Control-Allow-Origin: *
Age: 2318192
Cache-Control: public, max-age=315360000
Connection: keep-alive
Date: Mon, 18 Jul 2016 16:06:00 GMT
Server: Apache
Vary: Accept-Encoding
Via: 1.1 3dc30c7222755f86e824b93feb8b5b8c.cloudfront.net (CloudFront)
X-Amz-Cf-Id: TOl0FEm6uI4fgLdrKJx0Vao5hpkKGZULYN2TWD2gAWLtr7vlNjTvZw==
X-Backend-Server: developer6.webapp.scl3.mozilla.com
X-Cache: Hit from cloudfront
X-Cache-Info: cached

{}`, "\n")
	for i, actual := range dump {
		if expectedDump[i] != actual {
			t.Errorf("%dth row: expected %#v, actual %#v", i, expectedDump[i], actual)
		}
	}
}
