package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

func Message(status bool, message string) (map[string]interface{}) {

	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}


func AccessTokenGenerate (something string, duration time.Duration) string {

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true

	atClaims["something"] = something

	atClaims["exp"] = time.Now().Add(duration).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)

	token, _ := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	return token
}


func RefreshTokenGenerate (accessToken string) string {

	b64Token := base64.StdEncoding.EncodeToString([]byte(accessToken))

	return b64Token
}



func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// delete refresh token для того чтобы запретить повторное использование )
func DeleteRefresh (
	uID primitive.ObjectID,
	uRefID primitive.ObjectID,
) *mongo.UpdateResult {

	filterUID := map[string]interface{}{"_id": uID}

	update := map[string]interface{}{
		"$pull": map[string]interface{}{
			"tokens": map[string]interface{}{
				"id": uRefID,
			},
		},
	}
	

	result := MUpdateOne("users", "accounts", filterUID, update)

	return result
}

func MUpdateOne (
	db string, 
	table string, 
	filter map[string]interface{}, 
	update map[string]interface{},
) *mongo.UpdateResult {

	mongoURI := os.Getenv("MONGO_URI")

    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
        log.Fatal(err)
    }

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(db).Collection(table)
	
	result, errUpdate := collection.UpdateOne(context.TODO(), filter, update)

	
	if errUpdate != nil {

		log.Fatal(errUpdate)
	}

	return result
}


// func MFindOne (db string, table string, filter map[string]interface{}) map[string]interface{} {

// 	mongoURI := os.Getenv("MONGO_URI")

//     client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

// 	if err != nil {
//         log.Fatal(err)
//     }

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

// 	err = client.Connect(ctx)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer client.Disconnect(ctx)

// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection := client.Database(db).Collection(table)
	
// 	var result map[string]interface{}

// 	collection.FindOne(context.TODO(), filter).Decode(&result)

// 	return result
// }

// func MDeleteOne (db string, table string, filter map[string]interface{}) {

// 	mongoURI := os.Getenv("MONGO_URI")

//     client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

// 	if err != nil {
//         log.Fatal(err)
//     }

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

// 	err = client.Connect(ctx)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer client.Disconnect(ctx)

// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection := client.Database(db).Collection(table)
	
// 	_, errDelete := collection.DeleteOne(context.TODO(), filter)

// 	if errDelete != nil {
// 		log.Fatal(errDelete)
// 	}
// }

// func MInsert (db string, table string, data map[string]interface{}) interface{} {

// 	mongoURI := os.Getenv("MONGO_URI")

//     client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

// 	if err != nil {
//         log.Fatal(err)
//     }

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

// 	err = client.Connect(ctx)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer client.Disconnect(ctx)

// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection := client.Database(db).Collection(table)
	
// 	result, errInsert := collection.InsertOne(context.TODO(), data)

// 	if errInsert != nil {
// 		log.Fatal(errInsert)
// 	}

// 	return result.InsertedID
// }