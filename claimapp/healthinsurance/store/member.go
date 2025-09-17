package store

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type Member struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" gorm:"size:100;not null"`
	LastName  string `json:"last_name" gorm:"size:100;not null"`
	Email     string `json:"email" gorm:"size:100;unique;not null"`
	Phone     string `json:"phone" gorm:"size:15;not null"`
	Address   string `json:"address" gorm:"size:255;not null"`
	StartDate string `json:"start_date" gorm:"not null"`
	EndDate   string `json:"end_date" gorm:"not null"`
}

// Add other fields as necessary

func (m *Member) SaveMember() (bool, error) {
	db := ConnectionHelper()
	//defer db.Close()

	if err := db.Create(m).Error; err != nil {
		return false, err
	}
	return true, nil
}
