package common

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var Redisconnect *redis.Conn
func Initconn() redis.Conn {
	Redisconnect ,err:= redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Println(err)
		return nil
	}
		return Redisconnect

}
