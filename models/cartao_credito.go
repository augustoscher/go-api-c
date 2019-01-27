package models

//CartaoCredito representa as informações de determinado cartão de crédito
type CartaoCredito struct {
	ID            string `json:"id,omitempty"`
	Nome          string `json:"nome,omitempty"`
	Numero        string `json:"numero,omitempty"`
	Dv            string `json:"dv,omitempty"`
	Emissao       string `json:"emissao,omitempty"`
	Validade      string `json:"validade,omitempty"`
	Diavencimento string `json:"diavencimento,omitempty"`
}
