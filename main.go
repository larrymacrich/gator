package main

import (
	"log"

	"github.com/larrymacrich/gator/cli"
	_ "github.com/lib/pq"
)

func main() {
	err := cli.RunCommand()
	if err != nil {
		log.Fatal(err)
	}
}
