package models

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateUrl() string {
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	log.Println("GENERATING URL...")
	return base64.URLEncoding.EncodeToString(randomBytes)[:6]

}
