package models

type Item struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number" faker:"len=13"`
	Price       int    `json:"price" faker:"boundary_start=50, boundary_end=10000"`
	Rid         string `json:"rid" faker:"len=20"`
	Name        string `json:"name" faker:"name"`
	Sale        int    `json:"sale" faker:"boundary_start=0, boundary_end=99"`
	Size        string `json:"size" faker:"oneof: 0, 1, 2, 3, 4, 5, 6"`
	TotalPrice  int    `json:"total_price" faker:"boundary_start=50, boundary_end=10000"`
	NmId        int    `json:"nm_id" faker:"boundary_start=1, boundary_end=10000"`
	Brand       string `json:"brand" faker:"oneof: Moncler, Prada"`
	Status      int    `json:"status" faker:"boundary_start=200, boundary_end=500"`
}
