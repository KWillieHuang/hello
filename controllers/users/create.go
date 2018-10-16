package users

import (
	"hello/controllers"
	"hello/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CreateController struct {
	controllers.BaseController
}

func (this *CreateController) Get() {
	this.BaseController.Verify(2)
	this.TplName = "information.tpl"
}

func (this *CreateController) Post() {
	this.BaseController.Verify(2)
	beego.Info("I get the date")
	beego.Info("create")
	e := this.GetString("Email")
	if this.BaseController.IsEmail(e) {
		beego.Info("a")
		var user models.User
		user.Username = this.GetString("Username")
		user.Password = controllers.Encode(this.GetString("Password"))
		user.Email = this.GetString("Email")
		user.Role, _ = this.GetInt("Role")
		_, err := orm.NewOrm().Insert(&user)
		beego.Info("b")
		if err == nil {
			beego.Info("c")
			this.Redirect("/admin", 302)
			return
		} else {
			beego.Info("d")
			this.Redirect("/admin", 205)
			return
		}
	} else {
		beego.Info("e")
		this.Redirect("/admin", 205)
		return
	}
}
