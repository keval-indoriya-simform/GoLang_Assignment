package functions

import (
	"fmt"
	"strconv"
	"strings"
)

// This function will convert given string to int
func StrToNum(str string) (int, bool) {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	num, err := strconv.Atoi(str)
	if err == nil {
		return num, true
	}
	fmt.Println(strings.Repeat("-", LengthOfSeperator))
	fmt.Println("You Entered Something else but only numbers are required")
	fmt.Println(strings.Repeat("-", LengthOfSeperator))
	return 0, false
}
