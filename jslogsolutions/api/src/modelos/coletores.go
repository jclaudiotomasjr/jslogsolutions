package modelos

//Coletores representa os coletores do CD
type Coletores struct {
	NrColetor int64  `json:"nrcoletor,omitempty"`
	NrSerie   string `json:"nrserie,omitempty"`
	Defeito   bool   `json:"ativo"`
}
