// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "t_user"

// User mapped from table <t_user>
type User struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	Name       *string        `gorm:"column:name" json:"name"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time" json:"delete_time"`
	CreateAt   *time.Time     `gorm:"column:create_at" json:"create_at"`
	UpdateAt   *time.Time     `gorm:"column:update_at" json:"update_at"`
	ClassID    string         `gorm:"column:class_id;not null" json:"class_id"`
	Class      Class          `gorm:"foreignKey:class_id" json:"classInfo"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
