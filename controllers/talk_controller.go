package controllers

import (
	"encoding/json"

	"github.com/Yitsushi/lightningtalk/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TalkController struct {
	beego.Controller
}

func (this *TalkController) Get() {
	var talks []*models.Talk
	o := orm.NewOrm()
	o.QueryTable("talk").All(&talks)

	this.Data["json"] = talks
	this.ServeJSON()
}

func (this *TalkController) Find() {
	var id int64

	id, _ = this.GetInt64(":id")

	this.Data["json"] = struct {
		Test string
		Id   int64
	}{"test", id}
	this.ServeJSON()
}

func (this *TalkController) Post() {
	response := struct {
		Message string
		Status  bool
		Talk    *models.Talk
	}{"Success", true, &models.Talk{}}

	var err error

	err = json.Unmarshal(this.Ctx.Input.RequestBody, response.Talk)
	if err == nil {
		response.Talk.Id = 0
		response.Talk.Presenter = nil

		if response.Talk.Title != "" {
			models.TalkInsert(response.Talk)
		} else {
			response.Message = "Empty Title!"
			response.Status = false
		}

	} else {
		response.Message = err.Error()
		response.Status = false
	}

	this.Data["json"] = response

	this.ServeJSON()
}

func (this *TalkController) Update() {
	this.Data["json"] = "Updated..."
	this.ServeJSON()
}
