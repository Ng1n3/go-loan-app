package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"loan/account"
	"loan/fileops"
	"github.com/Pallinder/go-randomdata"
)

func Choice() {
    var choice int
    for {
        List()
        fmt.Print("\nEnter your choice: ")
        _, err := fmt.Scan(&choice)
        if err != nil {
            fmt.Println("Invalid input. Please enter a number from 1 to 6")
            // Clear the input buffer
            bufio.NewReader(os.Stdin).ReadString('\n')
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
                fmt.Printf("Error: %v\n", err)
                continue
            }

            userAccount.DisplayAccountNumber()

            err = fileops.WriteToFile(userAccount)
            if err != nil {
                fmt.Printf("Error saving account to file: %v\n", err)
            } else {
                fmt.Println("Account saved to file successfully")
            }

        case 2:
            accountNumber, err := getAccountDetails("Enter the account holder's account number to check loan status: ")
            if err != nil {
                fmt.Printf("Error getting account number: %v\n", err)
                continue
            }

            acc, err := fileops.CheckAccountByAccountNumber(accountNumber)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                continue
            }

            acc.CheckLoanStatus()
          case 3:
            accountNumber, err := getAccountDetails("Enter the account holder's account number to close: ")
            if err != nil {
              fmt.Printf("Error geting account number: %v\n", err)
              continue
            }

            loanAmountStr, err := getAccountDetails("Enter the amount of loan neede: ")
            if err != nil {
              fmt.Printf("error getting the required loan amount %v\n", err)
              continue
            }

            loanAMount, err := strconv.ParseFloat(loanAmountStr, 64)
            if err != nil || loanAMount <= 0 {
              fmt.Println("Invalid amount enterd")
              continue
            }

            err = fileops.GiveLoan(accountNumber, loanAMount)
            if err != nil {
              fmt.Printf("error: %v\n", err)
            }
        case 4:
            accountNumber, err := getAccountDetails("Enter the account holder's account number to close: ")
            if err != nil {
                fmt.Printf("Error getting account number: %v\n", err)
                continue
            }

            repaymentAmountStr, err := getAccountDetails("How much do you want to pay: ")
            if err != nil {
                fmt.Printf("error getting the required loan amount %v\n", err)
                continue
            }

            repaymentAmount, err := strconv.ParseFloat(repaymentAmountStr, 64)
            if err != nil || repaymentAmount <= 0 {
                fmt.Println("Invalid amount entered")
                continue
            }

            err = fileops.RepayLoan(accountNumber, repaymentAmount)
            if err != nil {
                fmt.Printf("error: %v\n", err)
            }


        case 5:
            accountNumber, err := getAccountDetails("Enter the account holder's account number to close: ")
            if err != nil {
                fmt.Printf("Error getting account number: %v\n", err)
                continue
            }

            err = fileops.DeleteAccountbyAccountNumber(accountNumber)

            if err != nil {
              fmt.Printf("Error: %v\n", err)
              continue
            }

        case 6:
            firstName, err := getAccountDetails("Enter your firstname for your account: ")
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                continue
            }

            _, err = fileops.GetAcountNumberByFirstName(firstName)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                continue
            }

        case 7:
            fmt.Printf("Goodbye!\nWe can't wait to see you again, Go Loan Bank.\n")
            return

        default:
            fmt.Println("Sorry please choose from the menu listed 1 to 7.")
        }
    }
}

func List() {
    fmt.Println("----- MENU -------")
    fmt.Printf("\n1) Create an Account\n2) Check Loan Status\n3) Get a Loan\n4) Repay Loan\n5) Close an Account\n6) Check account number\n7) Exit\n")
    fmt.Println("------------------")
}

func getAccountDetails(prompt string) (string, error) {
    fmt.Print(prompt)

    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')
    if err != nil {
        return "", fmt.Errorf("unable to get user input: %w", err)
    }

    text = strings.TrimSpace(text)
    if text == "" {
        return "", errors.New("input cannot be empty")
    }

    return text, nil
}

func getAccountData() (string, string, string, string, error) {
    fmt.Println("Please Enter the details required below to create your Go Loan Account")
    
    firstName, err := getAccountDetails("First name: ")
    if err != nil {
        return "", "", "", "", fmt.Errorf("error getting first name: %w", err)
    }

    lastName, err := getAccountDetails("Last name: ")
    if err != nil {
        return "", "", "", "", fmt.Errorf("error getting last name: %w", err)
    }

    phoneNumber, err := getAccountDetails("Phone number: ")
    if err != nil {
        return "", "", "", "", fmt.Errorf("error getting phone number: %w", err)
    }

    accountNumber := randomdata.StringNumber(4, "")

    // Convert to lowercase
    firstName = strings.ToLower(firstName)
    lastName = strings.ToLower(lastName)

    err = account.CheckPhoneNumber(phoneNumber)
    if err != nil {
        return "", "", "", "", fmt.Errorf("invalid phone number: %w", err)
    }

    return firstName, lastName, phoneNumber, accountNumber, nil
}