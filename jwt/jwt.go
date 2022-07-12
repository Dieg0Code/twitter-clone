package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dieg0code/twitter-clone/models"
)

func GenerateJWT(usr models.User) (string, error) {
	myKey := []byte("mySecretKey")

	payload := jwt.MapClaims{
		"email":     usr.Email,
		"name":      usr.Name,
		"lastname":  usr.LastName,
		"birthDate": usr.BirthDate,
		"bio":       usr.Bio,
		"location":  usr.Location,
		"webSite":   usr.WebSite,
		"_id":       usr.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
