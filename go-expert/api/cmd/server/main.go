package main

import (
	"fmt"

	"github.com/d4niells/api/configs"
)

func main() {
	cfg, err := configs.Load(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
