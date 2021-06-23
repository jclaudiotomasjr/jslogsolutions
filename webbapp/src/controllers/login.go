package controllers

import (
	"net/http"
)

//CarregarTelaDeLogin
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de Login"))
}
