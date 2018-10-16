package users

import (
	"hello/controllers"
	"hello/models"
	"time"

	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	this.TplName = "index.tpl"
}

func (this *LoginController) Post() {
	uname := this.GetString("username")
	pwd := controllers.Encode(this.GetString("password"))
	var users models.User
	qs := orm.NewOrm().QueryTable("user")
	err := qs.Filter("username", uname).Filter("password", pwd).One(&users)
	if err == nil {
		this.SetSession("Userid", users.Userid)
		this.SetSession("Username", users.Username)
		this.SetSession("Password", users.Password)
		this.SetSession("Role", users.Role)
		this.SetSession("Email", users.Email)
		this.SetSession("Time", time.Now().Format("15:04:05"))
		if this.GetSession("Role") == 2 {
			this.Redirect("/admin", 302)
			return
		} else {
			this.Redirect("/user", 302)
			return
		}
	} else {
		this.TplName = "index.tpl"
	}
}
