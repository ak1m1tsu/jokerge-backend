package main

import (
	"log"

	"github.com/ak1m1tsu/jokerge/internal/app/api"
)

func main() {
	if app, err := api.New(); err != nil {
		log.Println(err.Error())
	} else {
		log.Println(app.Run())
	}
}
