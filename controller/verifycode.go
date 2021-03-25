package controller

import (
	"ACAT/common"
	"ACAT/model"
	"ACAT/response"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)
// @Summary verifycode
// @Description 验证验证码是否正确
// @Tags operateverifycode
// @Accept json
// @Produce json
// @Param  email body  string true "接收的json数据email"
// @Param  capture body  string true "接收的json数据capture验证码"
// @Success 200 {string} json "{"code":200,"msg":"验证成功","data":nil}"
// @Failure 403 {string} json "{"code":403,"msg":"验证失败","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"服务器出错验证失败","data":nil}"
// @Router /api/user/verifycode [POST]
func Verifycode(c *gin.Context){
	var post model.Verify
     var flag int
	c.ShouldBind(&post)
	fmt.Println(post)
	conn := common.Initconn()
	defer conn.Close()
	code ,err := redis.String(conn.Do("GET",post.Email))
	if err != nil {
		fmt.Println("写入redis失败",err)
		response.Fail(c,nil,"服务器出错")
		return
	}
	if code == post.Capture{
         flag=1
         fmt.Println(flag)
		response.Success(c,gin.H{"data":flag},"验证成功")
	} else {
          flag=0
		response.Fail(c,gin.H{"data":flag},"验证失败")
	}
}
