package model

type ImageModel struct {
	IdImage   int    `gorm:"primary_key" json:"id_image"`
	Url       string `json:"url"`
	IdExhibit int    `json:"id_exhibit"`
}
