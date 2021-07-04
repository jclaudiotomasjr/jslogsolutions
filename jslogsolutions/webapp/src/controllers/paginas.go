package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//CarregarTelaDeLogin carregar login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

//CarregarPaginaPrincipal carrega pagina home
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "home.html", nil)
}
