package modelos

//Colaboradores representa os colaboradores
type Colaboradores struct {
	Matricula int64  `json:"matricula,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Setor     string `json:"setor,omitempty"`
	Turno     string `json:"turno,omitempty"`
}
