package model

import (
	"gorm.io/gorm"
	"zg5/z311/framework/mysql"
)

type Shop struct {
	gorm.Model
	Name  string
	Price float64
	Num   int64
}

func NewShop() *Shop {
	return new(Shop)
}

func (s *Shop) Create(info *Shop) (*Shop, error) {
	err := mysql.DB.Create(info).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (s *Shop) Select() ([]*Shop, error) {
	var ss []*Shop
	err := mysql.DB.Find(&ss).Error
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (s *Shop) Update(id, num int64) error {
	err := mysql.DB.Model(s).Where("id = ?", id).Update("num", gorm.Expr("num + ?", num)).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Shop) Delete(id int64) error {
	err := mysql.DB.Where("id = ?", id).Delete(s).Error
	if err != nil {
		return err
	}
	return nil
}
