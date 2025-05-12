package entity

type Group struct {
	ID        int    `json:"id"`
	ChapterID int    `json:"chapter_id"`
	CodeStart string `json:"code_start"`
	CodeEnd   string `json:"code_end"`
	Name      string `json:"name"`
}
