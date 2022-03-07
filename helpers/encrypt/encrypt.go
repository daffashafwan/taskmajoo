package encrypt

import (
	"crypto/md5"
	"encoding/hex"

)

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func Compare(password, hash string) bool {
	return (hash == GetMD5Hash(password))
}
