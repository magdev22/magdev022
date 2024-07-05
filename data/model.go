package model

type Product struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Product string  `json:"product"`
	Price   float64 `json:"price"`
}

var Products = []Product{
	{ID: "1", Title: "kanfetka", Product: "chi orio", Price: 52.52},
	{ID: "2", Title: "moloko", Product: "ne moloko", Price: 77.77},
	{ID: "3", Title: "chokoladka", Product: "alenka", Price: 66.66},
}
