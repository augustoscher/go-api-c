# api-c

### Pré-Requisitos 
1) Ter Go instalado.  
2) Ter ferramenta para testes de endpoints (Postman, Curl).  

### Passos para iniciar a aplicação:  
1) Clonar o projeto em seu workspace go (GOPATH).  
2) Através do terminal, acessar o diretório API-C em seu workspace.  
3) Executar: go run main.go  

### Execução de testes  
Através do Postman:  

1) Fazer login na api:    
executar chamada POST em http://localhost:3000/login  
body:  
{
	"username": "augusto",
	"password": "testing"
}
será retornado um JWT, por exemplo: 
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InRlc3RpbmciLCJ1c2VybmFtZSI6ImF1Z3VzdG8ifQ.77xiMA_tBkEhwigv9n5iAk_i59Y63eRLWuk3AehARO8"}

2) Buscar histórico de consultas por CPF.  
adicionar o Bearer Token nas autorizações.  
executar chamada GET http://localhost:3000/consultascpf/02206795012  
será retornada o histórico de consulta para o cpf informado.  
A url http://localhost:3000/consultascpf retorna todo o historico de consultas.  

3) Buscar histórico de movimentações financeira por cpf  
adicionar o Bearer Token nas autorizações.  
executar chamada GET http://localhost:3000/movimentosfinanceiro/02506196013  
será retornada o histórico de movimentos financeiros do cpf informado.  
A url http://localhost:3000/movimentosfinanceiro retorna todo o historico de movimentações.  

4) Buscar histórico de movimentações de cartão de crédito por cpf  
adicionar o Bearer Token nas autorizações.  
executar chamada GET http://localhost:3000/movimentoscartao/02506196013.  
será retornada o histórico de transações de cartão de crédito relacionadas ao cpf informado.  
A url http://localhost:3000/movimentoscartao retorna todo o historico de movimentações de cartão.  