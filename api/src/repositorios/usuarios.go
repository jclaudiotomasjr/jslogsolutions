package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

//Usuarios representa um repositorio de Usuarios
type Usuarios struct {
	db *sql.DB
}

//NovoRepositoriosDeUsuarios cria um repositorio de Usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//função
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (int64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, email, senha) values(?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro

	}
	return ultimoIDInserido, nil
}

func (repositorio Usuarios) Buscar(nomeOuEmail string) ([]modelos.Usuario, error) {
	nomeOuEmail = fmt.Sprintf("%%%s%%", nomeOuEmail) //%nomeOuEmial%

	linhas, erro := repositorio.db.Query(
		"select id, nome, email, criadoEm from usuarios where nome LIKE ? or email LIKE ?",
		nomeOuEmail, nomeOuEmail)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

//BuscarPorID retorna um usuario do banco
func (repositorio Usuarios) BuscarPorID(ID int64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()
	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}

	}
	return usuario, nil
}

func (repositorio Usuarios) Atualizar(ID int64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		return erro
	}
	return nil

}
