# go-api-c #

#### GOlang API with JWT authentication example ####

* GOlang-API-C  

##### Requirements:  
- Go needs to be installed.
- Some tool to endpoint testing (curl, postman)


### Setting Up 
----
##### 1. Clone the repositorie in your GOPATH  
> git clone https://augustoscher@bitbucket.org/augustoscher/api-c.git   

##### 2. Access the repositorie
> cd go-api-c  

##### 3. Install the dependencies
> go get  

or one by one  

> got get github.com/codegangsta/negroni  
> got get github.com/dgrijalva/jwt-go  
> got get github.com/gorilla/context  
> got get github.com/gorilla/mux  
> got get github.com/mitchellh/mapstructure  

##### 4. Running
> go run main.go  


### Test
----

##### 1. Login   
> POST http://localhost:3000/login  
```json
{
  "username": "augusto",
  "password": "testing"
}
```

It should return: 
```json
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InRlc3RpbmciLCJ1c2VybmFtZSI6ImF1Z3VzdG8ifQ.77xiMA_tBkEhwigv9n5iAk_i59Y63eRLWuk3AehARO8"}
```
You have to add Bearer Token in header of every request.

##### 2. Get the history of searchs by CPF     
> GET http://localhost:3000/consultascpf/02206795012  
or  
> http://localhost:3000/consultascpf  

##### 3. Get the history of financial transactions by CPF  
> GET http://localhost:3000/movimentosfinanceiro/02506196013  
or  
> GET http://localhost:3000/movimentosfinanceiro 


##### 4. Get the history of credit card transactions by CPF  
> GET http://localhost:3000/movimentoscartao/02506196013  
or  
> A url http://localhost:3000/movimentoscartao

### Next steps
----
Dockerize the API and HTTP Rest Client with docker-compose
