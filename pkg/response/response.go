package response

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}


/*
func EditArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
	return


 */