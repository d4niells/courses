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

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	print(string(body))
}
