package models

//self refrential structure

type Claim struct {
	ID     uint     `json:"id" gorm:"primaryKey"`
	Amount int      `json:"amount"`
	Claims []*Claim `json:"claims"`
}

func (c *Claim) TotalClaimsAmount() int64 {
	var total int64
	total += int64(c.Amount)
	for _, claim := range c.Claims {
		total += int64(claim.TotalClaimsAmount())
	}
	return total
}
