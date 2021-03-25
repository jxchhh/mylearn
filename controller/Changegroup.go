package controller

import (
	"ACAT/common"
	"ACAT/model"
	"ACAT/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"fmt"
	"github.com/gin-gonic/gin"
)
//@Summary Changegroup
//@Description 更改组别
//@Tags operatechangegroup
//@Accept json
//@Produce json
//@Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
//@Param groupname  body  string  true "发送的json数据只有groupname字段"
//@Success 200 {string} json "{"code":200,"msg":"更改成功","data":nil}"
//@Failure 403 {string} json "{"code":403,"msg":"已经处于面试状态不能更改组别","data":nil}"
//@Router /api/user/changegroup [POST]
type QUEUE struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
    Data int `json:"data"`
}

func getqueue(uid interface{}) int{
	var post QUEUE
	c := &http.Client{}
    userid := uid.(string)
	str := "1"
	url := "http://192.168.1.197:8000/queue/peekQueueIndexByUserId/"+userid+str
	req,err:= http.NewRequest("GET",url,nil)
	if err != nil {
		fmt.Println(err)
	}
	resp,er :=c.Do(req)
	if er != nil {
		fmt.Println(er)
	}

	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body,&post)
	return post.Data}
func Changegroup(c *gin.Context){

    // fmt.Println(c)
	var gid int
	var student model.User
	db := common.InitDB()
	defer db.Close()
	var post model.Changegroup
	stuid, _:= c.Get("studentid")
	db.Table("tb_user").Where("id = ?", stuid).First(&student)
	status := getqueue(stuid)
    if status <0{
	if student.Itw_status==1||student.Itw_status==0{
	c.ShouldBind(&post)
	if post.Groupname== "Java后端" {
		gid =1
	}
	if post.Groupname== "Go后端" {
		gid =2
	}
	if post.Groupname== "Web前端" {
		gid =3
	}
	if post.Groupname== "机器学习" {
		gid =4
	}
	if post.Groupname== "服务端" {
		gid =5
	}
 // db.Table("tb_user").Where("id = ?",stuid).Update("group_id",gid)
	db.Table("tb_user").Where("id = ?", stuid).Updates(map[string]interface{}{"group_id": gid})
 //fmt.Println(result.Error)
	    response.Success(c, nil, "更改成功")} else {
		response.Fail(c,nil,"已经处于面试状态不能更改组别")
		return
	}
} else {
	response.Fail(c,nil,"当前正在面试中不能更改组别")
	}
}
