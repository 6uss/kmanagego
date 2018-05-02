package main

import (
	"log"

	"github.com/6uss/k_manage/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
