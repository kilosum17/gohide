package core

import (
	"crypto/sha256"
)

func deriveKey(password string) *[32]byte {
	hash := sha256.Sum256([]byte(password))
	var key [32]byte
	copy(key[:], hash[:])
	return &key
}

type Options struct {
	Mode     string
	Src      string
	Dest     string
	Password string
}

func Consum(args ...any) {
	// Cosume args
}
