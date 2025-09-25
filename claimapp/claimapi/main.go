package main

import (
	"net/http"

	"github.com/eswaribala/claimapp/claimapi/store"
)

// @title Claim API
// @version 1.0
// @description This is api service for managing Claims
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email parameswaribala@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7072
// @BasePath /
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /claims/v1.0", store.GetClaims)
	mux.HandleFunc("POST /claims/v1.0", store.SaveClaim)
	mux.HandleFunc("GET /claims/v1.0/{id}", store.GetClaimByID)
	mux.HandleFunc("PUT /claims/v1.0/{id}", store.UpdateClaim)
	mux.HandleFunc("DELETE /claims/v1.0/{id}", store.DeleteClaim)
	http.ListenAndServe(":7072", mux)

}
