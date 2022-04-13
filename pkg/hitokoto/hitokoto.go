package hitokoto

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gookit/goutil/dump"
	"github.com/itachilee/furion/models"
	"github.com/itachilee/furion/pkg/bark"
	"github.com/itachilee/furion/pkg/cache"
)

const (
	redisPrefix = "hitokoto"
)

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

func GetByApi() (err error) {

	resp, err := http.Get(GlobalUrl)
	if err != nil {
		dump.P("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	hitokoto := &Hitokoto{}

	err = json.Unmarshal(body, hitokoto)
	if err != nil {
		panic(err)
	}
	dump.P(hitokoto)
	hitokoto.saveToRedis()
	hitokoto.saveToMysql()
	hitokoto.PushToBark()
	return
}

func (h *Hitokoto) PushToBark() {

	msg := &bark.BarkMessage{
		Title: h.From,
		Body:  h.Hitokoto,
	}
	bark.PushToBark(msg)
}

func (h *Hitokoto) saveToRedis() {

	key := fmt.Sprintf("%s:%s:%d", redisPrefix, h.Type, h.ID)
	str, _ := json.Marshal(h)
	cache.Rdb.Set(context.Background(), key, string(str), 0)

}

func (h *Hitokoto) saveToMysql() {
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
	models.AddHitokoto(hitokoto)

}
