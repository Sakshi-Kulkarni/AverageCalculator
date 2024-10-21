package utils

import (
  	"strings"
)

func IsValidInput(input string) bool {
	nums := strings.Fields(input)

	for _, numStr := range nums {
		if !isDigits(numStr) {
			return false
		}
	}
	return true
}

func isDigits(s string) bool {
	for _, char := range s {

		if char < '0' || char > '9' {
 			return false
		}
 	}
	return true
}

//gerics and string 