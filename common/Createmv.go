package common

import (
	"ACAT/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)


type Claims struct {
	Userid int
	jwt.StandardClaims
}

var jwtkey = []byte("hhh")
func Createmw(user model.User) (string,error){


	expireTime := time.Now().Add(7*24*time.Hour)
	claims :=&Claims{
		Userid: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "jxc",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenstring, err := token.SignedString(jwtkey)
	if err !=nil {
		//fmt.Println(claims.Userid)
		fmt.Println("加密失败",err)
		return "", err
	}
	return tokenstring, nil
}
func ParseToken(tokenstring string) (*jwt.Token,*Claims,error) { //解析加密后的令牌,并返回未加密的令牌

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenstring, claims,func(token *jwt.Token) (i interface{},err error){
		return jwtkey,nil
	})
	return token,claims,err
}



