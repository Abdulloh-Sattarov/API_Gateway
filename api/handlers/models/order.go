package models

type Order struct {
	OrderId     string `json:"OrderId"`
	BookId      string `json:"BookId"`
	BookName    string `json:"BookName"`
	AuthorId    string `json:"AuthorId"`
	AuthorName  string `json:"AuthorName"`
	Description string `json:"Description"`
	CreatedAt   string `json:"CreatedAt"`
	UpdatedAt   string `json:"UpdatedAt"`
}

type UpdateOrder struct {
	OrderId     string `json:"OrderId"`
	BookId      string `json:"BookId"`
	Description string `json:"Description"`
}

type CreateOrder struct {
	BookId      string `json:"BookId"`
	Description string `json:"Description"`
}

type ListOrders struct {
	Orders []Order `json:"Orders"`
}
