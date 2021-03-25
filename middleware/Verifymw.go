package middleware

import (
	"ACAT/common"
	"ACAT/model"
	"ACAT/response"
	"github.com/gin-gonic/gin"
	"strings"
)
func Verifymw() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Fail(c,nil,"请求头为空")
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader," ",2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {

			response.Fail(c,nil,"请求头格式有错")
			c.Abort()
			return
			
		}
		//处理正确tokenstring

		stringtoken, claims, err := common.ParseToken(parts[1])
		if err != nil || !stringtoken.Valid {

			response.Fail(c,nil,"无效token")
			c.Abort()
			return
		}
		userId := claims.Userid
		var student model.User
		db := common.InitDB()
	//	fmt.Println(userId)
	db.Table("tb_user").Where("id=?",userId).First(&student)
		defer db.Close()
     //  fmt.Println(student)
		if student.Id == 0 {
			response.Fail(c,nil,"用户不存在")
			c.Abort()
			return
		}
		c.Set("studentid",userId)//该学生存在,把该学生在表中对应的主键写入上下文
		//fmt.Println("hhhh")
	//	fmt.Println(c)
		c.Next()
	}
}