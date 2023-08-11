package coinbase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func CreateSignature(accessTimestamp int64, requestPath string) (*string, error) {
	godotenv.Load(".env")
	
	// accessPassphrase := os.Getenv("ACCESS_PASSPHRASE")
	secret := os.Getenv("COINBASE_SECRET")
		if secret == "" {
		fmt.Println("COINBASE_SECRET is not set")
		return nil, fmt.Errorf("COINBASE_SECRET is not set")
	}
	
	method := "GET"
	message := fmt.Sprintf("%v%s%s%s", accessTimestamp, method, requestPath, "")

	cbAccessSign := encodeMessage(secret, message)

	return &cbAccessSign, nil
}

func encodeMessage(secret string, message string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}