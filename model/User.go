package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"user_name"`
	Password  string    `gorm:"size:255" json:"user_password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Vendors   []Vendor  `json:"vendors"`
	Roles     string    `json:"roles"`
}

func (self *User) Insert(db *gorm.DB) (*User, error) {
	if e := db.Create(self).Error; e != nil {
		return self, e
	}
	return self, nil
}

func (self *User) Delete(db *gorm.DB) (*User, error) {
	if e := db.Delete(self).Error; e != nil {
		return self, e
	}
	return self, nil
}

func (self *User) FindUsers(db *gorm.DB) ([]*User, error) {
	var userList []*User
	if e := db.Debug().Model(self).Find(&userList).Error; e != nil {
		return userList, e
	}
	return userList, nil
}
func (self *User) FindUserByCredential(db *gorm.DB) ([]*User, error) {
	var userList []*User
	if e := db.Debug().Where(&User{Name: self.Name, Password: self.Password}).First(&userList).Error; e != nil {
		return userList, e
	}
	return userList, nil
}
func (self *User) ListUsers(db *gorm.DB, page int, pageSize int) ([]*User, int, error) {
	users := make([]*User, 0)
	var total int = 0
	db.Model(&User{}).Count(&total)
	if e := db.Debug().Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&users).Error; e != nil {
		return users, total, e
	}
	return users, total, nil
}
