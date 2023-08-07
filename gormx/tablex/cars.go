package tablex

import "github.com/google/uuid"

func (Cars) TableName() string {
	return `cars`
}

type Cars struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Car      string    `gorm:"" json:"car"`
	Name     string    `gorm:"" json:"name"`
	Year     string    `gorm:"" json:"year"`
	Distance float64   `gorm:"" json:"distance"`
	Owner    int64     `gorm:"" json:"owner"`
	Fuel     string    `gorm:"" json:"fuel"`
	Location string    `gorm:"" json:"location"`
	Drive    string    `gorm:"" json:"drive"`
	Type     string    `gorm:"" json:"type"`
	Price    float64   `gorm:"" json:"price"`
}
