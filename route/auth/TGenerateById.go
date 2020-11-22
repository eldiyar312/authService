package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TGenerateById (w http.ResponseWriter, r *http.Request) {
	
	type user struct {
		Id string
	}

	var uId user

	json.NewDecoder(r.Body).Decode(&uId)

	uID, _ := primitive.ObjectIDFromHex(uId.Id)

	// delete all tokens for this user
	result := DeleteUserAllTokens(uID)

	if result.MatchedCount == 0 {

		message := utils.Message(false, "not found user")

		utils.Respond(w, message)
	} else {

		const duration = time.Minute * 10

		// Generate
		accessToken := utils.AccessTokenGenerate(uId.Id, duration)
	
		refreshToken := utils.RefreshTokenGenerate(accessToken)
	
		hashRefresh, _ := utils.HashPassword(refreshToken) // ignore error for the sake of simplicity
		
		generateObjectID := primitive.NewObjectID()
	
		// Save in DB as bycrypt hash
		filter := map[string]interface{}{"_id": uID}
	
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
				Duration time.Duration
			}
	
			tokens := Tokens{generateObjectID, accessToken, refreshToken, duration}
			
			// send user
			json.NewEncoder(w).Encode(tokens)
		}
	}
}

// удаление всех токенов при генерации по GUID
func DeleteUserAllTokens (uID primitive.ObjectID) *mongo.UpdateResult {

	filter := map[string]interface{}{"_id": uID}

	update := map[string]interface{}{
		"$pull": map[string]interface{}{
			"tokens": map[string]interface{}{},
		},
	}

	result := utils.MUpdateOne("users", "accounts", filter, update)

	return result
}