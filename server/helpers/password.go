package helpers

import "golang.org/x/crypto/bcrypt"

func HassingPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordHash), err
}

func VerifyPassword(hashPasword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPasword), []byte(password))
}
