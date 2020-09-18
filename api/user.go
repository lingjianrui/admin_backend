package api

import (
	"backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserData struct {
	Total int           `json:"total"`
	Items []*model.User `json:"items"`
}

//获取用户详情
func (self *Server) GetUserInfo(c *gin.Context) {
	token := c.Query("token")
	fmt.Printf(token)
	res := make(map[string]interface{})
	res["code"] = 20000
	d := struct {
		Roles         []string `json:"roles"`
		Instroduction string   `json:"introduction"`
		Avatar        string   `json:"avatar"`
		Name          string   `json:"name"`
	}{
		Roles:         []string{"admin"},
		Instroduction: "test",
		Avatar:        "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:          "xiaohei",
	}
	res["data"] = d
	c.JSON(http.StatusOK, res)
}

//获取用户列表
func (self *Server) GetUserList(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	page_n, _ := strconv.Atoi(page)
	limit_n, _ := strconv.Atoi(limit)
	res := make(map[string]interface{})
	user := &model.User{}
	vlist, total, e := user.ListUsers(self.DB, page_n, limit_n)
	res["code"] = 20000
	if e != nil {
		res["code"] = 1
		res["error"] = "Get User List Error"
		res["message"] = e.Error()
	} else {
		res["data"] = &UserData{Total: total, Items: vlist}
	}
	c.JSON(http.StatusOK, res)
}

//增加用户
func (self *Server) CreateUser(c *gin.Context) {
	username := c.PostForm("name")
	pwd := c.PostForm("password")
	roles := c.PostForm("roles")
	//vlist := []model.Vendor{}
	//vlist = append(vlist, model.Vendor{Name: "aaaa"})
	user := &model.User{Name: username, Password: pwd, Roles: roles}
	res := make(map[string]interface{})
	u, e := user.Insert(self.DB)
	res["code"] = 20000
	if e != nil {
		res["code"] = 1
		res["error"] = "create user error"
		res["message"] = e.Error()
	} else {
		res["data"] = u
	}
	c.JSON(http.StatusOK, res)
}

//删除用户
func (self *Server) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	id_n, _ := strconv.Atoi(id)
	user := &model.User{ID: uint(id_n)}
	v, e := user.Delete(self.DB)
	res := make(map[string]interface{})
	res["code"] = 20000
	if e != nil {
		res["code"] = 1
		res["error"] = "delete user error"
		res["message"] = e.Error()
	} else {
		res["data"] = v
	}
	c.JSON(http.StatusOK, res)
}
