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
