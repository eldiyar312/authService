package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	filter := []interface{}{bson.M{"_id": uID}, bson.M{"tokens.id": uTokensID}}

	update := bson.M{
		"$pull": bson.M{
			"tokens": bson.M{
				"id": uTokensID,
			},
		},
	}

	result, _ := utils.MFingOneUpdateOne("users", "accounts", filter, update)

	fmt.Println(result)

	// if result.MatchedCount == 0 {

	// 	message := utils.Message(false, "not found token")
		
	// 	utils.Respond(w, message)
	// } else {
		
	// 	// send user
	// 	message := utils.Message(true, "success remove token")
	
	// 	utils.Respond(w, message)
	// }
}