package main

import (
	"flag"
	"fmt"
	"os"
)

var opt struct {
	port int
	dump bool
}

func init() {
	flag.IntVar(&opt.port, "p", 8000, "Listen port")
	flag.BoolVar(&opt.dump, "dump", false, "Dump RSS feeds")
}

func main() {
	flag.Parse()

	if opt.dump {
		podcastWriter(os.Stdout)
		return
	}

	addr := fmt.Sprintf(":%d", opt.port)
	fmt.Printf("Listening %s\n", addr)
	serve(addr)
}
