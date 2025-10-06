package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Claim struct {
	ClaimID     string
	Amount      float64
	Description string
	Status      bool
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

func GetClaims(c *gin.Context) {
	c.JSON(http.StatusOK, claims)
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/claims", AddClaim)
	r.GET("/claims", GetClaims)
	return r
}

func TestRouter(url string) *gin.Engine {
	r := gin.Default()
	r.POST("/claims", AddClaim)
	return r
}
