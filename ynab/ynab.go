package ynab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bairrya.com/go/ynab-coinbase/schema"
)

func GetAccountBalance(token string, bID string, aID string) (*int64, error) {
	if token == "" {
		fmt.Println("YNAB_TOKEN is not set")
		return nil, fmt.Errorf("YNAB_TOKEN is not set")
	}
	if bID == "" {
		fmt.Println("BUDGET_ID is not set")
		return nil, fmt.Errorf("BUDGET_ID is not set")
	}
	if aID == "" {
		fmt.Println("ACCOUNT_ID is not set")
		return nil, fmt.Errorf("ACCOUNT_ID is not set")
	}

	url := fmt.Sprintf("http://api.youneedabudget.com/v1/budgets/%s/accounts/%s?access_token=%s", bID, aID, token)
	// auth := fmt.Sprintf("Bearer %s", token)
    
	req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
		fmt.Println("Error creating YNAB API request:", err)
        return nil, err
    }

    // req.Header.Add("Authorization", auth)
    // req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
    res, err := client.Do(req)

	var response schema.YNABResponse
	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
		fmt.Println("Error sending YNAB API request:", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error sending YNAB API request:", err)
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON from YNAB API response:", err)
		return nil, err
	}
	account := schema.Budget{ Balance: response.Data.Account.Balance }
	balance := account.Balance / 1000.0
	fmt.Println(balance)
	return &balance, nil
}