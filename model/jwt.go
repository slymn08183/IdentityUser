package model

import "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	Email    string `json:"email" bson:"email"`
	UserName string `json:"userName" bson:"user_name"`
	Uid      string `json:"uid" bson:"uid"`
	jwt.StandardClaims
}
