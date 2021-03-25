package controller

import (
	"ACAT/common"
	"ACAT/model"
	"ACAT/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sync"
)
// @Summary register
// @Description 注册
// @Tags operateregister
// @Accept json
// @Produce json
// @Param username  body  string true "发送的json数据--学号"
// @Param password  body  string true "发送的json数据--密码"
// @Param email  body  string true "发送的json数据--邮箱"
// @Param real_name  body  string true "发送的json数据--真实姓名"
// @Param major  body  string true "发送的json数据--专业"
// @Param status  body  int true "验证码成功之后直接发1"
// @Success 200 {string} json "{"code":200,"msg":"注册成功","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"服务器出错注册失败","data":nil}"
// @Failure 403 {string} json "{"code":403,"msg":"已经注测过不能重复注册","data":nil}"
// @Router /api/user/register [POST]
type Queue struct {
	Msg string `json:"msg"`
	Code int `json:"code"`

}
func getcode(usercode string) string{
	var post model.Message

	c := &http.Client{}
    url := "http://192.168.1.197:8100/user/encode/"+usercode
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
func Register(c *gin.Context){
	var lock sync.Mutex
	var registerdata model.User1
    var registerdata2 model.User
	c.ShouldBind(&registerdata)
	fmt.Println(registerdata.Password)
	registerdata.Password = getcode(registerdata.Password)
	fmt.Println(registerdata)
    db := common.InitDB()
    defer  db.Close()
	lock.Lock()
   db.Table("tb_user").Where("username=?",registerdata.Username).First(&registerdata2)
    if registerdata2.Id==0{
		result:= db.Table("tb_user").Create(&registerdata)

		if result.Error != nil {
			response.Fail(c,nil,"注册失败")
		}else {
			response.Success(c,nil,"注册成功")}
		} else {
		  response.Fail(c,nil,"已经注册过不能重复注册")
	}
	lock.Unlock()
}
