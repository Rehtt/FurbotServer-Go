/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/9 21:29
 */

package models

import (
	"FurbotServer-Go/extends/sql"
	"FurbotServer-Go/extends/util"
	"github.com/spf13/viper"
	"math/rand"
	"strconv"
	"time"
)

// GetAuthKeyFromQQ 获取qq对应的auth key
func GetAuthKeyFromQQ(qq string) string {
	var authKey string
	sql.DB.Self.Model(&AuthTable{}).Where(map[string]interface{}{"qq": qq}).Pluck("auth_key", &authKey)
	return authKey
}

// GetFursuitRand 随机获取
func GetFursuitRand() (fur FursuitTable) {
	var count int64
	sql.DB.Self.Model(&FursuitTable{}).Count(&count)
	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(int(count))
	sql.DB.Self.Offset(offset).First(&fur)
	return
}

// GetFursuitByID 根据id获取
func GetFursuitByID(fid string) (fur FursuitTable) {
	sql.DB.Self.Where(map[string]interface{}{"fid": fid}).Find(&fur)
	return
}

// GetFursuitByName 根据名字获取
func GetFursuitByName(name string) (fur FursuitTable) {
	sql.DB.Self.Where(map[string]interface{}{"name": name}).Find(&fur)
	return
}

// GetVisitorAuth 根据条件获取auth
func GetVisitorAuth(query interface{}) (auth []AuthTable) {
	querySQL := sql.DB.Self.Model(&AuthTable{})
	if query != nil {
		querySQL.Where(query)
	}
	querySQL.Find(&auth)
	return
}

// DeleteVisitorAuth 删除auth
func DeleteVisitorAuth(qq string) {
	sql.DB.Self.Where(map[string]interface{}{"qq": qq}).Delete(&AuthTable{})
}

// AddVisitorAuth 添加auth
func AddVisitorAuth(qq, authKey string) bool {
	if len(GetVisitorAuth(map[string]interface{}{"qq": qq})) != 0 {
		return false
	}
	sql.DB.Self.Save(&AuthTable{
		QQ:      qq,
		AuthKey: authKey,
	})
	return true
}

// FixVisitorAuth 修改auth
func FixVisitorAuth(qq, authKey string) bool {
	auth := GetVisitorAuth(map[string]interface{}{"qq": qq})
	if len(auth) == 0 {
		return false
	}
	sql.DB.Self.Model(&AuthTable{}).Where(map[string]interface{}{"qq": qq}).Update("auth_key", authKey)
	return true
}

// AddFursuit 添加fursuit
func AddFursuit(name, imageStr string) error {
	var fid int
	sql.DB.Self.Model(&FursuitTable{}).Select("fid + 1 as fid").Pluck("fid", &fid)
	sql.DB.Self.Save(&FursuitTable{
		Name: name,
		Fid:  fid,
	})
	imageData, err := util.Base64ToImageBytes(imageStr)
	if err != nil {
		return err
	}
	return util.SaveFile(imageData, viper.GetString("imagePath")+"/"+strconv.Itoa(fid)+".jpg")
}

// DeleteFursuit 删除fursuit
func DeleteFursuit(fid string) {
	sql.DB.Self.Where(map[string]interface{}{"fid": fid}).Delete(&FursuitTable{})
}
