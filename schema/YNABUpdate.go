package schema

type TransactionData struct {
	Transaction Transaction `json:"transaction"`
}

type Transaction struct {
	AccountID   string  `json:"account_id"`
	Date        string  `json:"date"`
	Amount      float64 `json:"amount"`
	PayeeName   string  `json:"payee_name"`
	Cleared     string  `json:"cleared"`
	Approved    bool    `json:"approved"`
}