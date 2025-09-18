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

func (m *Member) GetAllMembers() ([]*Member, error) {
	db := ConnectionHelper()
	var members []*Member
	if err := db.Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (m *Member) GetMemberByID(id uint) (*Member, error) {
	db := ConnectionHelper()
	var member Member
	if err := db.First(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *Member) UpdateMember(email string, contactNo string) (bool, error) {
	db := ConnectionHelper()
	if err := db.Model(m).Updates(Member{Email: email, Phone: contactNo}).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (m *Member) DeleteMember(id uint) (bool, error) {
	db := ConnectionHelper()
	if err := db.Delete(&Member{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (m *Member) GetMemberByEmail(email string) (*Member, error) {
	db := ConnectionHelper()
	var member Member
	if err := db.Where("email = ?", email).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}
func (m *Member) GetMemberByPhone(phone string) (*[]Member, error) {
	db := ConnectionHelper()
	var members []Member

	if err := db.Where("phone = ?", phone).Find(&members).Error; err != nil {
		return nil, err
	}
	return &members, nil
}
