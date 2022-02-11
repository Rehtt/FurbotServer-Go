/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/9 21:06
 */

package middleware

import (
	"FurbotServer-Go/models"
	"github.com/gin-gonic/gin"
)

func VisitorAuth(ctx *gin.Context) {
	qq := ctx.Query("qq")
	timestamp := ctx.Query("timestamp")
	sign := ctx.Query("sign")
	apiPath := ctx.Request.URL.Path[1:]
	if !models.VisitorAuth(qq, timestamp, sign, apiPath) {
		ctx.Status(401)
		ctx.Abort()
		return
	}
	ctx.Next()
}

func AdminAuth(ctx *gin.Context) {
	auth := ctx.GetHeader("admin-auth")
	if !models.AdminAuth(auth) {
		ctx.Status(401)
		ctx.Abort()
		return
	}
	ctx.Next()
}
