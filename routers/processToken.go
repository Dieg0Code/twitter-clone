package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dieg0code/twitter-clone/db"
	"github.com/dieg0code/twitter-clone/models"
)

// Email value of the email used in every endpoint
var Email string

// UserID is the ID returned by the model, used in every endpoint
var UserID string

// ProcessToken process the token to get the data
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("mySecretKey")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, finded, _ := db.CheckIfUserExists(claims.Email)
		if finded {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, finded, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
