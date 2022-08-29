package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(str string) bool {
	if str == "" {
		return true
	}
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	processedString := strings.ToLower(re.ReplaceAllString(str, ""))
	return isP(processedString)
}

func isP(str string) bool {
	if len(str) == 0 || len(str) == 1 {
		return true
	}
	if str[0] == str[len(str)-1] {
		return isP(str[1 : len(str)-1])
	}
	return false
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
