package v2

import (
	"gin-orm/pkg/e"
	"gin-orm/pkg/setting"
	"gin-orm/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-orm/app/models"
	"gin-orm/app/viewModels"
	"gin-orm/app/viewModels/emun"
)

//获取多个文章
func GetArticles(c *gin.Context) {
	maps := make(map[string]interface{})
	code := e.SUCCESS
	var viewArticles []viewModels.Article
	var viewArticle viewModels.Article
	articles, _ := models.GetArticles(util.GetPage(c), setting.PageSize, maps)
	for _, articles := range articles {
		viewArticle.Id = articles.ID
		viewArticle.Author = articles.CreatedBy
		viewArticle.DisplayTime = articles.ModifiedOn.String()
		viewArticle.Pageviews = 3474
		viewArticle.Status = emun.GetArticleStatus(articles.State)
		viewArticle.Title = articles.Title
		viewArticles = append(viewArticles, viewArticle)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": viewArticles,
	})
}
