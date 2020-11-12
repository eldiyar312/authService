package main

import (
	"authService/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main () {

    router := httprouter.New()

    router.POST("/api/token", Index)

    // cors enable
    handler := cors.Default().Handler(router)
    http.ListenAndServe(":8080", handler)
}


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    
    // allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"

    w.Header().Set("Access-Control-Allow-Origin", "*")
    // w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    // w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
    // w.Header().Set("Access-Control-Expose-Headers", "Authorization")

    respone := utils.Message(true, "hello")

    w.Header().Add("Content-Type", "application/json")

    utils.Respond(w, respone)
}