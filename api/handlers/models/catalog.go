package models

type Category struct {
	CategoryId     string `json:"categoryId"`
	Name           string `json:"name"`
	ParentUuid     string `json:"parentUuid"`
	ParentCategory string `json:"parentCategory"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

type CreateCategory struct {
	Name       string `json:"name"`
	ParentUuid string `json:"parentUuid"`
}

type UpdateCategory struct {
	Name       string `json:"name"`
	ParentUuid string `json:"parentUuid"`
}

type ListCategories struct {
	Categories []Category `json:"categories"`
}

type Book struct {
	BookId       string  `json:"BookId"`
	Name         string  `json:"Name"`
	AuthorId     string  `json:"AuthorId"`
	Price        float64 `json:"Price"`
	CategoryId   string  `json:"CategoryId"`
	CategoryName string  `json:"CategoryName"`
	CreatedAt    string  `json:"CreatedAt"`
	UpdatedAt    string  `json:"UpdatedAt"`
}

type CreateBook struct {
	Name       string  `json:"Name"`
	AuthorId   string  `json:"AuthorId"`
	Price      float64 `json:"Price"`
	CategoryId string  `json:"CategoryId"`
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

type Author struct {
	AuthorId  string `json:"AuthorId"`
	Name      string `json:"Name"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

type List struct {
	Author     Author     `json:"Author"`
	Book       Book       `json:"Book"`
	Categories []Category `json:"Categories"`
}
