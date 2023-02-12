package models

type Payment struct {
	Transaction  string `json:"transaction" faker:"len=13"`
	RequestId    string `json:"request_id" faker:"len=19"`
	Currency     string `json:"currency" faker:"currency"`
	Provider     string `json:"provider" faker:"oneof: cc, paypal, check, money order"`
	Amount       int    `json:"amount" faker:"boundary_start=1, boundary_end=1000"`
	PaymentDt    int    `json:"payment_dt" faker:"boundary_start=1, boundary_end=999999"`
	Bank         string `json:"bank" faker:"oneof: SBER, POCHTA, TIN, VTB, ALPHA"`
	DeliveryCost int    `json:"delivery_cost" faker:"boundary_start=100, boundary_end=1000"`
	GoodsTotal   int    `json:"goods_total" faker:"boundary_start=1, boundary_end=1000"`
	CustomFee    int    `json:"custom_fee" faker:"boundary_start=0, boundary_end=10"`
}
