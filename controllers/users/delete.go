package users

import (
	"hello/controllers"
	"hello/models"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type DeleteController struct {
	controllers.BaseController
}

func (this *DeleteController) Get() {
	this.BaseController.Verify(2)
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":userid"))
	orm.NewOrm().Delete(&models.User{Userid: id})
	this.Redirect("/admin", 302)
	return
}
