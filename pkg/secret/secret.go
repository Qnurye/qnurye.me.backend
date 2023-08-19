package secret

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

type Secret struct {
	Jwt      string
	Database struct {
		Key string
		IV  string
	}
	Client struct {
		Key string
		IV  string
	}
}

var s Secret

func Generate() *Secret {
	jwtKey := randRead(32)
	dAesKey := randRead(32)
	cAesKey := randRead(32)
	dIv := randRead(16)
	cIv := randRead(16)

	s.Jwt = base64.StdEncoding.EncodeToString(jwtKey)
	s.Database.Key = base64.StdEncoding.EncodeToString(dAesKey)
	s.Database.IV = base64.StdEncoding.EncodeToString(dIv)
	s.Client.Key = base64.StdEncoding.EncodeToString(cAesKey)
	s.Client.IV = base64.StdEncoding.EncodeToString(cIv)

	return &s
}

func randRead(len int) []byte {
	var err error

	b := make([]byte, len)
	_, err = rand.Read(b)
	if err != nil {
		log.Fatalln("Error on generating b: ", err)
	}
	return b
}
