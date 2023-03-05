package tool

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyClaim struct {
	EmployeeID      int 
	jwt.StandardClaims
}

func TokenCreate(employeeID int) string {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaim{
		EmployeeID:      employeeID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	signedStr, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return signedStr
}

func TokenValidate(t string) (*jwt.Token, error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	token, _ := jwt.ParseWithClaims(t, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error while decoding the token")
		}

		return jwtKey, nil
	})

	return token, nil
}
