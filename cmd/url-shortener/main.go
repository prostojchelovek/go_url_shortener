package main

import (
	"fmt"

	"github.com/prostojchelovek/go_url_shortener/iternal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
