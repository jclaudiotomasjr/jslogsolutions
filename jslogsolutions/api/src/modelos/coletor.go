package modelos

import (
	"errors"
	"strings"
	"time"
)

//Coletores representa os coletores do CD
type Coletor struct {
	ID        int64     `json:"id,omitempty"`
	NrColetor string    `json:"nrcoletor,omitempty"`
	NrSerie   string    `json:"nrserie,omitempty"`
	Marca     string    `json:"marca,omitempty"`
	AutorID   int64     `json:"autor_id,omitempty"`
	AutorNome string    `json:"autornome,omitempty"`
	Estado    string    `json:"estado,omitempty"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar a publicação recebida
func (coletor *Coletor) Preparar() error {
	if erro := coletor.validar(); erro != nil {
		return erro
	}

	coletor.formatar()
	return nil
}

func (coletor *Coletor) validar() error {

	if coletor.NrColetor == "" {
		return errors.New("O número do coletor é obrigatório e não pode estar em branco")
	}

	if coletor.NrSerie == "" {
		return errors.New("O número de serie do coletor é obrigatório e não pode estar em branco")
	}
	if coletor.Marca == "" {
		return errors.New("A marca do coletor é obrigatório e naõ estar em branco")
	}

	if coletor.Estado == "" {
		return errors.New("O estado do coletor é obrigatório e não pode estar em branco")
	}

	return nil
}

func (coletor *Coletor) formatar() {
	coletor.NrColetor = strings.TrimSpace(coletor.NrColetor)
	coletor.NrSerie = strings.TrimSpace(coletor.NrSerie)
	coletor.Marca = strings.TrimSpace(coletor.Marca)
	coletor.Estado = strings.TrimSpace(coletor.Estado)

}
