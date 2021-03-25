package controller

import (
	"ACAT/common"
	_ "ACAT/common"
	"ACAT/model"
	"ACAT/response"
	"github.com/gin-gonic/gin"
	"time"
)
// @Summary Sighup
// @Description 报名
// @Tags operatesighup
// @Accept json
// @Produce json
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param groupname  body  string true "发送的json数据分别是 groupname和telephone"
// @Param telephone  body  string true "发送的json数据分别是 groupname和telephone"
// @Success 200 {string} json "{"code":200,"msg":"报名成功","data":nil}"
// @Failure 403 {string} json "{"code":403,"msg":"已经报过不能重复报名","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"服务器出错报名失败","data":nil}"
// @Router /api/user/sighup [POST]
func Sighup(c *gin.Context){
   // fmt.Println(c)
   var gid int
	var post model.Group
	var student model.User

	db := common.InitDB()
	defer db.Close()
	stuid, _:= c.Get("studentid")
	db.Table("tb_user").Where("id = ?", stuid).First(&student)
	if student.Group_id !=0{
		response.Fail(c,nil,"已经报过名不能重复报名")
		return
	} else {
		c.ShouldBind(&post)
		gname := post.Groupname
		tele := post.Telephone
    //fmt.Println(gname)
	if gname == "Java后端" {
			gid = 1
		}
		if gname == "Go后端" {
			gid = 2
		}
		if gname == "Web前端" {
			gid = 3
		}
		if  gname== "机器学习" {
			gid = 4
		}
		if gname == "服务端" {
			gid = 5
		}

		result :=db.Table("tb_user").Where("id = ?", stuid).Updates(map[string]interface{}{ "phone":tele,"group_id": gid,"itw_status":1}) //看一下表名
		if result.Error==nil{response.Success(c, nil, "报名成功")} else {
			response.Fail(c,nil,"服务器出错报名失败")
		}
		user := model.Role{User_id: stuid,Role_id:1,Status: "1",Created: time.Now()}
		db.Table("tb_user_role").Create(&user)

	}
}
