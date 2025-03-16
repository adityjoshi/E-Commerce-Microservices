package models

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type UserType string

const (
	Buyer  UserType = "Buyer"
	Seller UserType = "Seller"
	Admin  UserType = "Admin"
)

type Users struct {
	ID            uint     `gorm:"primaryKey"`
	Full_Name     string   `json:"Full_Name" gorm:"not null"`
	GenderInfo    Gender   `json:"GenderInfo"`
	ContactNumber string   `json:"ContactNumber" gorm:"not null"`
	Email         string   `json:"Email" gorm:"not null;unique"`
	Password      string   `json:"Password"`
	Region        string   `json:"region"`
	User_type     UserType `json:"User_type"`
}

type UserLogin struct {
	Email    string `json:"Email" gorm:"not null;unique"`
	Password string `json:"Password" gorm:"not null"`
}
