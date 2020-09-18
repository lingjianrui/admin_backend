package api

import (
	"backend/middleware"
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Data struct {
	Total int             `json:"total"`
	Items []*model.Vendor `json:"items"`
}

//获取商户列表
func (self *Server) GetVendorList(c *gin.Context) {
	res := make(map[string]interface{})
	res["code"] = 20000
	page := c.Query("page")
	limit := c.Query("limit")
	page_n, _ := strconv.Atoi(page)
	limit_n, _ := strconv.Atoi(limit)
	vendor := &model.Vendor{}
	vlist, total, e := vendor.ListVendors(self.DB, page_n, limit_n)
	if e != nil {
		middleware.Logger().Error(e.Error(), "get vendor list error")
		res["code"] = 1
		res["error"] = "db error"
		res["message"] = e.Error()
	} else {
		res["data"] = &Data{Total: total, Items: vlist}
	}
	c.JSON(http.StatusOK, res)
}

//增加商户
func (self *Server) CreateVendor(c *gin.Context) {
	res := make(map[string]interface{})
	res["code"] = 20000
	n := c.PostForm("vendor_name")
	phone := c.PostForm("phone")
	banner := c.PostForm("banner")
	image720 := c.PostForm("image720")
	desc := c.PostForm("description")
	vendor := &model.Vendor{Name: n, Phone: phone, Banner: banner, Image720: image720, Description: desc}
	v, e := vendor.Insert(self.DB)
	if e != nil {
		middleware.Logger().Error(e.Error(), "create vendor error")
		res["code"] = 1
		res["error"] = e
		res["message"] = e.Error()
	} else {
		res["data"] = v
	}
	c.JSON(http.StatusOK, res)
}

//删除商户
func (self *Server) DeleteVendor(c *gin.Context) {
	res := make(map[string]interface{})
	res["code"] = 20000
	id := c.Query("id")
	id_n, _ := strconv.Atoi(id)
	vendor := &model.Vendor{ID: uint(id_n)}
	v, e := vendor.Delete(self.DB)
	if e != nil {
		middleware.Logger().Error(e.Error(), "DeleteVendor")
		res["code"] = 1
		res["error"] = "delete vendor error"
		res["message"] = e.Error()
	} else {
		res["data"] = v
	}
	c.JSON(http.StatusOK, res)
}
