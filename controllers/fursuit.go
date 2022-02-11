/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/9 18:31
 */

package controllers

import (
	"FurbotServer-Go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strings"
)

// GetFursuitImage 获取fursuit图片
func GetFursuitImage(ctx *gin.Context) {
	f, err := os.Open(fmt.Sprintf("%s/%s", viper.GetString("imagePath"), ctx.Param("img")))
	if err != nil {
		ctx.Abort()
		return
	}
	defer f.Close()
	ctx.Status(http.StatusOK)
	ctx.Header("Content-Type", "image/jpeg")
	io.Copy(ctx.Writer, f)
}

// GetFursuitRand 随机获取
func GetFursuitRand(ctx *gin.Context) {
	res := models.GetFursuitRand()
	returnFursuitResponse(ctx, res)
}

// GetFursuitByID 根据id获取
func GetFursuitByID(ctx *gin.Context) {
	id := ctx.Query("fid")
	res := models.GetFursuitByID(id)
	returnFursuitResponse(ctx, res)
}

// GetFursuitByName 根据名字获取
func GetFursuitByName(ctx *gin.Context) {
	name := ctx.Query("name")
	res := models.GetFursuitByName(name)
	returnFursuitResponse(ctx, res)
}

func returnFursuitResponse(ctx *gin.Context, fur models.FursuitTable) {
	u := strings.Split(ctx.Request.RequestURI, "/")
	data := getFursuitResponse{}
	data.FursuitTable = fur

	if fur.Model != nil {
		data.URL = fmt.Sprintf("http://%s%s/image/%d.jpg", ctx.Request.Host, strings.Join(u[:len(u)-1], "/"), fur.Fid)

	}
	ctx.JSON(http.StatusOK, data)
}

// AddFursuit 添加
func AddFursuit(ctx *gin.Context) {
	var fur FursuitRequest
	if err := ctx.ShouldBindJSON(&fur); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.AddFursuit(fur.Name, fur.Image); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
}

// DeleteFursuit 删除
func DeleteFursuit(ctx *gin.Context) {
	models.DeleteFursuit(ctx.Param("fid"))
}
