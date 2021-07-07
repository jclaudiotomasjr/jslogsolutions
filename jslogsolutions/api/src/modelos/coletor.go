package modelos

//Coletores representa os coletores do CD
type Coletor struct {
	IDNrColetor int64  `json:"idnrcoletor,omitempty"`
	NrSerie     string `json:"nrserie,omitempty"`
	AutorID     int64  `json:"autorId,omitempty"`
	Status      string `json:"defeito"`
}
