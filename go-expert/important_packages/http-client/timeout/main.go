package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{
		Timeout: time.Second,
	}

	res, err := c.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	print(body)
}
