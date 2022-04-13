package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itachilee/furion/models"
	"github.com/itachilee/furion/pkg/gushici"
	"github.com/itachilee/furion/pkg/hitokoto"
	"github.com/itachilee/furion/pkg/setting"
	v1 "github.com/itachilee/furion/routers/api/v1"
)

const (
	KEY Request = "trace_id"
)

type Request string

func NewRequestID() Request {
	requestId := strings.Replace(uuid.New().String(), "-", "", -1)
	return Request(requestId)
}

func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
	return ctx
}

func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
}

func GetContextValue(ctx context.Context, k Request) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "dream worker")
}
func main() {

	// ProcessEnter(NewContextWithTraceID())

	setting.Setup("./conf")
	models.Setup()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hitokotos", v1.GetHitokotos)
	go hitokoto.CronRun()

	r.Run(":8084")

	gushici.GetGushici()
}
