package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/loan/account"
	"example.com/loan/fileops"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("Welcome to Go Loan Bank")
	var choice int
	for {
		Menu()
		fmt.Print("\nEnter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number from 1 to 6")
			continue
		}
		switch choice {
		case 1:
			firstName, lastName, phoneNumber, accountNumber, err := getAccountData()

			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			userAccount, err := account.New(firstName, lastName, phoneNumber, accountNumber)

			if err != nil {
				fmt.Printf("Error: %v", err)
				return
			}

			userAccount.DisplayAccountNumber()

			err = fileops.WriteToFile(userAccount)

			if err != nil {
				fmt.Printf("Error saving account to file: %v\n", err)
			} else {
				fmt.Println("Account saved to file successfully")
			}

		case 2:
			fmt.Println("Sorry currently unavailable")
		case 7:
			fmt.Printf("Goodbye!\nWe can't wait to see you again, Go Loan Bank.")
			return
		default:
			fmt.Println("Sorry please choose from the menu listed 1 to 7.")
		}
	}
}

func Menu() {
	fmt.Println("----- MENU -------")
	fmt.Printf("\n1) Create an Account\n2) Check Loan Status\n3) Get a Loan\n4) Repay Loan\n5) Close an Account\n6) Check account number\n7) Exit\n")
	fmt.Println("------------------")
}

func getAccountDetails(prompt string) string {
	fmt.Printf("%v", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getAccountData() (string, string, string, string, error) {
	fmt.Println("Please Enter the details required below to create your Go Loan Account")
	firstName := getAccountDetails("First name: ")
	lastName := getAccountDetails("Last name: ")
	phoneNumber := getAccountDetails("Phone number: ")
	accountNumber := randomdata.StringNumber(4, "")

	err := account.CheckPhoneNumber(phoneNumber)

	if err != nil {

		return "", "", "", "", fmt.Errorf("invalid phone number: %v", err)
	}

	return firstName, lastName, phoneNumber, accountNumber, nil
}
