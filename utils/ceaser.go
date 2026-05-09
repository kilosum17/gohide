package utils

import (
	"path/filepath"
	"strings"
)

const (
	ObfPrefix = "!!" // Marker to identify an obfuscated filename
)

func shiftFileName(fullPath string, password string, encrypt bool) string {
	dir, name := filepath.Split(fullPath)
	key := PasswordToKey(password)

	if !encrypt {
		key = -key
		name = strings.TrimPrefix(name, ObfPrefix)
	}

	runes := []rune(name)
	for i := range runes {
		runes[i] = rune(int(runes[i]) + key)
	}

	finalName := string(runes)

	if encrypt {
		finalName = ObfPrefix + finalName
	}

	return filepath.Join(dir, finalName)
}

func GetEncryptPath(filePath string, password string, shouldChange bool) string {
	// Only obfuscate the name if requested
	if shouldChange {
		filePath = shiftFileName(filePath, password, true)
	}
	// Always add .enc so the tool knows it's an encrypted file
	return CreateEncPath(filePath)
}

func GetDecryptPath(filePath string, password string) string {
	// 1. Remove the .enc extension first
	cleanPath := CreateDecPath(filePath)

	// 2. Extract the filename to check for the obfuscation marker
	_, name := filepath.Split(cleanPath)

	// 3. If the name starts with our marker, reverse the shift
	if strings.HasPrefix(name, ObfPrefix) {
		return shiftFileName(cleanPath, password, false)
	}

	// 4. Otherwise, just return the path without .enc
	return cleanPath
}
