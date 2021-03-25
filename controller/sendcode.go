package controller

import (
	"ACAT/common"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)
func Sendcode(c *gin.Context) {
	uemail ,_:=  c.Get("email")
    uuemail := uemail.(string)//interface转字符串
	captcha := common.Createcode()
     conn := common.Initconn()
    defer conn.Close()
    _, err := conn.Do("SET",uemail,captcha)
	if err != nil {
		fmt.Println(err)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "1308570778@qq.com")
	m.SetHeader("To",uuemail)
    body:= `<div style="background-color:#ECECEC; padding: 35px;">
    <table cellpadding="0" align="center"
        style="width: 600px; margin: 0px auto; text-align: left; position: relative; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 5px; border-bottom-left-radius: 5px; font-size: 14px; font-family:微软雅黑, 黑体; line-height: 1.5; box-shadow: rgb(153, 153, 153) 0px 0px 5px; border-collapse: collapse; background-position: initial initial; background-repeat: initial initial;background:#fff;">
        <tbody>
            <tr>
                <th valign="middle"
                    style="height: 25px; line-height: 25px; padding: 15px 35px; border-bottom-width: 1px; border-bottom-style: solid; border-bottom-color: #42a3d3; background-color: #49bcff; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px;">
                    <font face="微软雅黑" size="5" style="color: rgb(255, 255, 255); ">[ACAT计算机应用技术协会纳新系统]邮箱绑定</font>
                </th>
            </tr>
            <tr>
                <td>
                    <div style="padding:25px 35px 40px; background-color:#fff;">
                        <p>&nbsp;&nbsp;&nbsp;&nbsp;您收到这封这封电子邮件是因为您 (也可能是某人冒充您的名义) 申请了[ACAT计算机应用技术协会报名系统]邮箱绑定。若这不是您本人所申请,
                            请不用理会这封电子邮件。如有问题请联系管理员。此验证码10分钟内有效，请在有效时间内使用完成邮箱绑定：`+captcha+`<br>
                        <div style="width:700px;margin:0 auto;">
                            <div
                                style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;">
                                <p>此为系统邮件，请勿回复，如有问题请联系ACAT实验室负责人<br></p>
                                <p>©ACAT</p>
                            </div>
                        </div>
                    </div>
                </td>
            </tr>
        </tbody>
    </table>
</div>`
	m.SetHeader("Subject", "ACAT计算机应用技术协会报名系统") // 邮件标题
	m.SetBody("text/plain", body) // 邮件内容
	d := gomail.NewDialer("smtp.qq.com", 465, "1308570778@qq.com", "dbyaynjidpqzbagh")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}
