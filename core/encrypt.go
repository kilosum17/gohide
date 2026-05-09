package core

import (
	"crypto/rand"
	"os"

	"com.gosafe/utils"
	"golang.org/x/crypto/nacl/secretbox"
)

func EncryptFile(filePath string, c *Config) error {
	key := utils.DeriveKey(c.Password)

	plain, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return err
	}

	encrypted := secretbox.Seal(nonce[:], plain, &nonce, key)

	dstPath := utils.GetEncryptPath(filePath, c.Password, c.Rename)
	err = os.WriteFile(dstPath, encrypted, 0600)
	if err != nil {
		return err
	}
	if !c.KeepOriginal {
		return os.Remove(filePath)
	}
	return nil
}
