package main

import (
	"errors"
	"fmt"
	"strconv"
  "github.com/Pallinder/go-randomdata"
)

const accountFile = "account.txt"

func main() {
	fmt.Println("Welcome to Go Loan Bank")
	var choice int
  for {
    Menu()
    fmt.Print("\nEnter your choice: ")
    _,err := fmt.Scan(&choice)
    if err != nil {
      fmt.Println("Invalid input. Please enter a number from 1 to 6")
      continue
    }
    switch choice {
    case 1:
      CreateAccount()
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

func CreateAccount() {
  var name, phoneNumber string

  fmt.Print("Enter your name: ")
  fmt.Scan(&name)
  fmt.Print("Enter your Your phone number: ")
  fmt.Scan(&phoneNumber)
  err := CheckPhoneNumber(phoneNumber)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }

  accountNumber := randomdata.StringNumber(4, "")
  fmt.Printf("Hey %v welcome once again to Go Loan Bank.\nYour phone number %v.\nHere is you account number: %v\n", name, phoneNumber, accountNumber)

}

func CheckPhoneNumber(phoneNumber string) error {
  if len(phoneNumber) != 10 {
    return errors.New("phone number must contain exactly 10 digits")
  } 

  _, err := strconv.Atoi(phoneNumber)
  if err != nil {
    return errors.New("phone number must contain only digits")
  }

  return nil
}

func WriteToFile() {

}

func ReadFromFile() {

}
