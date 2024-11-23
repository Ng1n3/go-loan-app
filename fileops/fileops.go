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

func WriteToFile(newAccountInfo account.Account) error {
	var accounts []account.Account

	existingAccountFile, err := os.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(existingAccountFile, &accounts)

		if err != nil {
			return errors.New("failed to parse existing account")
		}
	}

	accounts = append(accounts, newAccountInfo)

	file, err := os.Create(filename)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	err = encoder.Encode(accounts)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}
	return nil
	
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

func DeleteAccountbyAccountNumber(accountNumber string) error {
  data, err := os.ReadFile(filename)

	if err != nil {
		return  err
	}

	// declare a slice to hold all accounts
	var accounts []account.Account

	// Parse the JSON data into the slice
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return  err
	}

	var updatedAccounts []account.Account
	accountFound := false

	for _, acc := range accounts {
		if acc.AccountNumber == accountNumber {
			accountFound = true
			continue
		}
		updatedAccounts = append(updatedAccounts, acc)
	}

	if !accountFound {
		return errors.New("account not found")
	}

	// convert the updated slice back to JSON
	updatedData, err := json.MarshalIndent(updatedAccounts, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, updatedData, 0644)
	if err != nil {
		return err
	}

	//if no record of account is found, return an error
	return nil
}