package routers

import (
	"gin-orm/pkg/response"
	"github.com/gin-gonic/gin"

	v1 "gin-orm/app/controller/api/v1"
	v2 "gin-orm/app/controller/api/v2"
	"gin-orm/app/middleware/cors"
	"gin-orm/app/middleware/myjwt"
	//"gin-orm/pkg/e"
	"gin-orm/pkg/setting"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	//var authMiddleware = myjwt.GinJWTMiddlewareInit(&myjwt.AllUserAuthorizator{})

	//404 handler
	r.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1, "pong", nil)
	})

	//r.POST("/login", authMiddleware.LoginHandler)

	//测试登陆返回token
	r.GET("/login", v1.Login)

	apiV1 := r.Group("/api/v1")

	apiV1.Use(myjwt.JWT()) // token 验证
	apiV1.GET("/userInfo", v1.GetUserInfo)

	//gorm获取文章列表
	r.GET("/get/articles", v1.GetArticles)
	r.GET("/articles", v2.GetArticles)
	//获取指定文章
	r.GET("/article/:id", v1.GetArticle)
	//新建文章
	r.POST("/article", v1.AddArticle)
	//更新指定文章
	r.PUT("/article/:id", v1.EditArticle)
	//删除指定文章
	r.DELETE("/article/:id", v1.DeleteArticle)

	//获取标签
	r.GET("/get/tags", v1.GetTags)
	//新建标签
	r.POST("/add/tag", v1.AddTag)
	//删除指定标签
	r.DELETE("/tag/:id", v1.DeleteTag)
	//更新指定标签
	r.PUT("/tag/:id", v1.EditTag)

	r.GET("/info", v1.GetUserInfo)
	r.POST("/logout", v1.Logout)

	//var adminMiddleware = myjwt.GinJWTMiddlewareInit(&myjwt.AdminAuthorizator{})
	//apiv1 := r.Group("/api/v1")
	////使用AdminAuthorizator中间件，只有admin权限的用户才能获取到接口
	//apiv1.Use(adminMiddleware.MiddlewareFunc())
	//{
	//	//获取table信息
	//	apiv1.GET("/table/list", v2.GetArticles)
	//	//获取标签列表
	//	apiv1.GET("/tags", v1.GetTags)
	//	//新建标签
	//	apiv1.POST("/tags", v1.AddTag)
	//	//更新指定标签
	//	apiv1.PUT("/tags/:id", v1.EditTag)
	//	//删除指定标签
	//	apiv1.DELETE("/tags/:id", v1.DeleteTag)
	//
	//	//获取文章列表
	//	apiv1.GET("/articles", v1.GetArticles)
	//	//获取指定文章
	//	apiv1.GET("/articles/:id", v1.GetArticle)
	//	//新建文章
	//	apiv1.POST("/articles", v1.AddArticle)
	//	//更新指定文章
	//	apiv1.PUT("/articles/:id", v1.EditArticle)
	//	//删除指定文章
	//	apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	//}

	//var testMiddleware = myjwt.GinJWTMiddlewareInit(&myjwt.TestAuthorizator{})
	//apiv2 := r.Group("/api/v2")
	//apiv2.Use(testMiddleware.MiddlewareFunc())
	//{
	//	//获取文章列表
	//	apiv2.GET("/articles", v2.GetArticles)
	//}

	return r
}
