package main

import (
	"github.com/authService/token"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main () {
    
    godotenv.Load() // Загрузить файл .env
    var (
        port = os.Getenv("PORT")
        originAllowed = []string{os.Getenv("ORIGIN_ALLOWED")}
    )

    router := mux.NewRouter()
    http.Handle("/", router)


    // ROUTES
    router.HandleFunc("/api/token", token.GenerateTokens).Methods("POST")


    // CORS
    handler := cors.Default().Handler(router)
    
    handler = cors.New(cors.Options{
        AllowedOrigins: originAllowed,
        AllowCredentials: true,
    }).Handler(router)


    // START
    fmt.Println(port)
    err := http.ListenAndServe(":" + port, handler)

	if err != nil {
		fmt.Print(err)
	}
}