package main

import (
	"fmt"

	"github.com/simonlewi/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	err = cfg.SetUser("simon")
	if err != nil {
		fmt.Println("Error setting user:", err)
		return
	}

	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading updated config:", err)
		return
	}

	fmt.Printf("Config: %+v,\n", updatedCfg)
}
