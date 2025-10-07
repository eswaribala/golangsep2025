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

func TestGetClaims(t *testing.T) {
	router := SetUpRouter()
	req, _ := http.NewRequest("GET", "/claims", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var returnedClaims []Claim
	err := json.Unmarshal(w.Body.Bytes(), &returnedClaims)
	assert.Nil(t, err)
	assert.Equal(t, claims, returnedClaims)
}

func TestFetchClaim(t *testing.T) {
	m := mocha.New(t).CloseOnCleanup(t)
	m.Start()

	claim := Claim{ClaimID: "12345", Amount: 1001.50, Description: "Test Claim", Status: true}

	m.AddMocks(mocha.
		Get(expect.URLPath("/claims/12345")).
		Reply(reply.OK().Status(
			http.StatusOK,
		).BodyJSON(claim)))
	claim, _ = FetchClaim(m.URL())

	//assert.NoError(t, err)
	assert.Equal(t, "12345", claim.ClaimID)
	assert.Equal(t, 1000.50, claim.Amount)
	assert.Equal(t, "Test Claim", claim.Description)
	assert.Equal(t, true, claim.Status)
}

func TestCreateClaim(t *testing.T) {
	m := mocha.New(t).CloseOnCleanup(t)
	m.Start()

	claim := Claim{ClaimID: "12348", Amount: 1000, Description: "Accident", Status: true}

	m.AddMocks(mocha.
		Post(expect.URLPath("/claims")).
		Reply(reply.OK().Status(
			http.StatusCreated,
		).BodyJSON(claim)))
	got, code, err := CreateClaimClient(m.URL(), claim)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, code)
	assert.Equal(t, claim, got)
	assert.Contains(t, claim.ClaimID, "12345")
	//assert.Contains(t, claim.Amount, 1000)
}
