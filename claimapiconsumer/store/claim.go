package store

type Claim struct {
	ClaimID     uint   `json:"id" gorm:"primaryKey"`
	ClaimAmount int    `json:"amount" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Status      string `json:"status" gorm:"not null"`
	CreatedAt   string `json:"created_at" gorm:"not null"`
}
