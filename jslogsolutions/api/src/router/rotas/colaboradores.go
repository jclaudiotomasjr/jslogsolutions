package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasColaboradores = []Rota{
	{
		URI:                "/colaboradores",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarColaborador,
		RequerAutenticacao: true,
	},
	{
		URI:                "/colaboradores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarColaboradores,
		RequerAutenticacao: true,
	},
	{
		URI:                "/colaboradores/{colaboradorId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarColaborador,
		RequerAutenticacao: true,
	},
	{
		URI:                "/colaboradores/{colaboradorId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarColaborador,
		RequerAutenticacao: true,
	},
	{
		URI:                "/colaboradores/{colaboradorId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarColaborador,
		RequerAutenticacao: true,
	},
}
