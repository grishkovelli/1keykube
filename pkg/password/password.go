package password

import (
	"math/rand"
)

type Spec struct {
	Length    int
	Uppercase int
	Numbers   int
	Symbols   int
}

const (
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	numbers   = "0123456789"
	symbols   = "!@#$%^&*()_+-=[]{}|;':\",./<>?"
)

func Generate(s *Spec) string {
	str := make([]byte, s.Length)

	for i := 0; i < s.Length; i++ {
		switch {
		case s.Uppercase > 0:
			str[i] = uppercase[rand.Intn(len(uppercase))]
			s.Uppercase--
		case s.Numbers > 0:
			str[i] = numbers[rand.Intn(len(numbers))]
			s.Numbers--
		case s.Symbols > 0:
			str[i] = symbols[rand.Intn(len(symbols))]
			s.Symbols--
		default:
			str[i] = lowercase[rand.Intn(len(lowercase))]
		}
	}

	rand.Shuffle(len(str), func(i, j int) {
		str[i], str[j] = str[j], str[i]
	})

	return string(str[:])
}
