package models

//Endereco representa as informações de um endereço de uma pessoa, instituição, estabelecimento.
type Endereco struct {
	ID     string `json:"id,omitempty"`
	Rua    string `json:"rua,omitempty"`
	Numero string `json:"numero,omitempty"`
	Cep    string `json:"cep,omitempty"`
	Cidade string `json:"cidade,omitempty"`
	Estado string `json:"estado,omitempty"`
	Pais   string `json:"pais,omitempty"`
}
