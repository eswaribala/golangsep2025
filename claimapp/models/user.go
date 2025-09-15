package models

//basic structure

type User struct {
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

type Admin struct {
	User
	AssignRole bool `gorm:"default:false" json:"assign_role"`
}

type Claimofficer struct {
	User
	ClaimsHandled int  `gorm:"default:0" json:"claims_handled"`
	ProcessClaims bool `gorm:"default:true" json:"process_claims"`
}

type ClaimAuthorizer struct {
	User
	FundAuthorization bool `gorm:"default:false" json:"fund_authorization"`
}

type Customer struct {
	User
	FullName string `gorm:"not null" json:"full_name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Phone    string `gorm:"unique;not null" json:"phone"`
}
