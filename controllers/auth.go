/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/9 21:00
 */

package controllers

import (
	"FurbotServer-Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuthList(ctx *gin.Context) {
	data := models.GetVisitorAuth(nil)
	ctx.JSON(http.StatusOK, data)
}

func AddAuth(ctx *gin.Context) {
	var request AuthRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if ok := models.AddVisitorAuth(request.QQ, request.AuthKey); !ok {
		ctx.String(http.StatusBadRequest, "QQ重复")
	}
	ctx.Status(http.StatusOK)
}

func FixAuth(ctx *gin.Context) {
	var request AuthRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if ok := models.FixVisitorAuth(request.QQ, request.AuthKey); !ok {
		ctx.String(http.StatusBadRequest, "找不到QQ")
	}
	ctx.Status(http.StatusOK)
}

func DeleteAuth(ctx *gin.Context) {
	models.DeleteVisitorAuth(ctx.Param("qq"))
	ctx.Status(http.StatusOK)
}
