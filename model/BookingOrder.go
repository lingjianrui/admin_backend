package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BookingOrder struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"vendor_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Phone     string
	UserId    uint `json:"user_id"`
}

func (self *BookingOrder) Insert(db *gorm.DB) (*BookingOrder, error) {
	if e := db.Create(self).Error; e != nil {
		return self, e
	}
	return self, nil
}

func (self *BookingOrder) Delete(db *gorm.DB) (*BookingOrder, error) {
	if e := db.Delete(self).Error; e != nil {
		return self, e
	}
	return self, nil
}

func (self *BookingOrder) FindBookingOrders(db *gorm.DB) ([]*BookingOrder, error) {
	var vendorList []*BookingOrder
	if e := db.Debug().Model(self).Find(&vendorList).Error; e != nil {
		return vendorList, e
	}
	return vendorList, nil
}

func (self *BookingOrder) ListBookingOrders(db *gorm.DB, page int, pageSize int) ([]*BookingOrder, int, error) {
	orders := make([]*BookingOrder, 0)
	var total int = 0
	db.Model(&BookingOrder{}).Count(&total)
	if e := db.Debug().Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&orders).Error; e != nil {
		return orders, total, e
	}
	return orders, total, nil
}
