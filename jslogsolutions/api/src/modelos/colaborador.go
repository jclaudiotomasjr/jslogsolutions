package modelos

import (
	"errors"
	"strings"
	"time"
)

//Colaboradores representa os colaboradores
type Colaborador struct {
	ID        int64     `json:"id,omitempty"`
	Matricula string    `json:"matricula,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	AutorID   int64     `json:"autor_id,omitempty"`
	AutorNome string    `json:"autornome,omitempty"`
	Setor     string    `json:"setor,omitempty"`
	Turno     string    `json:"turno,omitempty"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar a publicação recebida
func (colaborador *Colaborador) Preparar() error {
	if erro := colaborador.validar(); erro != nil {
		return erro
	}

	colaborador.formatar()
	return nil
}

func (colaborador *Colaborador) validar() error {

	if colaborador.Matricula == "" {
		return errors.New("A matrícula do colaborador é obrigatório e não pode estar em branco")
	}

	if colaborador.Nome == "" {
		return errors.New("O nome do colaborador é obrigatório e não pode estar em branco")
	}
	if colaborador.Setor == "" {
		return errors.New("O setor do colaborador é obrigatório e naõ estar em branco")
	}

	if colaborador.Turno == "" {
		return errors.New("O turno do colaborador é obrigatório e não pode estar em branco")
	}

	return nil
}

func (colaborador *Colaborador) formatar() {
	colaborador.Matricula = strings.TrimSpace(colaborador.Matricula)
	colaborador.Nome = strings.TrimSpace(colaborador.Nome)
	colaborador.Setor = strings.TrimSpace(colaborador.Setor)
	colaborador.Turno = strings.TrimSpace(colaborador.Turno)

}
