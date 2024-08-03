package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"pos-api/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJwt(UserID string, Username string, Fullname string, RoleID string, PrefixID string, Component []string) (string, error) {
	var privateKey = config.GetEnvConfig("SECRET_KEY")
	claims := jwt.MapClaims{
		"user_id":    UserID,
		"username":   Username,
		"fullname":   Fullname,
		"role_id":    RoleID,
		"components": Component,
		"prefix_id":  PrefixID,
		"exp":        time.Now().Add(time.Hour * 10).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return "Fail", err
	}
	return t, nil
}

func CreateTwoFactor() string {
	randomBytes := make([]byte, 25)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	base32Encoded := base32.StdEncoding.EncodeToString(randomBytes)
	return base32Encoded
}

func Sign2fa(secret string) (string, error) {
	decodedSecret, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", fmt.Errorf("error decoding secret: %v", err)
	}
	counter := time.Now().Unix() / 30
	buffer := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		buffer[i] = byte(counter & 0xff)
		counter >>= 8
	}
	hmacHash := hmac.New(sha1.New, decodedSecret)
	hmacHash.Write(buffer)
	hmacResult := hmacHash.Sum(nil)
	code := dynamicTruncationFn(hmacResult)
	return fmt.Sprintf("%06d", code%1000000), nil
}

func dynamicTruncationFn(hmacResult []byte) uint32 {
	offset := hmacResult[len(hmacResult)-1] & 0x0F
	truncatedHash := hmacResult[offset : offset+4]
	code := uint32(truncatedHash[0]&0x7F)<<24 |
		uint32(truncatedHash[1])<<16 |
		uint32(truncatedHash[2])<<8 |
		uint32(truncatedHash[3])

	return code
}
