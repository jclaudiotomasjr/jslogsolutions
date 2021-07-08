package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

//Coletores representa um repositorio de coletores
type Coletores struct {
	db *sql.DB
}

//NovoRepositorioDeColetores cria um repositorio de coletores
func NovoRepositorioDeColetores(db *sql.DB) *Coletores {
	return &Coletores{db}
}

//Criar insere um coletor do BD
func (repositorio Coletores) Criar(coletor modelos.Coletor) (int64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into coletores (nrcoletor, nrserie, marca, autor_id, estado) values (?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(coletor.NrColetor, coletor.NrSerie, coletor.Marca, coletor.AutorID, coletor.Estado)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return int64(ultimoIDInserido), nil

}
