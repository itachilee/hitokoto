package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/itachilee/furion/models"
	"github.com/itachilee/furion/pkg/hitokoto"
	"github.com/itachilee/furion/pkg/setting"
	v1 "github.com/itachilee/furion/routers/api/v1"
)

var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var seed = rand.NewSource(time.Now().UnixNano())

func handler(c *gin.Context) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(c.Writer, f) // Render 可接收多个 io.Writer 接口
}

func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, int(seed.Int63())%50)
	}
	return r
}

func main() {

	setting.Setup("./conf")
	models.Setup()
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8081", nil)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", handler)
	r.GET("/hitokotos", v1.GetHitokotos)
	go hitokoto.CronRun()

	r.Run(":8081")
}
