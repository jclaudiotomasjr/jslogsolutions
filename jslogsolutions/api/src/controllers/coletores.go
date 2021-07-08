package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CriarColetor adiciona um coletor no BD
func CriarColetor(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var coletor modelos.Coletor
	if erro = json.Unmarshal(corpoRequisicao, &coletor); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	coletor.AutorID = usuarioID

	if erro = coletor.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeColetores(db)
	coletor.ID, erro = repositorio.Criar(coletor)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, coletor)

}

//BuscarColetores retorna todos os coletores do BD
func BuscarColetores(w http.ResponseWriter, r *http.Request) {

}

//BuscarColetor retorna um coletor específico do BD
func BuscarColetor(w http.ResponseWriter, r *http.Request) {

}

//AtualizarColetor atualiza um coletor específico no BD
func AtualizarColetor(w http.ResponseWriter, r *http.Request) {

}

//DeletarColetor deleta um coletor específico do BD
func DeletarColetor(w http.ResponseWriter, r *http.Request) {

}
