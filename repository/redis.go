package repository

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

/*note
-Redis Do 是執行指令
 若get value ,回傳的值為byte ,需轉換為string ; 若是不存在的key,則回傳nil
 經由redis.String ,轉換型態後 , nil會變成空字串

 */
func Conn() {
	c, _ := redis.Dial("tcp", ":6379")
	defer c.Close()
	//v, _ := redis.String(c.Do("get", "test"))
	aa, _ := c.Do("get", "aabb")
	fmt.Println(aa == nil)
	return
}
