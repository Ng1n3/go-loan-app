package fileops

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/loan/account"
)

const filename = "account.json"

func WriteToFile(newAccount account.Account) error {

	data, err := os.ReadFile(filename)
	if err != nil {
		//If the file doesn't exist, create a new one with an array
		if os.IsNotExist(err) {
			accounts := []account.Account{newAccount}
			jsonData, err := json.MarshalIndent(accounts, "", " ")
			if err != nil {
				return err
			}
			return os.WriteFile(filename, jsonData, 0644)
		}
		return err
	}

	//unMarshal existing data into a slice
	var accounts []account.Account
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return err
	}

	accounts = append(accounts, newAccount)

	// Marshal the updated slice back to JSON
	jsonData, err := json.MarshalIndent(accounts, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func CheckAccountByAccountNumber(accountNumber string) (*account.Account, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	// declare a slice to hold all accounts
	var accounts []account.Account

	// Parse the JSON data into the slice
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return nil, err
	}

	for _, acc := range accounts {
		if acc.AccountNumber == accountNumber {
			// return all matching account
			return &acc, nil
		}
	}

	//if no record of account is found, return an error
	return nil, errors.New("account not found")
}

func GetAcountNumberByFirstName(firstName string) (*account.Account, error) {
	firstName = strings.ToLower(firstName)

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	// declare a slice to hold all accounts
	var accounts []account.Account

	// Parse the JSON data into the slice
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return nil, err
	}

	for _, acc := range accounts {
		if acc.Firstname == firstName {
			fmt.Printf("Your account number is %v\n", acc.AccountNumber)
			return &acc, nil
		}
	}

	return nil, errors.New("account not found")

}
