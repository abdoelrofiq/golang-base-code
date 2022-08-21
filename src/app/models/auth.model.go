package models

type JWTClaims struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
