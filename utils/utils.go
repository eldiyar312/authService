package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func AccessTokenGenerate (uId string) string {

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = uId
	atClaims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, _ := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	return token
}


func RefreshTokenGenerate (accessToken string) string {

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(accessToken))

	return token
}