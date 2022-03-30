package hitokoto

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gookit/goutil/dump"
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

func HttpGet(url string) (err error) {
	resp, err := http.Get(url)
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
	return
}

func (h *Hitokoto) saveToRedis() {
	rdb, err := cache.NewRedis(&redis.Options{
		Addr:     "localhost:6380",
		DB:       0,
		Password: "123456",
	})
	if err != nil {
		panic(err)
	}
	key := fmt.Sprintf("%s:%s:%d", redisPrefix, h.Type, h.ID)
	str, _ := json.Marshal(h)
	rdb.Set(context.Background(), key, string(str), 0)

}
