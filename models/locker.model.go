package models

type Locker struct {
	ID     int    `json:"id" form:"id" gorm:"primaryKey"`
	Code   string `json:"code" form:"code" gorm:"requires,not null"`
	UserId int    `json:"user_id" form:"user_id" gorm:"required,not null"`
	User   UserResponse   // Embeded
}

type LockerResponse struct {
	ID     int    `json:"id" form:"id"`
	Code   string `json:"code" form:"code"`
	UserId int    `json:"-" form:"user_id"`
}

func (LockerResponse) TableName() string {
	return "lockers"
}
