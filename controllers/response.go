/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/10 9:38
 */

package controllers

import "FurbotServer-Go/models"

type getFursuitResponse struct {
	models.FursuitTable
	Url   string `json:"url"`
	Thumb string `json:"thumb"`
}
