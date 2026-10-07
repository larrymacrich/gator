package main

import (
	"fmt"
	"log"

	"github.com/larrymacrich/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("larrymacrich")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
