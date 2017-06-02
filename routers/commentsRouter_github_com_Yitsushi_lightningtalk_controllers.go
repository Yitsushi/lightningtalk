package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/Yitsushi/lightningtalk/controllers:TalkController"] = append(beego.GlobalControllerRouter["github.com/Yitsushi/lightningtalk/controllers:TalkController"],
		beego.ControllerComments{
			"Find",
			`/talk/:id`,
			[]string{"get"},
			nil})

}
