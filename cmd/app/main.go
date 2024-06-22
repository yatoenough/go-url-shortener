package main

import (
	"fmt"

	"github.com/yatoenough/go-url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
