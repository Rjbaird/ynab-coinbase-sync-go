package main

import (
	"os"

	"bairrya.com/go/ynab-coinbase/ynab"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load(".env")
	ynabToken := os.Getenv("YNAB_TOKEN")
	budgetID := os.Getenv("BUDGET_ID")
	accountID := os.Getenv("ACCOUNT_ID")

	ynab.GetAccountBalance(ynabToken, budgetID, accountID)
}