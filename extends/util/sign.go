/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/9 21:18
 */

package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

// VisitorAuth 生成签名
func VisitorAuth(apiPath, timestamp, authKey string) string {
	w := md5.New()
	io.WriteString(w, fmt.Sprintf("%s-%s-%s", apiPath, timestamp, authKey))
	return fmt.Sprintf("%x", w.Sum(nil))
}

// AdminAuth 生成签名
func AdminAuth(timestamp, authKey string) string {
	w := md5.New()
	io.WriteString(w, fmt.Sprintf("admin-%s-%s", timestamp, authKey))
	return fmt.Sprintf("%s,%x", timestamp, w.Sum(nil))
}
