package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdateAt  time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	Mobel    string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	Nickname string     `gorm:"type:varchar(20) comment '昵称'"`
	Birthday *time.Time `gorm:"type:datetime comment '生日'"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female是女性,male是男性'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1是普通用户，2是管理员'"`
}
