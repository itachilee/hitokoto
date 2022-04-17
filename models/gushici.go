package models

import "gorm.io/gorm"

type Gushici struct {
	gorm.Model
	Content  string `gorm:"comment:'正文'"`
	Origin   string `gorm:"comment:'起源'"`
	Author   string `gorm:"comment:'作者'"`
	Category string `gorm:"comment:'分类'"`
}

func GetGushicis(pageNum int, pageSize int, maps interface{}) ([]*Gushici, error) {
	var gushicis []*Gushici
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&gushicis).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return gushicis, nil
}

func GetGushiciTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Gushici{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
