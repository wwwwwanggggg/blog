package controller

import (
	"blog/common"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var adminName, adminPassword string = "savanna", convert("123456")

// 加密密码的函数
func convert(input string) (output string) {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

type LoginInfo struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Admin struct {
}

// 登录函数
func (a *Admin) Login(c *gin.Context) {
	fmt.Println(c.FullPath())
	var info LoginInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	if !(info.Name == adminName && convert(info.Password) == adminPassword) {
		c.Error(common.ErrNew(errors.New("用户名或密码输入错误"), common.UserErr))
		return
	}
	SessionSet(c, "admin", UserSession{
		ID:       1,
		Username: info.Name,
		Level:    2,
	})
	resp := struct {
		Name string `json:"name"`
	}{
		Name: info.Name,
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp, "成功登录"))
}

// 登出函数
func (a *Admin) Logout(c *gin.Context) {
	SessionClear(c)
	c.JSON(http.StatusOK, ResponseNew(c, nil, "成功退出登录"))
}

// 测试函数
func (a *Admin) Test(c *gin.Context) {
	fmt.Println(c.FullPath())
	var info struct {
		TestMsg string `json:"test_msg"`
	}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, map[string]string{
		"key": "value",
	})
	// c.JSON(http.StatusOK, ResponseNew(c, map[string]string{
	// 	"msg": info.TestMsg,
	// }, "操作成功"))
	fmt.Println("Hello World")
}
