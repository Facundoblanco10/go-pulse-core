package main

import (
	"log"

	"github.com/Facundoblanco10/go-pulse-core/internal/api"
)

func main() {
	r := api.NewRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
