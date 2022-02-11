/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/9 21:26
 */

package models

import (
	"FurbotServer-Go/extends/util"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

// 验证
func VisitorAuth(qq, timestamp, sign, apiPath string) bool {
	t, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return false
	}
	// 时间相差1分钟
	if time.Now().Unix()-t > 60 {
		return false
	}

	if util.VisitorAuth(apiPath, timestamp, GetAuthKeyFromQQ(qq)) != sign {
		return false
	}
	return true
}

// 验证
func AdminAuth(auth string) bool {
	if util.AdminAuth(strings.Split(auth, ",")[0], viper.GetString("adminKey")) == auth {
		return true
	}
	return false
}
