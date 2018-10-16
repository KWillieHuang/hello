package users

import (
	"hello/controllers"
	"hello/models"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ResetpwdController struct {
	controllers.BaseController
}

func (this *ResetpwdController) Get() {
	this.BaseController.Verify(1)
	this.BaseController.SendToTpl()
	this.TplName = "resetpwd.tpl"
}

func (this *ResetpwdController) Post() {
	this.BaseController.Verify(1)
	id := this.GetSession("Userid").(int)
	surepwd := this.GetString("SurePassword")
	newpwd := this.GetString("NewPassword")
	oldpwd := this.GetString("OldPassword")
	pwd := controllers.Encode(oldpwd)
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	result := qs.Filter("Userid", id).Filter("password", pwd)
	if strings.EqualFold(surepwd, newpwd) && result.Exist() {
		user := models.User{Userid: id}
		user.Password = controllers.Encode(newpwd)
		o.Update(&user, "Password")
		this.Redirect("/resetpwd", 302)
		return
	} else {
		this.Redirect("/resetpwd", 205)
		return
	}
}
