package twitter

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

type Authentication struct{}

func NewAuthentication() *Authentication {
	return &Authentication{}
}

func (*Authentication) AuthNonce() string {
	key := make([]byte, 32)
	rand.Read(key)

	s := base64.StdEncoding.EncodeToString(key)
	for _, v := range []string{"+", "/", "="} {
		s = strings.ReplaceAll(s, v, "")
	}
	return s
}
