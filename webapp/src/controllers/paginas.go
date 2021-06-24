package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//CarregarTelaDeLogin carregar login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
