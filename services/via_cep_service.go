package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jvmr535/fullcycle-goexpert-desafio-multithread/dtos"
)

func FetchViaCEP(cep string) (*dtos.ViaCEPResponse, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching ViaCEP: %s", resp.Status)
	}
	var result dtos.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
