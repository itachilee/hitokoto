package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/itachilee/furion/pkg/app"
	"github.com/itachilee/furion/pkg/e"
	"github.com/itachilee/furion/pkg/setting"
	"github.com/itachilee/furion/pkg/util"
	"github.com/itachilee/furion/service/hitokoto_service"
	"github.com/unknwon/com"
)

func GetHitokotos(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}

	tagId := -1
	if arg := c.PostForm("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	hitokotoService := hitokoto_service.Hitokoto{
		// TagID:    tagId,
		// State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppCinfig.PageSize,
	}

	total, err := hitokotoService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	articles, err := hitokotoService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
