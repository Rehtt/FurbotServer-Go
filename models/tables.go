package models

import "gorm.io/gorm"

// FursuitTable fursuit表
type FursuitTable struct {
	*gorm.Model `json:"-"`
	Name        string `json:"name"`
	Fid         int    `json:"id"`
}

// AuthTable auth表
type AuthTable struct {
	*gorm.Model `json:"-"`
	QQ          string `json:"qq"`
	AuthKey     string `json:"auth_key"`
}
