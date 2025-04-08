package utils

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func randomStr(length int) (string, error) {
	var b = make([]byte, length)

	_, err := rand.Read(b)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return hex.EncodeToString(b), nil
}

func HashPassword(password string) (string, error) {
	//salt, _ := randomStr(16)
	//saltPass := fmt.Sprintf("%s.%s", salt, password)
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func CheckPasswordHash(password, hash string) error {
	//// Split salt and actual hash
	//var salt, hash string
	//n, _ := fmt.Sscanf(fullHash, "%[^$]$%s", &salt, &hash)
	//if n != 2 {
	//	return fmt.Errorf("invalid hash format")
	//}
	//saltedPassword := fmt.Sprintf("%s.%s", salt, password)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return err
	}
	return nil
}
