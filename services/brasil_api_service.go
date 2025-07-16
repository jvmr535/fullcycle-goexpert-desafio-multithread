package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jvmr535/fullcycle-goexpert-desafio-multithread/dtos"
)

func FetchBrasilAPI(cep string) (*dtos.BrasilAPIResponse, error) {
	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching BrasilAPI: %s", resp.Status)
	}
	var result dtos.BrasilAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
