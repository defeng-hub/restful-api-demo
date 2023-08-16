package middleware

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/user/common"
	"restful-api-demo/apps/user/common/response"
	"restful-api-demo/apps/user/service"
)

var casbinService = service.CasbinService{}

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//value, exists := c.Get("user")

		waitUse, _ := common.GetClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
