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

func (repositorio Colaboradores) BuscarPorID(colaboradorID int64) (modelos.Colaborador, error) {
	linha, erro := repositorio.db.Query(
		`select f.*, u.nome from
		colaboradores f inner join usuarios u
		on u.id = f.autor_id where f.id = ?`,
		colaboradorID,
	)
	if erro != nil {
		return modelos.Colaborador{}, erro
	}
	defer linha.Close()

	var colaborador modelos.Colaborador

	if linha.Next() {
		if erro = linha.Scan(
			&colaborador.ID,
			&colaborador.Nome,
			&colaborador.Matricula,
			&colaborador.AutorID,
			&colaborador.Setor,
			&colaborador.Turno,
			&colaborador.CriadoEm,
			&colaborador.AutorNome,
		); erro != nil {
			return modelos.Colaborador{}, erro
		}
	}
	return colaborador, nil

}
