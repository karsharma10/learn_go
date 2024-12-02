package config

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

type HashingKey struct {
	KeyPhrase string
}

func MdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:])
}

func ShaHashing(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}
