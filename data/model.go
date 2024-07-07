package model

type Bill struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Balance float64 `json:"balance"`
}

var Bills = []Bill{
	{ID: "1", Name: "Isa", Surname: "Isaev", Balance: 52.52},
	{ID: "2", Name: "Suleyma", Surname: "Suleymanov", Balance: 77.77},
	{ID: "3", Name: "Mag", Surname: "Magov", Balance: 14.88},
}
