package models

import "gorm.io/gorm"

type Hitokoto struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

func GetHitokotos(pageNum int, pageSize int, maps interface{}) ([]*Hitokoto, error) {
	var hitokotos []*Hitokoto
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&hitokotos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return hitokotos, nil
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetHitokotoTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Hitokoto{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func AddHitokoto(data map[string]interface{}) error {
	hitokoto := Hitokoto{
		ID:         data["ID"].(int),
		UUID:       data["UUID"].(string),
		Hitokoto:   data["Hitokoto"].(string),
		Type:       data["Type"].(string),
		From:       data["From"].(string),
		FromWho:    data["FromWho"].(string),
		Creator:    data["Creator"].(string),
		CreatorUID: data["CreatorUID"].(int),
		Reviewer:   data["Reviewer"].(int),
		CommitFrom: data["CommitFrom"].(string),
		CreatedAt:  data["CreatedAt"].(string),
		Length:     data["Length"].(int),
	}
	if err := db.Create(&hitokoto).Error; err != nil {
		return err
	}

	return nil
}
