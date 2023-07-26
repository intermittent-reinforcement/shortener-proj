package app

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/itchyny/base58-go"
)

const idLen int8 = 8

// Returns SHA256 hash of user input URL string
func hashURL(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hash.Sum(nil)
}

// Encoding hashed users input to Base58
// Input must be a []byte encoded string representing decimal number
func toBase58(decNum []byte) string {
	ripple := base58.RippleEncoding
	encoded, err := ripple.Encode(decNum)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// Returns first idLen charecters of pre-hashed URL Base58 string
func GenerateID(origURL string) string {
	// Get SHA256 sum of user input URL string
	hashedURL := hashURL(origURL)
	// Convert hash bytes to string decimal number
	bigInt := fmt.Sprintf("%d", binary.BigEndian.Uint64(hashedURL))
	// Encode hashed URL to Base64 string and return idLen of first characters
	return toBase58([]byte(bigInt))[:idLen]
}
