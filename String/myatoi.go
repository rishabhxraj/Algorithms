package main

import (
	"fmt"
)

// func myAtoi(str string) int {
// 	if len(str) == 0 {
// 		return 0
// 	}
// 	i := 0
// 	for (i < len(str)) && (str[i] == ' ') {
// 		i++
// 	}
// 	str = str[i:]

// 	sign := +1
// 	var ans int64 = 0

// 	if str[0] == '-' {
// 		sign = -1
// 	}

// 	j := 0
// 	if (str[0] == '+') || (str[0] == '-') {
// 		j = 1
// 	}
// 	for j < len(str) {
// 		if str[i] == ' ' && !unicode.IsDigit(rune(str[i])) {
// 			break
// 		}
// 		ans = ans*10 + str[i] - '0'
// 	}

// 	return 0
// }

func main() {
	str := '4' - '0'
	fmt.Println(int(str))
}
