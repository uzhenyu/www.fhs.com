package model

import "zg5/z311/framework/mysql"

func AutoTable() error {
	return mysql.DB.AutoMigrate(new(Shop))
}
