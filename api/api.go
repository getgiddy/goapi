package api

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success code, usually 200
	Code int

	// Account balance
	Balance int64
}

// Error Response
type ErrorResponse struct {
	// Error code, usually 400
	Code int

	// Error message
	Message string
}
