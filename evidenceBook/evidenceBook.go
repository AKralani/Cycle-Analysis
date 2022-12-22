package evidenceBook

// Evidence Book
type EvidenceBook struct {
	EvidenceBookID     *int    `json:"evidenceBookId"`
	Date               string  `json:"date"`
	Time               string  `json:"time"`
	Coin               string  `json:"coin"`
	Price              float64 `json:"price"`
	USDValue           float64 `json:"usdValue"`
	Quantity           float64 `json:"quantity"`
	Fee                float64 `json:"fee"`
	BuySell            string  `json:"buySell"`
	ProfitLoss         string  `json:"profitLoss"`
	PercentProfitLoss  float64 `json:"percentProfitLoss"`
	USDValueProfitLoss float64 `json:"usdValueProfitLoss"`
	Snapshot           string  `json:"snapshot"`
}
