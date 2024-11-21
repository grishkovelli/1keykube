package utils

import (
	"time"
)

func Bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}

func Lifetime(days int) int64 {
	if days == 0 {
		return 0
	}

	now := time.Now()
	return now.AddDate(0, 0, days).Unix()
}
