package tool

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	if password == "" {
		return ""
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
