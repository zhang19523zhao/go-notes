package controllers

import (
	"cmdb/models"
	"github.com/astaxie/beego"
	"net/http"
)

// UserController 用户管理控制器
type UserController struct {
	beego.Controller
}

// 查询用户
func (c *UserController) Query() {
	sessionUser := c.GetSession("user")
	if sessionUser == nil {
		// 无Session信息(未登录)
		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
		return
	}

	name := c.GetString("name")
	users := []*models.User{}
	if c.Ctx.Input.IsPost() {
		users = models.QueryUser(name)
	}
	c.Data["users"] = users
	c.Data["q"] = name
	c.TplName = "user/query.html"
}
