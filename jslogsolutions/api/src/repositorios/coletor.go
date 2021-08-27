package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
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

//BuscarPorID traz um único coletor do BD
func (repositorio Coletores) BuscarPorID(coletorID int64) (modelos.Coletor, error) {
	linha, erro := repositorio.db.Query(`
		select c.*, u.nome from coletores c inner join usuarios u
		on u.id = c.autor_id where c.id = ?`,
		coletorID,
	)
	if erro != nil {
		return modelos.Coletor{}, erro
	}
	defer linha.Close()

	var coletor modelos.Coletor

	if linha.Next() {
		if erro = linha.Scan(
			&coletor.ID,
			&coletor.NrColetor,
			&coletor.NrSerie,
			&coletor.Marca,
			&coletor.AutorID,
			&coletor.Estado,
			&coletor.CriadoEm,
			&coletor.AutorNome,
		); erro != nil {
			return modelos.Coletor{}, erro
		}
	}
	return coletor, nil
}

func (repositorio Coletores) Buscar(numeroColetorOuNumeroSerie string) ([]modelos.Coletor, error){
	numeroColetorOuNumeroSerie = fmt.Sprintf("%%%s%%", numeroColetorOuNumeroSerie)

	linhas, erro := repositorio.db.Query(
		"select id, "
	)
}