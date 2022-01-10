package models

type Book struct {
	BookId       string   `json:"BookId"`
	Name         string   `json:"Name"`
	AuthorId     string   `json:"AuthorId"`
	Price        float64  `json:"Price"`
	CategoryId   []string `json:"CategoryId"`
	CategoryName string   `json:"CategoryName"`
	CreatedAt    string   `json:"CreatedAt"`
	UpdatedAt    string   `json:"UpdatedAt"`
}

type CreateBook struct {
	Name       string   `json:"Name"`
	AuthorId   string   `json:"AuthorId"`
	Price      float64  `json:"Price"`
	CategoryId []string `json:"CategoryId"`
}

type UpdateBook struct {
	Name       string  `json:"Name"`
	AuthorId   string  `json:"AuthorId"`
	Price      float64 `json:"Price"`
	CategoryId string  `json:"CategoryId"`
}

type ListBooks struct {
	Books []Book `json:"Books"`
}
