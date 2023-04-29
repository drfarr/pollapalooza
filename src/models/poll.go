package models

type Poll struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        string `json:"belongsTo"`
	CreatedAt   string `json:"createdAt"`
	UpdateAt    string `json:"updatedAt"`
	Published   bool   `json:"published"`
}
