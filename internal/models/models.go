package models

import "time"

type Service struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
    TitleEn       string     `json:"title_en"`
	TitleId       string     `json:"title_id"`
	DescriptionEn string     `json:"description_en"`
	DescriptionId string     `json:"description_id"`
	IconURL       string     `json:"icon_url"`
	CreatedAt     time.Time  `json:"created_at"`
}

type Team struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	RoleEn    string     `json:"role_en"`
	RoleId    string     `json:"role_id"`
	BioEn     string     `json:"bio_en"`
	BioId     string     `json:"bio_id"`
	ImageURL  string     `json:"image_url"`
	CreatedAt time.Time  `json:"created_at"`
}

type Article struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	TitleEn   string     `json:"title_en"`
	TitleId   string     `json:"title_id"`
	ContentEn string     `json:"content_en"`
	ContentId string     `json:"content_id"`
	ImageURL  string     `json:"image_url"`
	CreatedAt time.Time  `json:"created_at"`
}