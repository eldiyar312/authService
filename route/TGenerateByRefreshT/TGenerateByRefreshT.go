package TGenerateByRefreshT

import (
	"authService/utils"
	"encoding/json"
	"net/http"

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

	
	// Search for check match refresh tokens
	uID, _ := primitive.ObjectIDFromHex(uRef.Id)

	// searchRefreshT(uID, uRef.RefreshToken)

	
	// Generating
	generateObjectID := primitive.NewObjectID()

	uRefID, _ := primitive.ObjectIDFromHex(uRef.IdRefreshToken)

	// Генерируем токены, в зависимости от токенов
	accessToken := utils.AccessTokenGenerate(uRef.RefreshToken)

	refreshToken := utils.RefreshTokenGenerate(accessToken)

	hashRefresh, _ := utils.HashPassword(refreshToken)
	

	// Save in DB as bycrypt hash
	newData := map[string]interface{}{
		"id": generateObjectID,
		"access": accessToken,
		"refresh": hashRefresh,
	}

	filter := map[string]interface{}{"_id": uID}

	// update token in DB
	utils.MUpdateOne("users", "accounts", filter, newData)


	// delete this user refresh token (uRefID)
	deleteRefresh(uID, uRefID)


	// send user
	type Tokens struct {
		ID primitive.ObjectID
		Access string
		Refresh string
	}

	tokens := Tokens{generateObjectID, accessToken, refreshToken}

	json.NewEncoder(w).Encode(tokens)
}

// delete refresh token для того чтобы запретить повторное использование )
func deleteRefresh (uID primitive.ObjectID, uRefID primitive.ObjectID) {
	filter := map[string]interface{}{"_id": uID}

	update := map[string]interface{}{
		"$pull": map[string]interface{}{
			"tokens": map[string]interface{}{
				"id": uRefID,
			},
		},
	}

	utils.MUpdateOne("users", "accounts", filter, update)
}

// // search user refresh token for проверки бытья :D
// func searchRefreshT(uID primitive.ObjectID, uRefID string) {

// 	hashRefresh, _ := utils.HashPassword(uRefID)

// 	fmt.Println(uID)
// 	fmt.Println(hashRefresh)

// 	filterRefT := map[string]interface{}{"_id": uID, "tokens": []interface{}{map[string]interface{}{"refresh": hashRefresh}}}

// 	resultFind := utils.MFindOne("users", "accounts", filterRefT)

// 	fmt.Println(resultFind)
	
// }