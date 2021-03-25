package middleware

import (
	"ACAT/common"
	"ACAT/model"
	"ACAT/response"
	"github.com/gin-gonic/gin"
)

func Sendmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post model.Code
		c.ShouldBind(&post)
		var verify model.User
		db := common.InitDB()
		defer db.Close()
		db.Table("tb_user").Where("email = ?",post.Email).First(&verify)
		if verify.Id==0{
			flag:=1
			response.Success(c,gin.H{"data": flag},"可以注册")
			c.Set("email",post.Email)
			c.Next()
		} else {
			response.Fail(c,nil,"邮箱已经注册过了")
			c.Abort()
			return
		}
	}
}
