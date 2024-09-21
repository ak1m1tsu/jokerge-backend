package main

import (
	"log"

	"github.com/ak1m1tsu/jokerge/internal/app/api"
)

func main() {
	if err := api.New().Run(); err != nil {
		log.Println(err.Error())
	}
}
