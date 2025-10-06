package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ========= Models ========= */

type Role struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Username  string               `bson:"username" json:"username"`
	Password  string               `bson:"password" json:"-"` // hash; never return
	Email     string               `bson:"email,omitempty" json:"email,omitempty"`
	RoleIDs   []primitive.ObjectID `bson:"role_ids" json:"role_ids"`
	CreatedAt time.Time            `bson:"created_at" json:"created_at"`
}

type UserWithRoles struct {
	ID        primitive.ObjectID `json:"id"`
	Username  string             `json:"username"`
	Email     string             `json:"email,omitempty"`
	Roles     []Role             `json:"roles"`
	CreatedAt time.Time          `json:"created_at"`
}

/* ========= DTOs ========= */

type RegisterRequest struct {
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"` // e.g. ["admin","manager"]
}

// --- API payloads ---

type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"admin"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token" example:"<jwt>"`
}

type APIError struct {
	Message string `json:"message" example:"invalid credentials"`
}

type Profile struct {
	Username string `json:"username" example:"admin"`
	Role     string `json:"role" example:"demo-admin"`
	Issuer   string `json:"issuer" example:"jwtdemo"`
	Expires  string `json:"expires" example:"2025-09-26T13:00:00Z"`
}

// --- JWT claims ---

// RegisteredClaims mirrors jwt.RegisteredClaims but as an embedded type for examples.
type RegisteredClaims = jwt.RegisteredClaims

type TokenClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	RegisteredClaims
}

func TimePtr(t time.Time) *jwt.NumericDate {
	return jwt.NewNumericDate(t)
}

