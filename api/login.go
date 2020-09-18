package api

import (
	"backend/model"
	"backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Login 登录请求
func (server *Server) Login(c *gin.Context) {
	res := make(map[string]interface{})
	res["code"] = 20000
	username := c.PostForm("username")
	password := c.PostForm("password")
	u := &model.User{Name: username, Password: password}
	ul, e := u.FindUserByCredential(server.DB)
	if e != nil {
		fmt.Println(e.Error())
	}
	if len(ul) == 1 {
		signedToken, err := util.GenerateToken(username, password)
		if err != nil {
			res["code"] = 1
			res["error"] = "generate token error"
			res["message"] = err.Error()
		} else {
			token := struct {
				Token string `json:"token"`
			}{
				Token: signedToken,
			}
			res["data"] = token
		}
	} else {
		res["code"] = 1
		res["error"] = "login error"
		res["message"] = "用户名或密码不正确"
	}
	c.JSON(http.StatusOK, res)
}
