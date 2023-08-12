package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"bairrya.com/go/ynab-coinbase/coinbase"
	"bairrya.com/go/ynab-coinbase/ynab"
	"github.com/joho/godotenv"
)


func main() {
	start := time.Now()
    defer func() {
        fmt.Println("Execution Time: ", time.Since(start))
    }()
	godotenv.Load(".env")
	ynabToken := os.Getenv("YNAB_TOKEN")
	budgetID := os.Getenv("BUDGET_ID")
	accountID := os.Getenv("ACCOUNT_ID")
	var wg sync.WaitGroup
	wg.Add(2)

	var yAcct *float64
	var cAcct *float64
	
	go func() {
		coinStart := time.Now()
    	defer func() {
			fmt.Println("coinbase execution Time: ", time.Since(coinStart))
		}()
		wallet, err := coinbase.GetWalletData("")
		if err != nil {
			fmt.Println("Error getting wallet data:", err)
			return
		}
		
		balance, err := coinbase.GetAccountBalance(wallet)
		if err != nil {
			fmt.Println("Error getting account balance:", err)
			return
		}
		cAcct = balance
		wg.Done()
	}()
	
	go func() {
		ynabStart := time.Now()
    	defer func() {
			fmt.Println("ynab execution Time: ", time.Since(ynabStart))
		}()
		acct, err := ynab.GetAccountBalance(ynabToken, budgetID, accountID)
		if err != nil {
			fmt.Println("Error getting ynab account balance:", err)
			return
		}
		yAcct = acct
		wg.Done()
		
	}()
	

	wg.Wait()

	usd := fmt.Sprintf("%.0f", *cAcct)
	old := fmt.Sprintf("%.0f", *yAcct)
	coinbaseMessage := fmt.Sprintf("Coinbase Balance: $%v", usd)
	ynabMessage := fmt.Sprintf("YNAB Balance: $%v", old)
	
	fmt.Println(coinbaseMessage)
	fmt.Println(ynabMessage)

}