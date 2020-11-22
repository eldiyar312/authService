package auth

import (
	"encoding/json"
	"net/http"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func DeleteTokens (w http.ResponseWriter, r *http.Request) {

	type User struct {
		Id string
	}

	var user User

	json.NewDecoder(r.Body).Decode(&user)
	
	// Generate
	uID, _ := primitive.ObjectIDFromHex(user.Id)

	// delete 
	result := DeleteUserAllTokens(uID)

	if result.MatchedCount == 0 {

		message := utils.Message(false, "not found user")
	
		utils.Respond(w, message)	
	} else {
		
		// send user
		message := utils.Message(true, "success remove all tokens")
	
		utils.Respond(w, message)
	}
}