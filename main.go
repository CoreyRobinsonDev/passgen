package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
)

var (
	PASSWORD_LENGTH = 16
	LOW_CHARS = "abcdefghijklmnopqrstuvwxyz"
	UPP_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIG_CHARS = "0123456789"
	SYM_CHARS = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)


func main() {
	homeDir := Unwrap(os.UserHomeDir())
	fileName := homeDir + "/.passwords"
	file := Unwrap(os.OpenFile(
		fileName, 
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	))
	defer file.Close()

	password := ""

	for range PASSWORD_LENGTH {
		charArr := [4]string{LOW_CHARS, UPP_CHARS, DIG_CHARS, SYM_CHARS}
		charArrIdx := rand.Int() % 4
		charsIdx := rand.Int() % len(charArr[charArrIdx])
		password += string(charArr[charArrIdx][charsIdx])
	}
	fmt.Print(password)
}

func Expect(err error) {
	if err != nil {
		slog.Error(fmt.Sprintf("passgen: %v", err))
		os.Exit(1)
	}
}

func Unwrap[T any](val T, err error) T {
	if err != nil {
		slog.Error(fmt.Sprintf("passgen: %v", err))
		os.Exit(1)
	}
	return val
}
