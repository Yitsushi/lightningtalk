package models

type User struct {
	ID       uint64 `orm:"auto"`
	Nickname string `orm:"size(100)"`
}
