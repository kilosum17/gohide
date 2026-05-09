package utils

import (
	"strings"
)

const EncExtension = ".enc"

func CreateEncPath(srcPath string) string {
	return srcPath + EncExtension
}

func CreateDecPath(srcPath string) string {
	if before, ok := strings.CutSuffix(srcPath, EncExtension); ok {
		return before
	}
	return srcPath
}
