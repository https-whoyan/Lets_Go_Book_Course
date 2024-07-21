package main

import (
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	apl, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	apl.Run()
}
