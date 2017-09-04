package main

import (
	"flag"
	"os"
)

const defaultAddr = ":8000"

var (
	httpMode = flag.Bool("http", false, "Start HTTP serve mode")
	httpAddr = flag.String("addr", defaultAddr, "HTTP serve `address`")
)

func main() {
	flag.Parse()

	if *httpMode {
		serve(*httpAddr)
		return
	}

	podcastWriter(os.Stdout)
}
