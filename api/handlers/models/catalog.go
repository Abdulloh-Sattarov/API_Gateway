package models

type Category struct {
	CategoryId     string `json:"category_id"`
	Name           string `json:"name"`
	ParentUuid     string `json:"parent_uuid"`
	ParentCategory string `json:"parent_category"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type CreateCategory struct {
	CategoryId string `json:"category_id"`
	Name       string `json:"name"`
	ParentUuid string `json:"parent_uuid"`
}

type UpdateCategory struct {
	Name       string `json:"name"`
	ParentUuid string `json:"parent_uuid"`
}

type ListCategories struct {
	Categories []Category `json:"categories"`
}
