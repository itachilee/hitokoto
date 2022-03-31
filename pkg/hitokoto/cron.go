package hitokoto

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func CronRun() {
	// config := &logredis.HookConfig{
	// 	Addr:     "localhost:6380",
	// 	Password: "123456",
	// 	DB:       0,
	// }
	// h, err := logredis.NewHook(*config)
	// if err != nil {
	// 	panic(err)
	// }

	// logrus.AddHook(h)

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
