package password

import (
	"math"
	"math/rand"
	"strings"
)

type Spec struct {
	Length    int
	Uppercase int
	Numbers   int
	Symbols   int
}

type SymSet string

func (s SymSet) Rand() byte {
	return s[rand.Intn(len(s))]
}

func (s SymSet) Contains(c rune) bool {
	return strings.ContainsRune(string(s), c)
}

const (
	uppercase SymSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase SymSet = "abcdefghijklmnopqrstuvwxyz"
	numbers   SymSet = "0123456789"
	special   SymSet = "!@#$%^&*()_+-=[]{}|;':\",./<>?`~"
)

func Generate(s *Spec) string {
	str := make([]byte, s.Length)

	for i := 0; i < s.Length; i++ {
		switch {
		case s.Uppercase > 0:
			str[i] = uppercase.Rand()
			s.Uppercase--
		case s.Numbers > 0:
			str[i] = numbers.Rand()
			s.Numbers--
		case s.Symbols > 0:
			str[i] = special.Rand()
			s.Symbols--
		default:
			str[i] = lowercase.Rand()
		}
	}

	rand.Shuffle(len(str), func(i, j int) {
		str[i], str[j] = str[j], str[i]
	})

	return string(str[:])
}

func SymbolCount(p string) map[string]int {
	count := map[string]int{
		"uppercase": 0,
		"lowercase": 0,
		"numbers":   0,
		"special":   0,
	}

	for _, c := range p {
		switch {
		case uppercase.Contains(c):
			count["uppercase"]++
		case numbers.Contains(c):
			count["numbers"]++
		case special.Contains(c):
			count["special"]++
		case lowercase.Contains(c):
			count["lowercase"]++
		}
	}

	return count
}

func AlphabetSize(p string) int {
	count := SymbolCount(p)
	size := 0

	if count["uppercase"] > 0 {
		size += 26
	}
	if count["lowercase"] > 0 {
		size += 26
	}
	if count["numbers"] > 0 {
		size += 10
	}
	if count["special"] > 0 {
		size += 32
	}

	return size
}

func Entropy(p string) float64 {
	return float64(len(p)) * math.Log2(float64(AlphabetSize(p)))
}

func ClasifyEntropy(e float64) string {
	switch {
	case e < 36:
		return "very weak"
	case e < 60:
		return "weak"
	case e < 59:
		return "good"
	case e < 120:
		return "strong"
	default:
		return "very strong"
	}
}
