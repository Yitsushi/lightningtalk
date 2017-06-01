package models

type Talk struct {
	Id          uint64 `orm:"auto"`
	Title       string `orm:"size(100)"`
	Description string `orm:"size(500)"`
	Presenter   *User  `orm:"null;rel(one);on_delete(set_null)"`
}
