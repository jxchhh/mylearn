package model

import "time"

type User struct {
	Id         int  `json:"id"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Phone string  `json:"phone"`
	Email string  `json:"email"`
	Real_name string `json:"real_name"`
	Sex int `json:"sex"`
	Major string  `json:"major"`
	User_pic string  `json:"user_pic"`
	Group_id    int `json:"group_id"`
	Status     int  `json:"status"`
	Itw_status int  `json:"itw_status"`

}
type Message struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
	Data string `json:"data"`
}
type User1 struct {
	Username   string  `json:"username"`
	Password   string  `json:"password"` //密码
	Email string  `json:"email"`//邮箱
	Real_name string `json:"real_name"` //真实姓名
	Major string `json:"major"`//专业
	Status     int `json:"status"`//鉴别账号可用！
}
type Code struct {
	Email string `json:email`
}
type Post struct { //解析前端传来的json数据 登录用
	Username string `json:"username"`
	Password string `json:"password"`
}
type Group struct {//报名！
	Groupname string `json:"groupname"`
    Telephone string `json:"telephone"`
}
type Changegroup struct {
	Groupname string `json:groupname`
}
type Role struct {
	User_id interface{}
	Role_id int
   Status string
	Created time.Time
}
type Verify struct {
	Email string `json:"email"`
	Capture string `json:"capture"`
}
type Getinformation struct {
	Message string `json:"msg"`
	Code int `json:code`
	Data bool `json:"data"`
}