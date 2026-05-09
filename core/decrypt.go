package core

import (
	"errors"
	"os"

	"com.gosafe/utils"
	"golang.org/x/crypto/nacl/secretbox"
)

func DecryptFile(filePath string, c *Config) error {
	key := deriveKey(c.Password)

	ciphertext, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if len(ciphertext) < 24 {
		return errors.New("invalid file")
	}

	var nonce [24]byte
	copy(nonce[:], ciphertext[:24])

	decrypted, ok := secretbox.Open(nil, ciphertext[24:], &nonce, key)
	if !ok {
		return errors.New("wrong password or corrupt file")
	}

	dstPath := utils.CreateDecPath(filePath)
	return os.WriteFile(dstPath, decrypted, 0600)
}
