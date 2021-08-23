package entity

import "time"

type Account struct {
	Id              string     `gorm:"primary_key; size:32; not null"`
	Password        string     `gorm:"size:60; not null"`
	Name            string     `gorm:"size:15; not null"`
	Email           string     `gorm:"unique; size:50; not null"`
	PhoneNumber     string     `gorm:"unique; size:13; not null"`
	DepartmentCode  int        `gorm:"default: null"`
	PositionCode    int        `gorm:"default: null"`
	AuthorityCode   int        `gorm:"default: null"`
	FirstPaymentId  string     `gorm:"size:32"`
	FinalPaymentId  string     `gorm:"size:32"`
	WorkCode        int        `gorm:"not null"`
	TotalAnnual     float32    `gorm:"default: null"`
	UseAnnual       float32    `gorm:"default: null"`
	RemainAnnual    float32    `gorm:"default: null"`
	HireDate        string     `gorm:"not null"`
	UpdateDate      time.Time  `gorm:"default:current_timestamp; not null"`
	ResignationDate *time.Time `gorm:"default: null"`
}

func (Account) TableName() string {
	return "account"
}
