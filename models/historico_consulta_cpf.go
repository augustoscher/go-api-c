package models

//HistoricoConsultaCpf é responsável por representar o histórico de consultas a um cpf.
type HistoricoConsultaCpf struct {
	ID            string           `json:"id,omitempty"`
	Pessoa        *Pessoa          `json:"pessoa,omitempty"`
	Data          string           `json:"data,omitempty"`
	Situacaocpf   string           `json:"situacaocpf,omitempty"`
	Consultadopor *Estabelecimento `json:"consultadopor,omitempty"`
}
