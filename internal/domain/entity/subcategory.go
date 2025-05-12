package entity

type Subcategory struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id" db:"category_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
}
