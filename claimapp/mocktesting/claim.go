package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Claim struct {
	ClaimID     string
	Amount      float64
	Description string
	Status      bool
}
type Client struct {
	BaseURL string
	HTTP    *http.Client
}

var claims []Claim

func AddClaim(c *gin.Context) {
	var newClaim Claim
	if err := c.ShouldBindJSON(&newClaim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims = append(claims, newClaim)
	c.JSON(http.StatusCreated, newClaim)

}
func CreateClaimClient(baseURL string, c Claim) (Claim, int, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return Claim{}, 0, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/claims", baseURL), "application/json", bytes.NewReader(b))
	if err != nil {
		return Claim{}, 0, err
	}
	defer resp.Body.Close()

	var out Claim
	_ = json.NewDecoder(resp.Body).Decode(&out)
	return out, resp.StatusCode, nil
}
func FetchClaim(apiURL string) (Claim, error) {
	resp, err := http.Get(apiURL + "/claims/12345")
	if err != nil {
		return Claim{}, err
	}
	defer resp.Body.Close()

	var claim Claim
	err = json.NewDecoder(resp.Body).Decode(&claim)
	return claim, err
}

func GetClaims(c *gin.Context) {
	c.JSON(http.StatusOK, claims)
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/claims", AddClaim)
	r.GET("/claims", GetClaims)
	return r
}
func TestRouter(api string) *gin.Engine {
	r := gin.Default()
	r.POST("/claims/v1.0", AddClaim)

	return r
}
