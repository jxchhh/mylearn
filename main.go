package main

import (
	"ACAT/common"
	"github.com/gin-gonic/gin"
)
//@title Golang interview API
//@version 1.0
//@descrition use gin doing login&sighup&changegroup
//@contact.name API jxc
//@contact.email 1308570778@qq.com
//@host 127.0.0.1:8080
func main(){
	r := gin.Default()
	r = InitRouter(r)
 	db := common.InitDB()
	defer db.Close()
    r.Run(":8080")
}
