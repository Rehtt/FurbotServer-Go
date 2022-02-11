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
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"strings"
)

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

func SaveFile(data []byte, path string) error {
	return ioutil.WriteFile(path, data, 644)
}
