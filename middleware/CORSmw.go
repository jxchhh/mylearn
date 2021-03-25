package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*func CORSmw() gin.HandlerFunc {
	return func(ctx *gin.Context) {

			method := c.Request.Method
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")


		fmt.Println("hhhh")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}

}*/
func  CORSmw() gin.HandlerFunc {
return func(c *gin.Context) {
method := c.Request.Method
c.Header("Access-Control-Allow-Origin", "*")
c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
c.Header("Access-Control-Allow-Credentials", "true")

//放行所有OPTIONS方法
if method == "OPTIONS" {
c.AbortWithStatus(http.StatusNoContent)
}
// 处理请求
c.Next()
}
}


