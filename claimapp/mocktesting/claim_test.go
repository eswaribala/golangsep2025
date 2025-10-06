package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorsalgado/mocha/v3"
	"github.com/vitorsalgado/mocha/v3/expect"
	"github.com/vitorsalgado/mocha/v3/reply"
)

func TestAddClaim(t *testing.T) {
	router := SetUpRouter()
	claim := Claim{
		ClaimID:     "12345",
		Amount:      1000.50,
		Description: "Test Claim",
		Status:      true,
	}
	jsonValue, _ := json.Marshal(claim)
	req, _ := http.NewRequest("POST", "/claims", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	var returnedClaim Claim
	err := json.Unmarshal(w.Body.Bytes(), &returnedClaim)
	assert.Nil(t, err)
	assert.Equal(t, claim, returnedClaim)
}

func TestAddClaimInvalidJSON(t *testing.T) {
	router := SetUpRouter()
	invalidJSON := `{"ClaimID": "12345", "Amount": "invalid_amount", "Description": "Test Claim", "Status": true}`
	req, _ := http.NewRequest("POST", "/claims", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Contains(t, response["error"], "json: cannot unmarshal")
}

func TestClaimAPI_ExternalMock(t *testing.T) {
	m := mocha.New(t)
	defer m.Close()

	m.AddMocks(mocha.Post(expect.URLPath("/claims")).
		Header("Content-Type", expect.ToEqual("application/json")).
		Body(expect.ToEqual(`{"ClaimID":"12345","Amount":1000.5,"Description":"Test Claim","Status":true}`)).
		Reply(reply.Created()))
	claim := Claim{
		ClaimID:     "12345",
		Amount:      1000.50,
		Description: "Test Claim",
		Status:      true,
	}
	router := TestRouter(m.URL())
	jsonValue, _ := json.Marshal(claim)
	req, _ := http.NewRequest("POST", "/claims", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	var returnedClaim Claim
	err := json.Unmarshal(w.Body.Bytes(), &returnedClaim)
	assert.Nil(t, err)
	assert.Equal(t, claim, returnedClaim)

}
