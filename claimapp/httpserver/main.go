package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

func handleRequest(prefix string) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		resp := map[string]string{
			"message": fmt.Sprintf("%s %s", prefix, gofakeit.HackerPhrase()),
		}
		json.NewEncoder(response).Encode(resp)
	}
}

func handlePolicyRequest() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		resp := map[string]any{
			"id":         "P1001",
			"holder":     "Asha Rao",
			"type":       "HEALTH",
			"coverage":   1000000,
			"deductible": 5000,
		}
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(resp)
	}
}

func main() {
	http.HandleFunc("/", handleRequest("Admin says: "))
	http.HandleFunc("/policies", handlePolicyRequest())
	http.HandleFunc("/customers", handleRequest("Customer: "))
	http.HandleFunc("/claims", handleRequest("Claim: "))
	fmt.Println("Server started on port 7074")

	http.ListenAndServe(":7074", nil)

}
