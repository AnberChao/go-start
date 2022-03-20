package middleware

import (
	"go-start/framework"
	"log"
	"time"
)

func Cost() framework.ControllerHandler {
	// 使用函数回掉
	return func(c *framework.Context) error {
		//记录开始时间
		start := time.Now()

		log.Printf("api uri start:%v", c.GetRequest().RequestURI)
		// 使用next执行具体的业务逻辑
		c.Next()

		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Panicf("api uri end:%v,cost:%v", c.GetRequest().RequestURI, cost.Seconds())

		return nil
	}
}
