package schema

import "time"

type Wallet struct {
	Name string
	Balance float64
	AssetID string
}

type WalletResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []WalletData     `json:"data"`
	Warnings   []Warnings `json:"warnings"`
}
type Pagination struct {
	EndingBefore         any    `json:"ending_before"`
	StartingAfter        any    `json:"starting_after"`
	PreviousEndingBefore any    `json:"previous_ending_before"`
	NextStartingAfter    string `json:"next_starting_after"`
	Limit                int    `json:"limit"`
	Order                string `json:"order"`
	PreviousURI          any    `json:"previous_uri"`
	NextURI              string `json:"next_uri"`
}
type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	SortIndex    int    `json:"sort_index"`
	Exponent     int    `json:"exponent"`
	Type         string `json:"type"`
	AddressRegex string `json:"address_regex"`
	AssetID      string `json:"asset_id"`
	Slug         string `json:"slug"`
}
type Balance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
type Rewards struct {
	Apy          string `json:"apy"`
	FormattedApy string `json:"formatted_apy"`
	Label        string `json:"label"`
}
type Currency0 struct {
	Code                string `json:"code"`
	Name                string `json:"name"`
	Color               string `json:"color"`
	SortIndex           int    `json:"sort_index"`
	Exponent            int    `json:"exponent"`
	Type                string `json:"type"`
	AddressRegex        string `json:"address_regex"`
	AssetID             string `json:"asset_id"`
	DestinationTagName  string `json:"destination_tag_name"`
	DestinationTagRegex string `json:"destination_tag_regex"`
	Slug                string `json:"slug"`
}
type Currency1 struct {
	Code                string `json:"code"`
	Name                string `json:"name"`
	Color               string `json:"color"`
	SortIndex           int    `json:"sort_index"`
	Exponent            int    `json:"exponent"`
	Type                string `json:"type"`
	AddressRegex        string `json:"address_regex"`
	AssetID             string `json:"asset_id"`
	DestinationTagName  string `json:"destination_tag_name"`
	DestinationTagRegex string `json:"destination_tag_regex"`
	Slug                string `json:"slug"`
}
type WalletData struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Primary          bool      `json:"primary"`
	Type             string    `json:"type"`
	Currency         Currency  `json:"currency,omitempty"`
	Balance          Balance   `json:"balance"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Resource         string    `json:"resource"`
	ResourcePath     string    `json:"resource_path"`
	AllowDeposits    bool      `json:"allow_deposits"`
	AllowWithdrawals bool      `json:"allow_withdrawals"`
	Rewards          Rewards   `json:"rewards,omitempty"`
}
type Warnings struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	URL     string `json:"url"`
}