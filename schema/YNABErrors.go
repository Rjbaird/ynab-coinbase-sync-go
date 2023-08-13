package schema

type YNABError struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}