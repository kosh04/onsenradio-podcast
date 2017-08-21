package main

import (
	"flag"
	"fmt"
	"os"
)

var opt struct {
	http bool
	port int
}

func init() {
	flag.BoolVar(&opt.http, "http", false, "Start http mode")
	flag.IntVar(&opt.port, "p", 8000, "Litening port when http mode")
}

func main() {
	flag.Parse()

	if opt.http {
		addr := fmt.Sprintf(":%d", opt.port)
		fmt.Printf("Listening %s\n", addr)
		serve(addr)
		return
	}

	podcastWriter(os.Stdout)
}
