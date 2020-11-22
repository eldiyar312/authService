package auth

import (
	"encoding/json"
	"net/http"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func DeleteRefreshT (w http.ResponseWriter, r *http.Request) {

	type User struct {
		Id string
		TokenId string
	}

	var user User

	json.NewDecoder(r.Body).Decode(&user)
	
	// Generate
	uID, _ := primitive.ObjectIDFromHex(user.Id)

	uTokensID, _ := primitive.ObjectIDFromHex(user.TokenId)


	// remove refresh token
	result := DeleteRefreshToken(uID, uTokensID)

	// checking deleted token
	if result.ModifiedCount == 0 {

		message := utils.Message(false, "not found token or user")
		
		utils.Respond(w, message)
	} else {
		
		message := utils.Message(true, "success remove token")
	
		utils.Respond(w, message)
	}
}

// delete refresh token для того чтобы запретить повторное использование )
func DeleteRefreshToken (
	uID primitive.ObjectID,
	IDTokens primitive.ObjectID,
) *mongo.UpdateResult {

	filter := bson.M{"_id": uID}

	update := map[string]interface{}{
		"$pull": map[string]interface{}{
			"tokens": map[string]interface{}{
				"id": IDTokens,
			},
		},
	}
	

	result := utils.MUpdateOne("users", "accounts", filter, update)

	return result
}