package hitokoto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gookit/goutil/dump"
	"github.com/itachilee/furion/pkg/bark"
	"github.com/robfig/cron/v3"
)

func CronRun() {
	cron2 := newWithSeconds() //创建一个cron实例

	//执行定时任务（每5秒执行一次）
	_, err := cron2.AddFunc("*/10 * * * * *", GetHitokoto)
	if err != nil {
		fmt.Println(err)
	}

	//启动/关闭
	cron2.Start()
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}

}

func GetHitokoto() {
	GetByApi()
}

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())

}

func PushToBark() {
	resp, err := http.Get(GlobalUrl)
	if err != nil {
		dump.P("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	hitokoto := &Hitokoto{}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, hitokoto)
	if err != nil {
		panic(err)
	}
	dump.P(hitokoto.Hitokoto)
	msg := &bark.BarkMessage{
		Title: hitokoto.From,
		Body:  hitokoto.Hitokoto,
	}
	bark.PushToBark(msg)
}

// 定时推送
func CronRunBark() {

	cron2 := newWithSeconds() //创建一个cron实例

	//执行定时任务（每小时推送一次）
	_, err := cron2.AddFunc("00 00 */1 * * *", PushToBark)
	if err != nil {
		fmt.Println(err)
	}

	//启动/关闭
	cron2.Start()
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}

}
