package models

import "github.com/astaxie/beego/orm"

type Talk struct {
	ID          uint64 `orm:"auto";form:"-"`
	Title       string `orm:"size(100)"`
	Description string `orm:"size(500)"`
	Presenter   *User  `orm:"null;rel(one);on_delete(set_null)"`
}

func TalkInsert(talk *Talk) *Talk {
	o := orm.NewOrm()
	id, _ := o.Insert(talk)

	talk.ID = uint64(id)

	return talk
}
