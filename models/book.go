package models

type Book struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MediumPrice float32 `json:"medium_price"`
	Author      string  `json:"author"`
	ImageURL    string  `json:"img_url"`
}
