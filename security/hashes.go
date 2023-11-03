package security

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func Hash(content string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(content), -1)
	return base64.StdEncoding.EncodeToString(result)
}

func CompareHash(hash, original string) bool {
	raw, _ := base64.StdEncoding.DecodeString(hash)
	return bcrypt.CompareHashAndPassword(raw, []byte(original)) == nil
}
