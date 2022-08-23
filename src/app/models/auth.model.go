package models

type JWTClaims struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
