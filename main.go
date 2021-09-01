package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type stringSlice []string

func (ss *stringSlice) String() string {
	return fmt.Sprintf("%s", *ss)
}

func (ss *stringSlice) Set(s string) error {
	*ss = append(*ss, s)
	return nil
}

var (
	x  = flag.String("x", "GET", "method")
	d  = flag.String("d", "", "request body")
	v  = flag.Bool("v", false, "show request")
	hs stringSlice
)

func main() {
	flag.Var(&hs, "H", "request header")
	flag.Parse()
	url := os.Args[len(os.Args)-1]
	curl, err := NewCurl(*x, url, *d)
	if err != nil {
		log.Fatal(err)
	}

	for _, h := range hs {
		curl.AddHeader(h)
	}

	if *v {
		fmt.Println(curl.DumpRequest())
	}

	err = curl.Do()
	if err != nil {
		log.Fatal(err)
	}
	b, err := curl.ResponseBody()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
