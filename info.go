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

// OrderFillsByTime return user order fills by time
// RequestBody:
// type: "userFillsByTime"
// user: Address in 42-character hexadecimal format
// startTime: start time in milliseconds, inclusive
// endTime: end time in milliseconds, inclusive. Defaults to current time
func (c *Client) OrderFillsByTime() {

}

// QueryUserRateLimit return user rate limits
// RequestBody:
// user: address in 42-character hexadecimal format
// type: "userRateLimit"
func (c *Client) QueryUserRateLimit() {

}

// OrderStatus return order status by oid or cloid
// RequestBody:
// user: address in 42-character hexadecimal format
// type: "orderStatus"
// oid: Either u64 representing the orderID or 16-byte hex string representing the client order id
func (c *Client) OrderStatus() {

}

// L2BookSnapshot return l2 order book snapshot
// RequestBody:
// type: "l2Book"
// coin: coin
// nSigFigs: (optional) field to aggregate levels to {nSigFigs} significant figures.
//
//	Accepts values between 2 and 5
//
// mantissa: (optional) field to aggregate levelss. This field is only allowed
//
//	if nSigFigs is 5. Accepts values of 1, 2 or 5
func (c *Client) L2BookSnapshot() {

}

// CandleSnapshot return the most recent 5000 candles available
// RequestBody:
// type: "candleSnapshot"
// req: {"coin": <coin>, "interval": "15m", "startTime": <epoch millis>, "endTime": <epoch millis>}
func (c *Client) CandleSnapshot() {

}

// CheckBuilderFeeApproval return maximum fee approved
// Header:
// Content-Type: "application/json"
// RequestBody:
// type: "maxBuilderFee"
// user: address in 42-character hexadecimal format
// builder: address in 42-character hexadecimal format
func (c *Client) CheckBuilderFeeApproval() {

}
