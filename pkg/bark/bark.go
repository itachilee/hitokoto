package bark

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gookit/goutil/dump"
	"github.com/itachilee/furion/pkg/cache"
)

var ctx = context.Background()

type BarkMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// push msg to bark server
func PushToBark(msg *BarkMessage) {
	b, _ := json.Marshal(msg)
	resp, err := http.Post(BarkPushUrl, "application/json", bytes.NewBuffer(b))
	if err != nil {
		dump.P(err)
	}
	//
	defer resp.Body.Close()
	//io.Reader

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func CanPushToBark() bool {
	_, err := cache.Rdb.Get(ctx, RedisBarkExPrefix).Result()
	if err == redis.Nil {
		return true
	} else if err != nil {
		dump.V(err)
		return false
	} else {
		return false
	}
}
