package tools

import (
	"encoding/hex"
	"fmt"
	"regexp"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/scrypt"
)

const xaalAddrPattern = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"

// IsValidAddr : Check is the xAAL address is valid
func IsValidAddr(val string) bool {
	if val == "" {
		return false
	}
	re := regexp.MustCompile(xaalAddrPattern)
	fmt.Println(val)
	fmt.Println(re.MatchString(val))
	return re.MatchString(val)
}

// GetRandomUUID : Generates a new xAAL UUID
func GetRandomUUID() string {
	u1 := uuid.Must(uuid.NewV1()) // panic on error
	return u1.String()
}

// Pass2key : Generates key from passphrase using scrypt
func Pass2key(passphrase string) string {
	salt := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // buffer of zeros (crypto_pwhash_scryptsalsa208sha256_SALTBYTES)
	key, _ := scrypt.Key([]byte(passphrase), []byte(salt), 16384, 8, 1, 32)
	return hex.EncodeToString(key)
}
