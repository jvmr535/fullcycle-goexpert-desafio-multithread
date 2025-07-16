package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jvmr535/fullcycle-goexpert-desafio-multithread/dtos"
	"github.com/jvmr535/fullcycle-goexpert-desafio-multithread/services"
)

func fetchCEPInfo(ctx context.Context, cep string) (interface{}, string) {
	brasilApiCh := make(chan *dtos.BrasilAPIResponse)
	viaCepCh := make(chan *dtos.ViaCEPResponse)

	go func() {
		result, err := services.FetchBrasilAPI(cep)
		if err != nil {
			brasilApiCh <- nil
		} else {
			brasilApiCh <- result
		}
	}()

	go func() {
		result, err := services.FetchViaCEP(cep)
		if err != nil {
			viaCepCh <- nil
		} else {
			viaCepCh <- result
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil, "Tempo de requisição excedido"
		case brasilApiResult := <-brasilApiCh:
			if brasilApiResult != nil {
				return brasilApiResult, "Resultado da BrasilAPI"
			}
		case viaCepResult := <-viaCepCh:
			if viaCepResult != nil {
				return viaCepResult, "Resultado da ViaCEP"
			}
		}

		if ctx.Err() != nil {
			break
		}
	}
	return nil, "Nenhum resultado encontrado"
}

func main() {
	var cep string
	fmt.Print("Digite seu CEP: ")
	fmt.Scanln(&cep)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result, source := fetchCEPInfo(ctx, cep)
	fmt.Println(source + ":")
	if result != nil {
		fmt.Printf("%+v\n", result)
	} else {
		fmt.Println("Nenhum resultado encontrado")
	}
}
