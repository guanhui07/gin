package v1

import (
	"gin-orm/pkg/util"
	"github.com/astaxie/beego/validation"
	"net/http"

	"gin-orm/app/models"
	"gin-orm/app/viewModels"
	"gin-orm/pkg/e"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetUserInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userName := claims["userName"].(string)
	avatar := models.GetUserID(userName)

	code := e.SUCCESS
	userRoles := models.GetRoles(userName)
	data := viewModels.User{Roles: userRoles, Introduction: "", Avatar: avatar, Name: userName}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func Logout(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "success",
	})
}

func Login(c *gin.Context) {
	valid := validation.Validation{}
	var username string = ""
	if arg := c.Query("username"); arg != "" {
		username = com.ToStr(arg)
	}
	code := e.INVALID_PARAMS
	var Password string = ""
	if arg := c.Query("Password"); arg != "" {
		Password = com.ToStr(arg)
	}

	//todo:从service判断账号密码查出uid
	uid := 8
	valid.Min(uid, 1, "id").Message("ID必须大于0")

	token, err := util.GenerateToken(uid, username, Password)
	if err != nil {
		code = e.ERROR_NOT_EXIST_ARTICLE
	} else {
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": token,
	})
}
