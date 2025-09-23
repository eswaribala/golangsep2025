package models

type Member struct {
	ID                   int    `json:"id"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	LowIncome            bool   `json:"low_income"`
	PreferredContact     string `json:"preferred_contact"`
	Consent              bool   `json:"consent"`
	EmergencyContactName string `json:"emergency_contact_name"`
}
