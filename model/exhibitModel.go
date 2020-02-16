package model

type ExhibitModel struct {
	IdExhibit   int          `gorm:"primary_key" json:"id_exhibit"`
	MuseumId    int          `json:"museum_id"`
	Rating      float32      `json:"rating"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Year        string       `json:"year"`
	Author      string       `json:"author"`
	Audio       string       `json:"audio"`
	IsPopular   bool         `json:"is_popular"`
	Images      []ImageModel `gorm:"ForeignKey:IdExhibit" json:"images"`
}

func (e *ExhibitModel) RemoveFromPopular() {
	e.IsPopular = false
}

func (e *ExhibitModel) AddToPopular() {
	e.IsPopular = true
}
