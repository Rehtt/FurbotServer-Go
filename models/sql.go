/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/9 21:29
 */

package models

import (
	"FurbotServer-Go/extends/sql"
	"FurbotServer-Go/extends/util"
	"github.com/spf13/viper"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"strconv"
	"time"
)

func GetAuthKeyFromQQ(qq string) string {
	var authKey string
	sql.DB.Self.Model(&AuthTable{}).Where(map[string]interface{}{"qq": qq}).Pluck("auth_key", &authKey)
	return authKey
}

func GetFursuitRand() (fur FursuitTable) {
	var count int64
	sql.DB.Self.Model(&FursuitTable{}).Count(&count)
	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(int(count))
	sql.DB.Self.Offset(offset).First(&fur)
	return
}

func GetFursuitById(fid string) (fur FursuitTable) {
	sql.DB.Self.Where(map[string]interface{}{"fid": fid}).Find(&fur)
	return
}

func GetFursuitByName(name string) (fur FursuitTable) {
	sql.DB.Self.Where(map[string]interface{}{"name": name}).Find(&fur)
	return
}

func GetVisitorAuth(query interface{}) (auth []AuthTable) {
	querySQL := sql.DB.Self.Model(&AuthTable{})
	if query != nil {
		querySQL.Where(query)
	}
	querySQL.Find(&auth)
	return
}

func DeleteVisitorAuth(qq string) {
	sql.DB.Self.Where(map[string]interface{}{"qq": qq}).Delete(&AuthTable{})
}

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
func FixVisitorAuth(qq, authKey string) bool {
	auth := GetVisitorAuth(map[string]interface{}{"qq": qq})
	if len(auth) == 0 {
		return false
	}
	sql.DB.Self.Model(&AuthTable{}).Where(map[string]interface{}{"qq": qq}).Update("auth_key", authKey)
	return true
}

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

func DeleteFursuit(fid string) {
	sql.DB.Self.Where(map[string]interface{}{"fid": fid}).Delete(&FursuitTable{})
}
