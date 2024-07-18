package main

import (
	"github.com/https_whoyan/Lets_Go_Book_Course/config"
	_ "github.com/lib/pq"
)

func main() {
	apl, err := config.NewApplication()
	if err != nil {
		panic(err)
	}
	apl.Run()
}
