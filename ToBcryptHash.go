package strutils

import "golang.org/x/crypto/bcrypt"

// ToBcryptHash converts the string to bcrypt hash
func ToBcryptHash(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
