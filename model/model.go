package model

//Pessoa representa as informações de uma pessoa
type Pessoa struct {
	ID             string    `json:"id,omitempty"`
	Cpf            string    `json:"cpf,omitempty"`
	Datanascimento string    `json:"datanascimento,omitempty"`
	Nome           string    `json:"nome,omitempty"`
	Endereco       *Endereco `json:"endereco,omitempty"`
}

//MovimentacoesFinanceiras representa as movimentações financeiras de uma determinada pessoa
type MovimentacoesFinanceiras struct {
	ID                  string         `json:"id,omitempty"`
	Pessoa              *Pessoa        `json:"pessoa,omitempty"`
	Valor               float64        `json:"valor,omitempty"`
	Data                string         `json:"data,omitempty"`
	Tipo                string         `json:"tipo,omitempty"`
	Conta               *ContaBancaria `json:"conta,omitempty"`
	Nomeestabelecimento string         `json:"nomeestabelecimento,omitempty"`
}

//MovimentacoesCartaoCredito representa as movimentações financeiras de uma determinada pessoa
type MovimentacoesCartaoCredito struct {
	ID              string           `json:"id,omitempty"`
	Pessoa          *Pessoa          `json:"pessoa,omitempty"`
	Valor           float64          `json:"valor,omitempty"`
	Data            string           `json:"data,omitempty"`
	Estabelecimento *Estabelecimento `json:"estabelecimento,omitempty"`
	CartaoCredito   *CartaoCredito   `json:"cartaocredito,omitempty"`
}

//Estabelecimento representa as informações de determinado estabelecimento comercial
type Estabelecimento struct {
	ID       string    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Cnpj     string    `json:"cnpj,omitempty"`
	Ie       string    `json:"ie,omitempty"`
	Endereco *Endereco `json:"endereco,omitempty"`
}

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

//ContaBancaria representa as informações de uma conta bancária de uma pessoa
type ContaBancaria struct {
	ID          string `json:"id,omitempty"`
	Instituicao string `json:"instituicao,omitempty"`
	Agencia     string `json:"agencia,omitempty"`
	Conta       string `json:"conta,omitempty"`
	Digito      string `json:"digito,omitempty"`
}

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

//HistoricoConsultaCpf é responsável por representar o histórico de consultas a um cpf.
type HistoricoConsultaCpf struct {
	ID            string           `json:"id,omitempty"`
	Pessoa        *Pessoa          `json:"pessoa,omitempty"`
	Data          string           `json:"data,omitempty"`
	Situacaocpf   string           `json:"situacaocpf,omitempty"`
	Consultadopor *Estabelecimento `json:"consultadopor,omitempty"`
}
