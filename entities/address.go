package entities

import "AddressService/external"

type Address struct {
	external.SQLModel

	UserID     uint   `json:"user_id" gorm:"column:user_id"`
	Name       string `json:"name" gorm:"column:name"`
	Gender     bool   `json:"gender" gorm:"column:gender"`
	Phone      string `json:"phone" gorm:"column:phone"`
	ProvinceID uint   `json:"province_id" gorm:"column:province_id"`
	DistrictID uint   `json:"district_id" gorm:"column:district_id"`
	WardID     uint   `json:"ward_id" gorm:"column:ward_id"`
	Street     string `json:"street" gorm:"column:street"`
}
