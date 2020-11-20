package TGenerateByRefreshT

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func Refreshing (w http.ResponseWriter, r *http.Request) {

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

		// Generating
		generateObjectID := primitive.NewObjectID()
	
		uID, _ := primitive.ObjectIDFromHex(uTokens.Id)
		
		IDTokens, _ := primitive.ObjectIDFromHex(uTokens.IdTokens)
	
		// Генерируем токены, в зависимости от токенов
		const duration = time.Minute * 10
	
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
		resultUpdate := utils.MUpdateOne("users", "accounts", filterUID, update)
	
		if resultUpdate.MatchedCount == 0 {
	
			message := utils.Message(false, "not found user")
			
			utils.Respond(w, message)
		} else {
			
			// delete this user refresh token (uRefID)
			resultDelete := utils.DeleteRefresh(uID, IDTokens)
			
			if resultDelete.MatchedCount == 0 {
				
				message := utils.Message(false, "not found token")
				
				utils.Respond(w, message)
			} else {
		
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