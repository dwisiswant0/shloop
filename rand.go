package main

import (
	"strconv"

	rand "github.com/brianvoe/gofakeit"
)

func rint(l uint) string {
	return rand.DigitN(l)
}

func rstr(l uint) string {
	return rand.LetterN(l)
}

func uuid() string {
	return rand.UUID()
}

func increment() string {
	i++

	return strconv.Itoa(i)
}
