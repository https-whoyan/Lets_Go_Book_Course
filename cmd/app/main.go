package main

import (
	"flag"

	"github.com/https_whoyan/Lets_Go_Book_Course/config"
)

func main() {
	addr := flag.String("addr", ":4000", "http service address")
	flag.Parse()
	apl := config.NewApplication()
	apl.Run(*addr)
}
