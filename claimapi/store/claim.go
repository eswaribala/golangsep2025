package store



type Claim struct {
	ClaimID     uint   `json:"id" gorm:"primaryKey"`
	ClaimAmount int    `json:"amount"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}
