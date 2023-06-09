package functions

import (
	"fmt"
	"strings"
)

// This function will display the balance of that perticular user in consol
// user is passed by reference
func ViewBalance(user *User) {
	fmt.Println(strings.Repeat("-", LengthOfSeperator))
	fmt.Printf("Your Balance is : %.2f\n", user.Balance)
	fmt.Println(strings.Repeat("-", LengthOfSeperator))
}
