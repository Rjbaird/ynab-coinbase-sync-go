package coinbase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"bairrya.com/go/ynab-coinbase/schema"
	"github.com/joho/godotenv"
)

func GetWalletData(route string) (*[]schema.WalletData, error) {
	walletStart := time.Now()
	defer func() {
		fmt.Println("coinbase wallet execution time: ", time.Since(walletStart))
	}()
	if route == "" {
		route = "/v2/accounts?&limit=100"
	}
	godotenv.Load(".env")
	key := os.Getenv("COINBASE_API_KEY")
	if key == "" {
		fmt.Println("COINBASE_API_KEY is not set")
		return nil, fmt.Errorf("COINBASE_API_KEY is not set")
	}

	url := "https://api.coinbase.com"
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
		return nil, fmt.Errorf("error sending wallet request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error getting wallet data:", resp.Status)
		return nil, fmt.Errorf("error getting wallet data: %v - %v", resp.Status, resp.Body.Close().Error())
	}

	var walletResp schema.WalletResponse
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

	if walletResp.Pagination.NextURI != "" {
		nextWalletData, err := GetWalletData(walletResp.Pagination.NextURI)
		if err != nil {
			fmt.Println("Error getting next wallet data:", err)
			return nil, err
		}
		walletResp.Data = append(walletResp.Data, *nextWalletData...)
	}
	walletData := walletResp.Data
	return &walletData, nil
}

func GetAccountBalance(walletData *[]schema.WalletData) (*float64, error) {
	coinStart := time.Now()
	defer func() {
		fmt.Println("coin rate execution time: ", time.Since(coinStart))
	}()
	var acctBalance float64 = 0
	for _, account := range *walletData {
		if !strings.Contains(account.Balance.Amount, "0.0") {
			coinBalance, err := strconv.ParseFloat(account.Balance.Amount, 64)
			if err != nil {
				fmt.Println("Error parsing balance:", err)
			}
			rate, err := getCoinRate(account.Balance.Currency)
			if err != nil {
				fmt.Println("Error getting coin rate:", err)
			}
			usd := coinBalance * *rate
			acctBalance += usd
		}
	}
	return &acctBalance, nil
}

func getCoinRate(coin string) (*float64, error) {
	godotenv.Load(".env")
	key := os.Getenv("EXCHANGE_KEY")
	if key == "" {
		fmt.Println("EXCHANGE_KEY is not set")
		return nil, fmt.Errorf("EXCHANGE_KEY is not set")
	}

	url := "https://rest.coinapi.io/v1/exchangerate/"
	method := http.MethodGet

	req, err := http.NewRequest(method, url + coin, nil)
	if err != nil {
		fmt.Println("Error creating wallet request:", err)
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CoinAPI-Key", key)
	
	// Create Client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending wallet request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error getting wallet data:", resp.Status)
		return nil, fmt.Errorf("error getting wallet data: %v - %v", resp.Status, resp.Body.Close().Error())
	}

	var data schema.CoinRate
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading wallet response:", err)
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON from Coinbase API response:", err)
		return nil, err
	}

	var rate float64
	for _, coin := range data.Rates {
		if coin.AssetIDQuote == "USD" {
			rate = coin.Rate
		}
	}
	return &rate, nil
}