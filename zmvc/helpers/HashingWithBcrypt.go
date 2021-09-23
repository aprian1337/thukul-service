package helpers

import "golang.org/x/crypto/bcrypt"

func HashWithBcrypt(password string) string {
	strToByte := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(strToByte, bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	return string(hash)
}

func CompareHashWithBcrypt(hashPassword string, password string) bool {
	passToByte := []byte(password)
	hashToByte := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hashToByte, passToByte)
	if err != nil {
		return false
	}
	return true
}
