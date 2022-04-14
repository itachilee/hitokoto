package gushici_service

import "github.com/itachilee/furion/models"

type Gushici struct {
	Content  string
	Origin   string
	Author   string
	Category string

	PageNum  int
	PageSize int
}

func (h *Gushici) GetAll() ([]*models.Gushici, error) {
	var (
		hitokotos []*models.Gushici
	)
	hitokotos, err := models.GetGushicis(h.PageNum, h.PageSize, h.getMaps())
	if err != nil {
		return nil, err
	}

	return hitokotos, nil
}
func (a *Gushici) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	return maps
}
