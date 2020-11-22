package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func TGenerateByrefreshT (w http.ResponseWriter, r *http.Request) {

	type user struct {
		Id string
		IdTokens string
		RefreshToken string
		AccessToken string
	}

	var uTokens user

	json.NewDecoder(r.Body).Decode(&uTokens)
	
	match := checkRefreshT(uTokens.AccessToken, uTokens.RefreshToken)

	if !match {

		message := utils.Message(false, "access or refresh token not correct")

		utils.Respond(w, message)
	} else {

		// Generating in primitive id
		uID, _ := primitive.ObjectIDFromHex(uTokens.Id)
		
		IDTokens, _ := primitive.ObjectIDFromHex(uTokens.IdTokens)
		
		
		// delete this user refresh token
		resultDelete := utils.DeleteRefresh(uID, IDTokens)

		if resultDelete.MatchedCount == 0 {
	
			message := utils.Message(false, "not found user or dont correct id tokens")
			
			utils.Respond(w, message)
		} else {

		// Генерируем токены, в зависимости от токенов
		generateObjectID := primitive.NewObjectID()

		duration := time.Minute * 10
	
		accessToken := utils.AccessTokenGenerate(uTokens.RefreshToken, duration)
	
		refreshToken := utils.RefreshTokenGenerate(accessToken)
	
		hashRefresh, _ := utils.HashPassword(refreshToken)

			// Save in DB as bycrypt hash
			filterUID := map[string]interface{}{"_id": uID}
		
			newData := map[string]interface{}{
				"id": generateObjectID,
				"access": accessToken,
				"refresh": hashRefresh,
			}
		
			update := map[string]interface{}{
				"$push": map[string]interface{}{
					"tokens": newData,
				},
			}
		
			// create new access and refresh tokens in DB
			var _ = utils.MUpdateOne("users", "accounts", filterUID, update) //

			// send user
			type Tokens struct {
				ID primitive.ObjectID
				Access string
				Refresh string
				Duration time.Duration
			}
	
			tokens := Tokens{generateObjectID, accessToken, refreshToken, duration}
	
			json.NewEncoder(w).Encode(tokens)
		}
	}
}

// проверяем, тот ли refresh token который выдан вместе с access token
func checkRefreshT (access string, refresh string) bool {
	
	refreshToken := utils.RefreshTokenGenerate(access)

	if refresh == refreshToken {
		return true
	} else {
		return false
	}
}