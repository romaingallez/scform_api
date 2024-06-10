package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Write a countdown functuion that will countdown from a given number and allow to pass time.Duration

func Countdown(count int, duration time.Duration) {

	for i := 1; i < count; i++ {
		fmt.Println(i)
		time.Sleep(duration)
	}
}

func RemoveExtraSpaces(input string) string {
	// Utiliser des expressions régulières pour remplacer les espaces multiples par un seul espace
	re := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(re.ReplaceAllString(input, " "))
}
