package users

import (
	"hello/controllers"
	"hello/models"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	controllers.BaseController
}

func (this *AdminController) Get() {
	beego.Info(this.GetSession("Role"))
	this.BaseController.Verify(2)
	this.BaseController.SendToTpl()
	var users []*models.User
	num, _ := orm.NewOrm().QueryTable("user").All(&users)
	this.Data["num"] = num
	this.Data["users"] = users
	this.TplName = "admin.tpl"
}
