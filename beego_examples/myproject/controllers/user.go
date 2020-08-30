package controllers

import (
	"fmt"
	"myproject/models"
	"myproject/utils"
	"time"

	"github.com/astaxie/beego"
)

// RegisterController =>
type RegisterController struct {
	beego.Controller
}

// Get => RegisterController
func (c *RegisterController) Get() {
	fmt.Println("RegisterController Get!!!")
	c.TplName = "register.html"
}

// Post => RegisterController
func (c *RegisterController) Post() {
	fmt.Println("RegisterController Get!!!")
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	fmt.Println(username, password, repassword)

	id := models.QueryUserWithUsername(username)
	fmt.Println("id", id)
	if id > 0 {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "username exist!"}
		c.ServeJSON()
		return
	}

	password = utils.MD5(password)
	fmt.Println("After md5: ", password)

	user := models.User{
		ID:         0,
		Username:   username,
		Password:   password,
		Status:     0,
		Createtime: time.Now().Unix(),
	}
	_, err := models.InsertUser(user)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "register failed!"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "register successed!"}
	}
	c.ServeJSON()
}
