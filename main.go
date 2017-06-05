package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yitsushi/lightningtalk/controllers"
	"github.com/Yitsushi/lightningtalk/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	runmode := os.Getenv("RUNMODE")
	if runmode == "" {
		runmode = "prod"
	}
	beego.AppConfig.Set("runmode", runmode)

	orm.RegisterModel(
		new(models.User),
		new(models.Talk),
	)
	orm.RegisterDataBase("default", "mysql", "root:123456@/lightningtalk?charset=utf8", 30)

	orm.RunSyncdb("default", true, false)
}

func getAllFileContent(path string, extension string) []byte {
	final := []byte{}

	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), extension) {
			continue
		}

		content, _ := ioutil.ReadFile(filepath.Join(path, file.Name()))
		final = append(final, content...)
	}

	return final
}

func setupRouter() {
	beego.Router("/", &controllers.FrontpageController{})
	beego.Router("/talk", &controllers.TalkController{})
	beego.Router("/talk/:id", &controllers.TalkController{}, "get:Find")
	beego.Get("/static/js/app.js", func(ctx *context.Context) {
		content := []byte{}
		content = append(content, getAllFileContent("./js-src/helpers", ".js")...)
		content = append(content, getAllFileContent("./js-src/modules", ".js")...)
		content = append(content, getAllFileContent("./js-src", ".js")...)

		if beego.AppConfig.String("runmode") == "prod" {
			ioutil.WriteFile("./static/js/app.js", content, 0644)
		}

		ctx.Output.Header("Content-Type", "application/javascript; charset=utf-8")
		ctx.Output.Body(content)
	})
}

func main() {
	setupRouter()
	beego.Run()
}
