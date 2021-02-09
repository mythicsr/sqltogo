package sql_to_go_test

import (
	"time"
)

type USER struct {
	Id           int64     `json:"id" gorm:"column:id"`                       //primary key
	IpAddress    int64     `json:"ip_address" gorm:"column:ip_address"`       //ip_address
	Nickname     string    `json:"nickname" gorm:"column:nickname"`           //user note
	Description  string    `json:"description" gorm:"column:description"`     //user description
	CreatorEmail string    `json:"creator_email" gorm:"column:creator_email"` //creator email
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`       //create time
	DeletedAt    time.Time `json:"deleted_at" gorm:"column:deleted_at"`       //delete time
}
