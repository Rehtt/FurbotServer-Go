/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/10 16:26
 */

package controllers

type AuthRequest struct {
	QQ      string `json:"qq"`
	AuthKey string `json:"auth_key"`
}

type FursuitRequest struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
