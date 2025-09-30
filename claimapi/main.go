package main

import (
	"log"
	"net/http"

	_ "github.com/cog/claimapi/docs"
	"github.com/cog/claimapi/store"
	httpSwagger "github.com/swaggo/http-swagger"
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

	res := store.EnsureTopic()
	if res != nil {
		log.Fatalf("failed to ensure kafka topic: %v", res)
	}
	println("Kafka topic ensured")
	store.InitKafkaWriters()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /claims/v1.0", store.GetClaims)
	mux.HandleFunc("POST /claims/v1.0", store.SaveClaim)
	mux.HandleFunc("GET /claims/v1.0/{claimid}", store.GetClaimByID)
	mux.HandleFunc("GET /claims/v1.0/kafka/{claimid}", store.PublishClaimInfoByID)
	mux.HandleFunc("PUT /claims/v1.0/{claimid}", store.UpdateClaim)
	mux.HandleFunc("DELETE /claims/v1.0/{claimid}", store.DeleteClaim)
	// Swagger UI served at /swagger/
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Your own handlers
	// mux.HandleFunc("/claims", claimsHandler)

	log.Println("Server running at http://localhost:7072")
	log.Fatal(http.ListenAndServe(":7072", mux))

}
