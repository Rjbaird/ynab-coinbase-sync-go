package ynab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"bairrya.com/go/ynab-coinbase/schema"
)

func GetAccountBalance(token string, bID string, aID string) (*float64, error) {
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
	balance := float64(account.Balance / 1000.0)
	return &balance, nil
}

func UpdateAccountBalance(token string, bID string, aID string, change float64) (*string, error){
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

	url := fmt.Sprintf("http://api.youneedabudget.com/v1/budgets/%s/transactions?access_token=%s", bID, token)
	currentTime := time.Now()
	formattedDate := currentTime.Format("2006-01-02")

	amount := fmt.Sprintf("%.0f", change)

	requestBody := strings.NewReader(`
		{
			"transaction": {
				"account_id": "` + aID + `",
				"date": "` + formattedDate + `",
				"amount": ` + amount + `,
				"payee_name": "Coinbase",
				"cleared": "cleared",
				"approved": true
			}
		}
	`)

	req, err := http.NewRequest(http.MethodPost, url, requestBody)
    if err != nil {
		fmt.Println("Error creating YNAB API request:", err)
        return nil, err
    }

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
    res, resErr := client.Do(req)

	var response schema.YNABResponse
	var errorResponse schema.YNABError
	if res.StatusCode != 201 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading YNAB API body response:", err)
			return nil, err
		}
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			fmt.Println("Error unmarshalling JSON from YNAB API response:", err)
			return nil, err
		}
		fmt.Println(res.StatusCode)
		fmt.Println("Error sending YNAB API request:", resErr)
		return nil, err
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error sending YNAB API response body:", err)
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON from YNAB API response:", err)
		return nil, err
	}
	return nil, nil
}