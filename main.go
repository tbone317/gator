package main

import (
	"fmt"
	"log"

	"github.com/tbone317/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error to read config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("travis")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error to read config: %v", err)
	}
	fmt.Printf("Reading config again: %+v\n", cfg)
}
