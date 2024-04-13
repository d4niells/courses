package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type ViaCep struct {
	Cep         string `json:"cep"`
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
	mux := http.NewServeMux()

	mux.HandleFunc("/{cep}", handler)

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cepParam := r.PathValue("cep")

	if cepParam == "" {
		log.Println("Invalid query param cep")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(5)*time.Second)
	defer cancel()

	cep, err := getViaCep(ctx, cepParam)
	if err != nil {
		log.Printf("Error getting via cep: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	t, err := template.New("cep.html").ParseFiles("cep.html")
	if err != nil {
		log.Printf("Error parsing template file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, cep)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
}

func getViaCep(ctx context.Context, cepParam string) (*ViaCep, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json", cepParam)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var cep *ViaCep
	err = json.NewDecoder(res.Body).Decode(&cep)
	if err != nil {
		return nil, err
	}

	return cep, nil
}
