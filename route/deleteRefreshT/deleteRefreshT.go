package deleteRefreshT

import (
	"encoding/json"
	"net/http"

	"github.com/authService/token/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func DeleteRefreshT (w http.ResponseWriter, r *http.Request) {

	type user struct {
		Id string
		RefreshId string
	}

	var uRefId user

	json.NewDecoder(r.Body).Decode(&uRefId)
	
	// Generate
	uID, _ := primitive.ObjectIDFromHex(uRefId.Id)

	uRefID, _ := primitive.ObjectIDFromHex(uRefId.RefreshId)


	filter := map[string]interface{}{"_id": uID}

	update := map[string]interface{}{
		"$pull": map[string]interface{}{
			"tokens": map[string]interface{}{
				"id": uRefID,
			},
		},
	}

	utils.MUpdateOne("users", "accounts", filter, update)


	// send user
	message := utils.Message(true, "success remove token")

	utils.Respond(w, message)
}