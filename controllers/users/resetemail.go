package users

import (
	"hello/controllers"
	"hello/models"

	"github.com/astaxie/beego/orm"
)

type ResetemailController struct {
	controllers.BaseController
}

func (this *ResetemailController) Get() {
	this.BaseController.Verify(1)
	this.BaseController.SendToTpl()
	this.TplName = "resetemail.tpl"
}

func (this *ResetemailController) Post() {
	this.BaseController.Verify(1)
	e := this.GetString("Email")
	if this.BaseController.IsEmail(e) {
		id := this.GetSession("Userid").(int)
		user := models.User{Userid: id}
		user.Email = this.GetString("Email")
		orm.NewOrm().Update(&user, "Email")
		this.Redirect("/resetemail", 302)
		return
	} else {
		this.Redirect("/resetemail", 205)
		return
	}

}
