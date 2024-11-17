package fileops

import (
	"encoding/json"
	"os"
	"example.com/loan/account"
)

func WriteToFile(account account.Account) error {
  filename := "account.json"

  json, err := json.Marshal(account)

  if err != nil {
    return err
  }

  return os.WriteFile(filename, json, 0644)
}