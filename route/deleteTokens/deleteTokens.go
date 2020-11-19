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


	filter := map[string]interface{}{"_id": uID}

	update := map[string]interface{}{
		"$pull": map[string]interface{}{
			"tokens": map[string]interface{}{},
		},
	}

	utils.MUpdateOne("users", "accounts", filter, update)

	// send user
	message := utils.Message(true, "success remove all tokens")

	utils.Respond(w, message)
}