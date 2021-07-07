package modelos

import "time"

//Colaboradores representa os colaboradores
type Colaborador struct {
	Matricula int64     `json:"matricula,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	AutorID   int64     `json:"autorId,omitempty"`
	Setor     string    `json:"setor,omitempty"`
	Turno     string    `json:"turno,omitempty"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}
