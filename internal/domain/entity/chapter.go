package entity

type Chapter struct {
	ID        int    `json:"id"`
	Roman     string `json:"roman"`
	CodeStart string `json:"code_start"`
	CodeEnd   string `json:"code_end"`
	Name      string `json:"name"`
}
