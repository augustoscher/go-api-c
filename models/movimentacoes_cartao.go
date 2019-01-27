package models

//MovimentacoesCartaoCredito representa as movimentações financeiras de uma determinada pessoa
type MovimentacoesCartaoCredito struct {
	ID              string           `json:"id,omitempty"`
	Pessoa          *Pessoa          `json:"pessoa,omitempty"`
	Valor           float64          `json:"valor,omitempty"`
	Data            string           `json:"data,omitempty"`
	Estabelecimento *Estabelecimento `json:"estabelecimento,omitempty"`
	CartaoCredito   *CartaoCredito   `json:"cartaocredito,omitempty"`
}
