package common

import (
	"fmt"
	"regexp"

	"golang.org/x/exp/slices"
)

// ColorSliceStringGreen takes in a slice and a string. It will find the string
// in the slice and color it green.
func ColorSliceStringGreen(slice []string, item string) []string {
	currentIndex := slices.IndexFunc(slice, func(value string) bool { return value == item })
	slice[currentIndex] = fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, item)
	fmt.Println(slice[currentIndex])

	return slice
}

// RemoveGreenStringformatting will take the green formatting out of the string.
func RemoveGreenStringFormatting(s string) string {
	// Regular expression for the specific green ANSI color code
	re := regexp.MustCompile("\x1b\\[32m(.*?)\x1b\\[0m")
	return re.ReplaceAllString(s, "$1")
}
