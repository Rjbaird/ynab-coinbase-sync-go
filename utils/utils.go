package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetExchangeRate() {
	godotenv.Load(".env")
	exchangeKey := os.Getenv("EXCHANGE_KEY")
	headers := fmt.Sprintf("X-CoinAPI-Key: %v", exchangeKey)
	fmt.Println(exchangeKey)
	fmt.Println(headers)
	
}