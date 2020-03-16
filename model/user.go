package model

import "time"

type User struct {
	Name     string
	Password string
	TaobaoSid      string
	JingdongId	string
	RegTime	time.Time
	ExTime time.Time
}