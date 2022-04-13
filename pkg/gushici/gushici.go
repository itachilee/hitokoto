package gushici

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gookit/goutil/dump"
	"github.com/itachilee/furion/pkg/cache"
)

type Gushici struct {
	Content  string `json:"content"`
	Origin   string `json:"origin"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

var ctx = context.Background()

func GetGushici() {
	resp, err := http.Get(GushiciUrl)
	if err != nil {
		dump.P("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dump.P(err)
	}
	gushici := &Gushici{}

	err = json.Unmarshal(body, gushici)
	if err != nil {
		panic(err)
	}
	dump.P(gushici)
	gushici.saveToRedis()
}

func (h *Gushici) saveToRedis() {

	key := fmt.Sprintf("%s:%s", RedisPrefix, GetRedisId())
	str, _ := json.Marshal(h)
	cache.Rdb.Set(ctx, key, string(str), 0)

}

func GetRedisId() string {

	cache.Rdb.IncrBy(ctx, RedisExKey, 1)
	v, err := cache.Rdb.Get(ctx, RedisExKey).Result()
	if err != nil {
		dump.V(err)
	}
	return v
}
