package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int     `json:"customer_id"`
	Balance float64 `json:"amount"`
}

func main() {
	account := Account{Number: 89, Balance: 100.0}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	data := []byte(`{"customer_id": 1, "amount": 123.20}`)
	err = json.Unmarshal(data, &account)
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Balance)
}
