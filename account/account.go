package account

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Account struct {
	Firstname           string    `json:"firstname"`
	Lastname            string    `json:"lastname"`
	PhoneNumber         string    `json:"phone_number"`
	AccountNumber       string    `json:"account_number"`
	LoanStatus          bool      `json:"loan_status"`
	LoanAmountAvailable float64   `json:"loan_amount_available"`
	NumberOfLoans       int       `json:"number_of_loans"`
	CurrentLoan         float64   `json:"current_loan"`
	CreatedAt           time.Time `json:"created_at"`
}

func New(firstName, lastName, phoneNumber, accountNumber string) (Account, error) {

	if firstName == "" || lastName == "" || phoneNumber == "" {
		return Account{}, errors.New("this field  cannot be empty")
	}

	return Account{
		Firstname:           firstName,
		Lastname:            lastName,
		PhoneNumber:         phoneNumber,
		AccountNumber:       accountNumber,
		LoanStatus:          false,
		LoanAmountAvailable: 5000.0,
		CurrentLoan:         0.0,
		NumberOfLoans:       0,
		CreatedAt:           time.Now(),
	}, nil
}

func (account Account) DisplayAccountNumber() {
	// accountNumber := randomdata.StringNumber(4, "")
	fmt.Printf("Hey %v welcome once again to Go Loan Bank.\nYour phone number %v.\nHere is you account number: %v\n", account.Firstname, account.PhoneNumber, account.AccountNumber)
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

func (account Account) CheckLoanStatus() {
	if !account.LoanStatus {
		loanStatus := "You have no pending loans"
		fmt.Printf("\n Hey %v below is your current loan status\nLoan status: %v \nLoan amount available: %v\n", account.Firstname, loanStatus, account.LoanAmountAvailable)
	} else {
		loanStatus := "You have some pending loans, please pay us as soon as you can to enjoy better offers next time"
		fmt.Printf("\n Hey %v below is your current loan status\nLoan status: %v \nCurrent loan amount: %v\n", account.Firstname, loanStatus, account.CurrentLoan)
	}
}

