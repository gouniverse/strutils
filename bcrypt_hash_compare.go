package strutils

import "golang.org/x/crypto/bcrypt"

// BcryptHashCompare compares the string to a bcrypt hash
func BcryptHashCompare(str string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))

	return err == nil
}
