package main

import (
	"log"

	"github.com/DimitarL/interview-challenge-backend/pkg/app"
)

func main() {
	applicaiton := app.NewApplication()

	err := applicaiton.Start("127.0.0.1", 8080)
	if err != nil {
		log.Fatal(err)
	}
}
