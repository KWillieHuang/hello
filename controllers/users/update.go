package users

import (
	"hello/controllers"
	"hello/models"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

type UpdateController struct {
	controllers.BaseController
}

func (this *UpdateController) Get() {
	this.BaseController.Verify(2)
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":userid"))
	this.Data["Userid"] = id
	//if id = 0 ,information.tpl not have date
	if id != 0 {
		user := models.User{Userid: id}
		err := orm.NewOrm().Read(&user)
		if err == orm.ErrNoRows || err == orm.ErrMissPK {
			this.Redirect("/admin", 302)
			return
		} else {
			this.Data["Username"] = user.Username
			this.Data["Password"] = user.Password
			this.Data["Role"] = user.Role
			this.Data["Email"] = user.Email
		}
	}
	this.TplName = "information.tpl"
}

func (this *UpdateController) Post() {
	this.BaseController.Verify(2)
	id, _ := this.GetInt("Userid")
	beego.Info("I get the date")
	//if id = 0 ,user will be create
	//else user wil be update
	beego.Info("update")
	e := this.GetString("Email")
	if this.BaseController.IsEmail(e) {
		user := models.User{Userid: id}
		user.Username = this.GetString("Username")
		user.Password = controllers.Encode(this.GetString("Password"))
		user.Email = this.GetString("Email")
		user.Role, _ = this.GetInt("Role")
		_, err := orm.NewOrm().Update(&user)
		if err == nil {
			this.Redirect("/admin", 302)
			return
		} else {
			this.Redirect("/admin", 205)
			return
		}
	} else {
		this.Redirect("/admin", 205)
		return
	}
}
