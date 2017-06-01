package main

import (
	"github.com/Yitsushi/lightningtalk/controllers"
	"github.com/Yitsushi/lightningtalk/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(
		new(models.User),
		new(models.Talk),
	)
	orm.RegisterDataBase("default", "mysql", "root:123456@/lightningtalk?charset=utf8", 30)

	orm.RunSyncdb("default", true, true)
}

func setupRouter() {
	beego.Router("/", &controllers.FrontpageController{})
}

func main() {
	setupRouter()
	beego.Run()
}
