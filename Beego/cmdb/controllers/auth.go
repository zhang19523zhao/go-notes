package controllers

import (
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

// AuthController 认证控制器
type AuthController struct {
	beego.Controller
}

// Login 认证登录
func (c *AuthController) Login() {
	user := c.GetSession("user")
	if user != nil {
		c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
	}

	form := &forms.LoginForm{}

	errs := errors.New()
	// Get请求直接加载页面
	// Post请求进行数据验证
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			user := models.GetUserByName(form.Name)
			if user == nil {
				//用户不存在
				fmt.Println(form)
				errs.Add("default", "用户或密码错误")
				fmt.Println("用户不存在")
			} else if user.ValidPassword(form.Password) {
				//用户密码正确
				//fmt.Println(form)
				c.SetSession("user", user.ID)
				c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
			} else {
				//用户密码不正确
				errs.Add("default", "用户或密码错误")
			}
		} else {
			fmt.Println(form)
		}

		//验证成功
		//验证失败
	}
	c.Data["form"] = form
	c.Data["errors"] = errs
	errs.Errors()
	c.TplName = "auth/login.html"
}

// Logout用户退出登录
func (c *AuthController) Logout() {
	c.DestroySession()

	fmt.Println("des session")
	c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$")
}
