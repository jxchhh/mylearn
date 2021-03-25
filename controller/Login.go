package controller

import (
	"ACAT/common"
	"ACAT/model"
	"ACAT/response"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// @Summary Login
// @Description 用户登录
// @Tags operateLogin
// @Accept json
// @Produce json
// @Param Authorization	header string true "Bearer 31a1656dec616b1f8f3207b4273"
// @Param username formData string true "学号"
// @Param password formData string true "用户密码"
// @Success 200 {string} json "{"code":200,"msg":"登陆成功","data":token}"
// @Failure 403 {string} json "{"code":403,"msg":"信息错误","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"授权失败","data":nil}"
// @Router /api/user/login [POST]
type Java struct {
	Src string `json:"src"`
	Code string `json:"code"`
}

func postjava(information *Java) int {
	var post model.Getinformation
    json2, _ := json.Marshal(&information)
	req, err := http.NewRequest("POST", "http://192.168.1.197:8100/user/matching",bytes.NewBuffer(json2))//不能用newreader
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		fmt.Println(err)
	}
	c := &http.Client{}
	resp, er := c.Do(req)
	if er != nil {
		fmt.Println(er)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &post)

	if post.Data == true {
		return 1
	}else
	{return 0}

}


func UserLogin(c *gin.Context) {
	db := common.InitDB()
	defer db.Close()
	var flag int
	var post model.Post
	c.ShouldBind(&post)
	studentid := post.Username
	var s model.User
	db.Table("tb_user").Where("username=?",studentid).First(&s)
    if s.Id<=0 {
		response.Fail(c,nil,"用户不存在")
		return
	} else {
		var javaobject Java
		javaobject.Src = post.Password
		javaobject.Code = s.Password

		msg := postjava(&javaobject)
		if msg ==1{
        if s.Group_id!=0 {
        	mytoken, _ := common.Createmw(s)
        	flag=1
			response.Success(c, gin.H{"token": mytoken,"flag" : flag,"info" :s}, "登录成功")} else{
			mytoken, _ := common.Createmw(s)
			flag =0
			response.Success(c, gin.H{"token": mytoken,"flag" : flag}, "登录成功")//报名
		}
	}else{
			response.Fail(c,nil,"密码错误")
		}
	}
}