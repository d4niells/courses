package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for {
		f, err := os.Create(fmt.Sprintf("./tmp/chunk_%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.WriteString("Secret data...")
		i++
	}
}
