package models

type Delivery struct {
	Name    string `json:"name"  faker:"name, len=15"`
	Phone   string `json:"phone" faker:"phone_number, len=15"`
	Zip     string `json:"zip" faker:"len=18"`
	City    string `json:"city" faker:"oneof: moscow, piter, othercity"`
	Address string `json:"address" faker:"oneof: Pl 15, butlerova 10, 1905 goda"`
	Region  string `json:"region" faker:"word, len=15"`
	Email   string `json:"email" faker:"email, len=20"`
}
