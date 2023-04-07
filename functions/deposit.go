package functions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var LengthOfSeperator = 58

// This function will Credit ammount into perticular user's account
func Deposit(user *User) {
	var inputReader = bufio.NewReader(os.Stdin)
	typeOfTransaction := "Credit"
	openingBalance := user.Balance
	fmt.Println(strings.Repeat("-", LengthOfSeperator))
	fmt.Println("Enter amount you want to Deposite :")
	amtStr, _ := inputReader.ReadString('\n')
	amt, ok := AmmountToFloat(amtStr)
	if ok {
		if amt >= 0 {
			if amt > 20000 {
				fmt.Println(strings.Repeat("-", LengthOfSeperator))
				fmt.Println("Deposite ammount sholud be less than 20000 Rs.")
				fmt.Println(strings.Repeat("-", LengthOfSeperator))
			} else if amt >= 500 {
				if math.Mod(amt, 500) == 0 {
					user.Balance += amt
					CreateLog(user, typeOfTransaction, amt, openingBalance, user.Balance)
					fmt.Println(strings.Repeat("-", LengthOfSeperator))
					fmt.Println("Transaction Successful")
					ViewBalance(user)
				} else {
					fmt.Println(strings.Repeat("-", LengthOfSeperator))
					fmt.Println("Transaction Failed")
					fmt.Println("Amount sholud be multiple of 500")
					fmt.Println(strings.Repeat("-", LengthOfSeperator))
				}
			} else {
				fmt.Println(strings.Repeat("-", LengthOfSeperator))
				fmt.Println("Transaction Failed")
				fmt.Println("You have to Deposite Minimum 500 Rs.")
				fmt.Println(strings.Repeat("-", LengthOfSeperator))
			}
		} else {
			fmt.Println(strings.Repeat("-", LengthOfSeperator))
			fmt.Println("Transaction Failed")
			fmt.Println("Invalid Amount")
			fmt.Println(strings.Repeat("-", LengthOfSeperator))
		}
	}
	if IsContinue() {
		return
	} else {
		os.Exit(0)
	}
}
