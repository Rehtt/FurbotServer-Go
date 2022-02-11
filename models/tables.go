package models

import "gorm.io/gorm"

type FursuitTable struct {
	*gorm.Model `json:"-"`
	Name        string `json:"name"`
	Fid         int    `json:"id"`
}

type AuthTable struct {
	*gorm.Model `json:"-"`
	QQ          string `json:"qq"`
	AuthKey     string `json:"auth_key"`
}
