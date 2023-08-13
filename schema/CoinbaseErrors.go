package schema

type CoinbaseError struct {
	Errors []Errors `json:"errors"`
}
type Errors struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	URL     string `json:"url"`
}