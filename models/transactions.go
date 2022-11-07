package transactions

type Transactions struct {
	Name     string `json:"name"`
	Value    int64  `json:"value"`
	Fees     int64  `json:"fees"`
	Quantity int    `json:"quantity"`
}
