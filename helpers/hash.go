package helpers

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashTransactionToSlug(uuid string, keyString string, keyAdditional string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(uuid))
	encodeEncrypt := base64.StdEncoding.EncodeToString([]byte(uuid + "+" + keyAdditional))
	return fmt.Sprintf("%s/%s", encode, AesEncrypt(encodeEncrypt, keyString))
}
