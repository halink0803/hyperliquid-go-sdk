package hyperliquid

// AllMidsResponse ...
type AllMidsResponse map[string]string

// AllMids - retrieve mids for all actively traded coins
func (c *Client) AllMids() *AllMidsResponse {
	var (
		response AllMidsResponse
	)
	// TODO:
	return &response
}

// OpenOrder struct
type OpenOrder struct {
	Coin      string `json:"coin"`
	LimitPx   string `json:"limitPx"`
	OID       uint64 `json:"oid"`
	Side      string `json:"side"`
	Size      string `json:"sz"`
	Timestamp string `json:"timestamp"`
}

// GetOpenOrders - return user open order
// requestBody:
// type: "openOrders"
// user: Address in 42-character hexadecimal format
func (c *Client) GetOpenOrders() []OpenOrder {
	var (
		response []OpenOrder
	)
	return response
}

// OrderFills return user order fills
// RequestBody:
// type: "userFills"
// user: Address in 42-character hexadecimal format
func (c *Client) OrderFills() {

}
