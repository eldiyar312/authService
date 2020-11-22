package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/authService/token/route/auth"
	"github.com/authService/token/route/pages"

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
    router.HandleFunc("/", pages.RespondDoc).Methods("GET")
    router.HandleFunc("/api/token", auth.TGenerateById).Methods("POST")
    router.HandleFunc("/api/refresh", auth.TGenerateByrefreshT).Methods("POST")
    router.HandleFunc("/api/delete/refresh", auth.DeleteRefreshT).Methods("POST")
    router.HandleFunc("/api/delete/all/refresh", auth.DeleteTokens).Methods("POST")


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