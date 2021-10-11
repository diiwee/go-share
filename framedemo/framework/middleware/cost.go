package middleware

import (
	"framedemo/framework"
	"log"
	"time"
)

func Cost() framework.FrameHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		// 记录开始时间
		start := time.Now()

		// 使用next执行具体的业务逻辑
		c.Next()

		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("API: %v, COST: %v", c.GetRequest().RequestURI, cost.Seconds())

		return nil
	}
}
