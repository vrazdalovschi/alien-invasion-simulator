package main

import (
	"log"

	"github.com/vrazdalovschi/alien-invasion-simulator/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
