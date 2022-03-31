package hitokoto_service

import "github.com/itachilee/furion/models"

type Hitokoto struct {
	ID         int
	UUID       string
	Hitokoto   string
	Type       string
	From       string
	FromWho    string
	Creator    string
	CreatorUID int
	Reviewer   int
	CommitFrom string
	CreatedAt  string
	Length     int

	PageNum  int
	PageSize int
}

func (h *Hitokoto) GetAll() ([]*models.Hitokoto, error) {
	var (
		hitokotos []*models.Hitokoto
	)
	hitokotos, err := models.GetHitokotos(h.PageNum, h.PageSize, h.getMaps())
	if err != nil {
		return nil, err
	}

	return hitokotos, nil
}

func (h *Hitokoto) Add() error {
	hitokoto := map[string]interface{}{
		"ID":         h.ID,
		"UUID":       h.UUID,
		"Hitokoto":   h.Hitokoto,
		"Type":       h.Type,
		"From":       h.From,
		"FromWho":    h.FromWho,
		"Creator":    h.Creator,
		"CreatorUID": h.CreatorUID,
		"Reviewer":   h.Reviewer,
		"CommitFrom": h.CommitFrom,
		"CreatedAt":  h.CreatedAt,
		"Length":     h.Length,
	}
	if err := models.AddHitokoto(hitokoto); err != nil {
		return err
	}

	return nil
}
func (a *Hitokoto) Count() (int, error) {
	return models.GetHitokotoTotal(a.getMaps())
}

func (a *Hitokoto) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	return maps
}
