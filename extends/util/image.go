/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/10 17:36
 */

package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"  // gif 格式
	_ "image/jpeg" // jpeg 格式
	_ "image/png"  // png 格式
	"io/ioutil"
	"strings"
)

// Base64ToImageBytes base64转图片[]byte
func Base64ToImageBytes(imageStr string) ([]byte, error) {
	if strings.Contains(imageStr, "base64,") {
		imageStr = strings.Split(imageStr, "base64,")[1]
	}
	res, err := base64.StdEncoding.DecodeString(imageStr)
	if err != nil {
		return nil, fmt.Errorf("base64解码失败")
	}
	_, _, err = image.Decode(bytes.NewReader(res))
	if err != nil {
		return nil, fmt.Errorf("仅支持jpg/png/gif格式")
	}
	return res, nil
}

// SaveFile 保存文件
func SaveFile(data []byte, path string) error {
	return ioutil.WriteFile(path, data, 644)
}
