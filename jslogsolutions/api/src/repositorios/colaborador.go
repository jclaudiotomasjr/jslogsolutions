package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

//Colaboradores representa um repositorio de colaboradores
type Colaboradores struct {
	db *sql.DB
}

//NovoRepositorioDeColaboradores cria um repositorio de colaboradores
func NovoRepositorioDeColaboradores(db *sql.DB) *Colaboradores {
	return &Colaboradores{db}
}

func (repositorio Colaboradores) Criar(colaborador modelos.Colaborador) (int64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into colaboradores (matricula, nome, autor_id, setor, turno) values (?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(colaborador.Matricula, colaborador.Nome, colaborador.AutorID, colaborador.Setor, colaborador.Turno)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return int64(ultimoIDInserido), nil

}
