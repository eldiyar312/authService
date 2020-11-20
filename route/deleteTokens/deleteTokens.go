package deleteTokens

import (
	"encoding/json"
	"net/http"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func DeleteTokens (w http.ResponseWriter, r *http.Request) {

	type user struct {
		Id string
	}

	var uRefId user

	json.NewDecoder(r.Body).Decode(&uRefId)
	
	// Generate
	uID, _ := primitive.ObjectIDFromHex(uRefId.Id)

	// delete 
	filter := map[string]interface{}{"_id": uID}

	update := map[string]interface{}{
		"$pull": map[string]interface{}{
			"tokens": map[string]interface{}{},
		},
	}

	result := utils.MUpdateOne("users", "accounts", filter, update)

	if result.MatchedCount == 0 {

		message := utils.Message(false, "not found user")
	
		utils.Respond(w, message)	
	} else {
		
		// send user
		message := utils.Message(true, "success remove all tokens")
	
		utils.Respond(w, message)
	}
}