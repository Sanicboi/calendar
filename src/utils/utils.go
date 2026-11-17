package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	env "github.com/joho/godotenv"
)

func LoadEnv() {
	if err := env.Load(); err != nil {
		log.Fatal(".env file not found")
	}
}

func CreateToken(userId int64) (string, error) {
	key, present := os.LookupEnv("JWT_KEY")
	if present == false {
		log.Fatal("JWT key not present")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
	})
	res, err := t.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
		return res, fmt.Errorf("Could not sign token")
	}
	return res, nil

}

func ParseToken(token string) (int64, error) {

	type CustomClaims struct {
		Id int64 `json:"id"`
		jwt.RegisteredClaims
	}

	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		key, present := os.LookupEnv("JWT_KEY")
		if present == false {
			return nil, fmt.Errorf("Key not present")
		}
		return []byte(key), nil
	})

	if err != nil {
		fmt.Println(err)
		return -1, fmt.Errorf("Could not parse token")
	}

	if claims, ok := t.Claims.(*CustomClaims); ok {
		return claims.Id, nil
	}

	return -1, fmt.Errorf("Wrong claims type")
}
