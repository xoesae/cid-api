package entity

type Category struct {
	ID      int    `json:"id"`
	GroupID int    `json:"group_id" db:"group_id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
}
