package models

type User struct {
	Id       uint64 `orm:"auto"`
	Nickname string `orm:"size(100)"`
}
