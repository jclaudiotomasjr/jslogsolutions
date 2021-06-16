package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()
	fmt.Printf("Rodando API da JS!!! Escutando na Porta %d", config.Porta)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
