package models

type Calculator struct {
	ID       uint
	UserID   uint   `json:"user_id"`
	ImageUrl string `json:"image_url"`
	Text     string `json:"text"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}
