package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dieg0code/twitter-clone/models"
)

// GenerateJWT generates a JWT token
func GenerateJWT(u models.User) (string, error) {
	myKey := []byte("mySecretKey")

	payload := jwt.MapClaims{
		"email":     u.Email,
		"name":      u.Name,
		"lastname":  u.LastName,
		"birthDate": u.BirthDate,
		"bio":       u.Bio,
		"location":  u.Location,
		"webSite":   u.WebSite,
		"_id":       u.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
