package token

import (
	"authService/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GenerateTokens (w http.ResponseWriter, r *http.Request) {

	type user struct {
		Id string
	}
	var uId user

	json.NewDecoder(r.Body).Decode(&uId)


	// Access token generate
	accessToken := utils.AccessTokenGenerate(uId.Id)

	fmt.Println("accessToken:", accessToken)


	// Refresh token generate
	tokens := strings.Split(accessToken, ".")
	token := tokens[len(tokens) - 1]

	refreshToken := utils.RefreshTokenGenerate(token)

	fmt.Println("refreshToken:", refreshToken)

}