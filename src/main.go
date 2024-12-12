package main

import (
	"log"

	"github.com/thapasubham/imageProcess/src/router"
)

func main() {
	err := router.NewRouter()

	if err != nil {
		log.Fatal(err)
	}
}
