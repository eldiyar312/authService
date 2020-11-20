package TGenerateByRefreshT

import (
	"encoding/json"
	"net/http"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func Refreshing (w http.ResponseWriter, r *http.Request) {

	type user struct {
		Id string
		RefreshToken string
		IdRefreshToken string
	}

	var uRef user

	json.NewDecoder(r.Body).Decode(&uRef)

	
	// Generating
	generateObjectID := primitive.NewObjectID()

	uID, _ := primitive.ObjectIDFromHex(uRef.Id)
	
	uRefID, _ := primitive.ObjectIDFromHex(uRef.IdRefreshToken)

	// Генерируем токены, в зависимости от токенов
	accessToken := utils.AccessTokenGenerate(uRef.RefreshToken)

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
		resultDelete := utils.DeleteRefresh(uID, uRefID)
		
		if resultDelete.MatchedCount == 0 {
			
			message := utils.Message(false, "not found token")
			
			utils.Respond(w, message)
		} else {
	
			// send user
			type Tokens struct {
				ID primitive.ObjectID
				Access string
				Refresh string
			}
	
			tokens := Tokens{generateObjectID, accessToken, refreshToken}
	
			json.NewEncoder(w).Encode(tokens)
		}
	}
}