package main

import (
	"fmt"
	"log"

	"github.com/Dhar01/Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Println("Initial config:", cfg)

	err = cfg.SetUser("jane")

	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Println("Updated config: ", updatedCfg)
}
