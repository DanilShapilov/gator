package main

import (
	"fmt"
	"os"

	"github.com/DanilShapilov/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	cfg.SetUser("Danil")
	cfg, err = config.Read()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", cfg)

	os.Exit(0)
}
