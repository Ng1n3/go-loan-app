package account

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type Account struct {
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	PhoneNumber   string    `json:"phone_number"`
	AccountNumber string    `json:"account_number"`
	CreatedAt     time.Time `json:"created_at"`
}

func New(firstName, lastName, phoneNumber, accountNumber string) (Account, error) {

	if firstName == "" || lastName == "" || phoneNumber == "" {
		return Account{}, errors.New("this field  cannot be empty")
	}

	return Account{
		Firstname:     firstName,
		Lastname:      lastName,
		PhoneNumber:   phoneNumber,
		AccountNumber: accountNumber,
		CreatedAt:     time.Now(),
	}, nil
}

func (account Account) DisplayAccountNumber() {
	accountNumber := randomdata.StringNumber(4, "")
	fmt.Printf("Hey %v welcome once again to Go Loan Bank.\nYour phone number %v.\nHere is you account number: %v\n", account.Firstname, account.PhoneNumber, accountNumber)
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
