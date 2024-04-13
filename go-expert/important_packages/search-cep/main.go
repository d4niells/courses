package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching via cep api: %v", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshaling response: %v", err)
		}

		fmt.Println(data)

		file, err := os.Create("cities.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating a file", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %v", data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error saving data into the file", err)
		}
	}
}
