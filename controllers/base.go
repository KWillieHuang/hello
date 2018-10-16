package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

//parent class
type BaseController struct {
	beego.Controller
	User models.User
}

//verify if sufficient permissions
func (this *BaseController) Verify(permissions int) {
	role, err := this.GetSession("Role").(int)
	username := this.GetSession("Username")
	password := this.GetSession("Password")
	if !err || username == nil || password == nil || role < permissions {
		this.Redirect("/", 401)
		return
	}
}

//send which get from session to tpl
func (this *BaseController) SendToTpl() {
	this.Data["Time"] = this.GetSession("Time")
	this.Data["Userid"] = this.GetSession("Userid")
	this.Data["Username"] = this.GetSession("Username")
	this.Data["Password"] = this.GetSession("Password")
	this.Data["Role"] = this.GetSession("Role")
	this.Data["Email"] = this.GetSession("Email")
}
