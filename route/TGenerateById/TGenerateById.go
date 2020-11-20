package TGenerateById

import (
	"encoding/json"
	"net/http"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateTokens (w http.ResponseWriter, r *http.Request) {
	
	type user struct {
		Id string
	}

	var uId user

	json.NewDecoder(r.Body).Decode(&uId)


	// Generate
	accessToken := utils.AccessTokenGenerate(uId.Id)

	refreshToken := utils.RefreshTokenGenerate(accessToken)

	uObjectID, _ := primitive.ObjectIDFromHex(uId.Id)

	hashRefresh, _ := utils.HashPassword(refreshToken) // ignore error for the sake of simplicity
	
	generateObjectID := primitive.NewObjectID()


	// Save in DB as bycrypt hash
	filter := map[string]interface{}{"_id": uObjectID}

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

	// update token in DB
	result := utils.MUpdateOne("users", "accounts", filter, update)
	
	
	if result.MatchedCount == 0 {

		message := utils.Message(false, "user not found")

		utils.Respond(w, message)
	} else {

		// Output tokens
		type Tokens struct {
			ID primitive.ObjectID
			Access string
			Refresh string
		}

		tokens := Tokens{generateObjectID, accessToken, refreshToken}
		
		// send user
		json.NewEncoder(w).Encode(tokens)
	}
}