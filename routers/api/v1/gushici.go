package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itachilee/furion/pkg/app"
	"github.com/itachilee/furion/pkg/e"
	"github.com/itachilee/furion/pkg/setting"
	"github.com/itachilee/furion/pkg/util"
	"github.com/itachilee/furion/service/gushici_service"
)

func GetGushiciAll(c *gin.Context) {
	appG := app.Gin{C: c}
	gushiciService := gushici_service.Gushici{
		// TagID:    tagId,
		// State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppCinfig.PageSize,
	}
	articles, err := gushiciService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = 10

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
