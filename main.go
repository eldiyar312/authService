package main

import (
	refresh "authService/route/TGenerateByRefreshT"
	delete "authService/route/deleteRefreshT"
	deleteTokens "authService/route/deleteTokens"
	"fmt"
	"log"
	"net/http"
	"os"

	tokens "github.com/authService/token/route/TGenerateById"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main () {
    // load .env file
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
    var (
        port = os.Getenv("PORT")
        originAllowed = []string{os.Getenv("ORIGIN_ALLOWED")}
    )

    router := mux.NewRouter()
    http.Handle("/", router)


    // ROUTES
    router.HandleFunc("/api/token", tokens.GenerateTokens).Methods("POST")
    router.HandleFunc("/api/refresh", refresh.Refreshing).Methods("POST")
    router.HandleFunc("/api/delete/refresh", delete.DeleteRefreshT).Methods("POST")
    router.HandleFunc("/api/delete/all/refresh", deleteTokens.DeleteTokens).Methods("POST")


    // CORS
    handler := cors.Default().Handler(router)
    
    handler = cors.New(cors.Options{
        AllowedOrigins: originAllowed,
        AllowCredentials: true,
    }).Handler(router)


    // START
    fmt.Println(port)
    log := http.ListenAndServe(":" + port, handler)

	if log != nil {
		fmt.Print(log)
	}
}