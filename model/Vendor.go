package model

import (
	"github.com/jinzhu/gorm"
)

type Vendor struct {
	Base
	Name        string `gorm:"size:255;not null" json:"vendor_name"`
	Banner      string `json:"banner"`
	Phone       string `json:"phone"`
	Image720    string `json:"image720"`
	Description string `gorm:"size:255" json:"description"`
	UserId      uint   `json:"user_id"`
}

func (self *Vendor) Insert(db *gorm.DB) (*Vendor, error) {
	if e := db.Create(self).Error; e != nil {
		return self, e
	}
	return self, nil
}

func (self *Vendor) Delete(db *gorm.DB) (*Vendor, error) {
	if e := db.Delete(self).Error; e != nil {
		return self, e
	}
	return self, nil
}

func (self *Vendor) FindVendors(db *gorm.DB) ([]*Vendor, error) {
	var vendorList []*Vendor
	if e := db.Debug().Model(self).Find(&vendorList).Error; e != nil {
		return vendorList, e
	}
	return vendorList, nil
}

func (self *Vendor) ListVendors(db *gorm.DB, page int, pageSize int) ([]*Vendor, int, error) {
	vendors := make([]*Vendor, 0)
	var total int = 0
	db.Model(&Vendor{}).Count(&total)
	if e := db.Debug().Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&vendors).Error; e != nil {
		return vendors, total, e
	}
	return vendors, total, nil
}
