package schema

type YNABResponse struct {
	Data Data `json:"data"`
}
type DebtInterestRates struct {
}
type DebtMinimumPayments struct {
}
type DebtEscrowAmounts struct {
}
type Account struct {
	ID                  string              `json:"id"`
	Name                string              `json:"name"`
	Type                string              `json:"type"`
	OnBudget            bool                `json:"on_budget"`
	Closed              bool                `json:"closed"`
	Note                any                 `json:"note"`
	Balance             int64               `json:"balance"`
	ClearedBalance      int64               `json:"cleared_balance"`
	UnclearedBalance    int64               `json:"uncleared_balance"`
	TransferPayeeID     string              `json:"transfer_payee_id"`
	DirectImportLinked  bool                `json:"direct_import_linked"`
	DirectImportInError bool                `json:"direct_import_in_error"`
	LastReconciledAt    any                 `json:"last_reconciled_at"`
	DebtOriginalBalance any                 `json:"debt_original_balance"`
	DebtInterestRates   DebtInterestRates   `json:"debt_interest_rates"`
	DebtMinimumPayments DebtMinimumPayments `json:"debt_minimum_payments"`
	DebtEscrowAmounts   DebtEscrowAmounts   `json:"debt_escrow_amounts"`
	Deleted             bool                `json:"deleted"`
}
type Data struct {
	Account Account `json:"account"`
}