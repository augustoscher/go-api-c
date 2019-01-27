package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/augustoscher/API-C/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

//Login realiza login e gera token
func Login(w http.ResponseWriter, req *http.Request) {
	var user models.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	if isValidUser(user) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"password": user.Password,
		})
		tokenString, error := token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}
		json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString})
	} else {
		json.NewEncoder(w).Encode(models.Exception{Message: "Usuário inválido."})
	}
}

func isValidUser(user models.User) bool {
	return user.Username == "augusto" && user.Password == "testing"
}

//TestEndpoint retorna usuario e senha de determinado token
func TestEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	var user models.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}

//EndpointLivre retorna uma string
func EndpointLivre(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("endpoint livre")
}
