package models

type Author struct {
	AuthorId  string `json:"AuthorId"`
	Name      string `json:"Name"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

type CUAuthor struct {
	Name string `json:"Name"`
}

type ListAuthors struct {
	Authors []Author `json:"Authors"`
}
