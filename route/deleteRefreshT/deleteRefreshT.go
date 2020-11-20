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


	// remove refresh token
	result := utils.DeleteRefresh(uID, uRefID)

	if result.MatchedCount == 0 {

		message := utils.Message(false, "not found token")
		
		utils.Respond(w, message)
	} else {
		
		// send user
		message := utils.Message(true, "success remove token")
	
		utils.Respond(w, message)
	}
}