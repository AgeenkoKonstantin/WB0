package models

type Order struct {
	OrderUID          string `json:"order_uid" faker:"uuid_digit"`
	TrackNumber       string `json:"track_number" faker:"len=14"`
	Entry             string `json:"entry" faker:"oneof: WBIL, WBIL1, WBIL2"`
	Locale            string `json:"locale" faker:"oneof: en, ru"`
	InternalSignature string `json:"internal_signature" faker:"-"`
	CustomerId        string `json:"customer_id" faker:"len=10"`
	DeliveryService   string `json:"delivery_service" faker:"oneof: meest, someone_service, dev_service"`
	Shardkey          string `json:"shardkey" faker:"len=10"`
	SmId              int    `json:"sm_id" faker:"boundary_start=1, boundary_end=1000"`
	DateCreated       string `json:"date_created" faker:"timestamp"`
	OofShard          string `json:"oof_shard" faker:"len=2"`

	Delivery `json:"delivery"`
	Payment  `json:"payment"`
	Items    []Item `json:"items"`
}
