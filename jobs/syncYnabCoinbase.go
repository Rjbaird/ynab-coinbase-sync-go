package jobs

import (
	"fmt"
	"os"
	"sync"
	"time"

	"bairrya.com/go/ynab-coinbase/coinbase"
	"bairrya.com/go/ynab-coinbase/ynab"
	"github.com/joho/godotenv"
)

func SyncYnabCoinbase() {
	start := time.Now()
    defer func() {
        fmt.Println("ynab-coinbase sync execution Time: ", time.Since(start))
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
	
	if usd != old {
		fmt.Println("Updating YNAB balance...")
		up := (*cAcct - *yAcct) * 1000
		_, err := ynab.UpdateAccountBalance(ynabToken, budgetID, accountID, up)
		if err != nil {
			fmt.Println("Error updating ynab account balance:", err)
			return
		}
		fmt.Println("YNAB balance updated!")
	}	
}