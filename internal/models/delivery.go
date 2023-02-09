package models

type Delivery struct {
	Name    string `json:"name"  faker:"name"`
	Phone   string `json:"phone" faker:"phone_number"`
	Zip     string `json:"zip" faker:"len=20"`
	City    string `json:"city" faker:"oneof: moscow, piter, othercity"`
	Address string `json:"address" faker:"oneof: Ploshad Mira 15, ulitsa butlerova 10,ulitsa  1905 goda"`
	Region  string `json:"region" faker:"word"`
	Email   string `json:"email" faker:"email"`
}
