package currency

// Currency struct
type Currency struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"fee_currency"`
}

// GetNewAccessToken returns AccessToken
func GetNewCurrency(symbol string) *Currency {
	return &Currency{ID: symbol}
}
