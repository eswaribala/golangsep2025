package models


type Policy struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Coverage struct {
		MaxAmount float64 `gorm:"not null" json:"max_amount"`
		MinAmount float64 `gorm:"not null" json:"min_amount"`
	} `gorm:"embedded;embeddedPrefix:coverage_" json:"coverage"`
}
