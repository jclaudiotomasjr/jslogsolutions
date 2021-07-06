package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasColetores = []Rota{
	{
		URI:                "/coletores",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarColetor,
		RequerAutenticacao: true,
	},
	{
		URI:                "/coletores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarColetores,
		RequerAutenticacao: true,
	},
	{
		URI:                "/coletores/{coletorId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarColetor,
		RequerAutenticacao: true,
	},
	{
		URI:                "/coletores{coletorId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarColetor,
		RequerAutenticacao: true,
	},
	{
		URI:                "/coletores/{coletorId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarColetor,
		RequerAutenticacao: true,
	},
}
