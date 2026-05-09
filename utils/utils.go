package utils

import (
	"errors"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
)

func ReadPassword(mes string) string {
	fmt.Print(mes)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
		panic("Could not read!")
	}
	fmt.Println()
	password := string(bytePassword)
	return password
}

func GetPassword() (string, error) {
	password := ""
	tries := 0
	for {
		password = ReadPassword("Enter Password: ")
		if len(password) < 3 {
			fmt.Println("Short Password: Enter 3+ characters")
			continue
		}
		password2 := ReadPassword("Confirm Password: ")
		if password == password2 {
			return password, nil
		} else {
			fmt.Println("Passwords din't match")
		}
		tries++
		if tries > 3 {
			return "", errors.New("Getting password failed")
		}
	}
}

func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	} else {
		return b
	}
}
