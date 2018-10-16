package routers

import (
	"hello/controllers/users"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &users.LoginController{})
	beego.Router("/user", &users.UserController{})
	beego.Router("/admin", &users.AdminController{})
	beego.Router("/admin/delete/?:userid", &users.DeleteController{})
	beego.Router("/admin/create", &users.CreateController{})
	beego.Router("/admin/update/?:userid", &users.UpdateController{})
	beego.Router("/resetemail", &users.ResetemailController{})
	beego.Router("/resetpwd", &users.ResetpwdController{})
}
