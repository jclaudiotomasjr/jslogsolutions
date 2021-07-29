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
	"strconv"

	"github.com/gorilla/mux"
)

//CriarColaborador adiciona um novo colaborador no BD
func CriarColaborador(w http.ResponseWriter, r *http.Request) {
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

	var colaborador modelos.Colaborador
	if erro = json.Unmarshal(corpoRequisicao, &colaborador); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	colaborador.AutorID = usuarioID

	if erro = colaborador.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeColaboradores(db)
	colaborador.ID, erro = repositorio.Criar(colaborador)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusCreated, colaborador)

}

//BuscarColaboradores retorna todos os colabores do BD
func BuscarColaboradores(w http.ResponseWriter, r *http.Request) {

}

//BuscarColaborador retorna um colaborador específico do BD
func BuscarColaborador(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	colaboradorID, erro := strconv.ParseInt(parametros["colaboradorId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return	
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro) 
		return	
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeColaboradores(db)
	colaborador, erro := repositorio.BuscarPorID(colaboradorID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, colaborador)
}

//AtualizarColaborador atualiza um colaborador específico do BD
func AtualizarColaborador(w http.ResponseWriter, r *http.Request) {

}

//DeletarColaborador deleta um colaborador específico do BD
func DeletarColaborador(w http.ResponseWriter, r *http.Request) {

}
