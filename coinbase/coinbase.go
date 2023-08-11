package coinbase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"bairrya.com/go/ynab-coinbase/schema"
	"github.com/joho/godotenv"
)

func GetWalletData() (*[]schema.WalletData, error) {
	godotenv.Load(".env")
	key := os.Getenv("COINBASE_API_KEY")
	if key == "" {
		fmt.Println("COINBASE_API_KEY is not set")
		return nil, fmt.Errorf("COINBASE_API_KEY is not set")
	}
	url := "https://api.coinbase.com"
	route := "/v2/accounts?&limit=100"
	method := http.MethodGet
	currentTime := time.Now().UTC()
	timestamp := currentTime.Unix()

	signature, err := CreateSignature(timestamp, route)

	if err != nil {
		fmt.Println("Error creating wallet signature:", err)
		return nil, err
	}

	req, err := http.NewRequest(method, url + route, nil)
	if err != nil {
		fmt.Println("Error creating wallet request:", err)
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("CB-ACCESS-KEY", key)
	req.Header.Set("CB-ACCESS-SIGN", *signature)
	req.Header.Set("CB-ACCESS-TIMESTAMP", fmt.Sprintf("%v", timestamp))
	req.Header.Set("CB-VERSION", "2021-02-13")
	
	// Create Client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending wallet request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	// Send Request
	var resp404 schema.Coinbase404
	if resp.StatusCode != http.StatusOK {
		fmt.Println("wallet response Status:", resp.Status)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil, err
		}
		err = json.Unmarshal(body, &resp404)
		fmt.Println("coinbase response:", resp404)
		if err != nil {
			fmt.Println("Error unmarshalling JSON from Coinbase API response:", err)
			return nil, err
		}
		return nil, err
	}

	var walletResp schema.WalletResponse
	// Read Response Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading wallet response:", err)
		return nil, err
	}

	err = json.Unmarshal(body, &walletResp)
	if err != nil {
		fmt.Println("Error unmarshalling JSON from Coinbase API response:", err)
		return nil, err
	}
	fmt.Println("wallet response:", walletResp.Pagination.NextURI)
	walletData := walletResp.Data
	return &walletData, nil
}

func GetAccountBalance(walletData *[]schema.WalletData) {
	fmt.Println("Getting account balance")
	for _, account := range *walletData {
		if account.Balance.Amount != "0.00" {
			fmt.Println(account.Name, ":", account.Balance.Amount, account.Balance.Currency)
		}
	}
}
