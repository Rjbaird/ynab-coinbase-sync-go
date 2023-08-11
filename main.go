package main

import (
	"fmt"

	"bairrya.com/go/ynab-coinbase/coinbase"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load(".env")
	// ynabToken := os.Getenv("YNAB_TOKEN")
	// budgetID := os.Getenv("BUDGET_ID")
	// accountID := os.Getenv("ACCOUNT_ID")

	// ynab.GetAccountBalance(ynabToken, budgetID, accountID)
	wallet, err := coinbase.GetWalletData()
	if err != nil {
		fmt.Println("Error getting wallet data:", err)
		return
	}
	coinbase.GetAccountBalance(wallet)
}