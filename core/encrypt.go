package core

import (
	"crypto/rand"
	"fmt"
	"os"

	"golang.org/x/crypto/nacl/secretbox"
)

func EncryptFile(srcPath string, c *Config) error {
	key := deriveKey(c.Password)

	plain, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return err
	}

	encrypted := secretbox.Seal(nonce[:], plain, &nonce, key)

	dstPath := fmt.Sprintf("%s.enc", srcPath)
	return os.WriteFile(dstPath, encrypted, 0600)
}
