package users

import "hello/controllers"

type UserController struct {
	controllers.BaseController
}

func (this *UserController) Get() {
	this.BaseController.Verify(1)
	this.BaseController.SendToTpl()
	this.TplName = "user.tpl"
}
