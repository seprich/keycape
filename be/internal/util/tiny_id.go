package util

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/samber/lo"
)

func GenerateTinyId() string {
	bytes := make([]byte, 6)
	lo.Must(rand.Read(bytes))
	return base64.RawURLEncoding.EncodeToString(bytes)
}
