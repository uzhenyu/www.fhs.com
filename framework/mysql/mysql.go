package mysql

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zg5/z311/framework/nacos"
)

type T struct {
	Mysql struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
		Host     string `json:"Host"`
		Port     string `json:"Port"`
		Database string `json:"Database"`
	} `json:"Mysql"`
}

var DB *gorm.DB

func InitMysql(fileName string) error {
	config, err := nacos.GetConfig(fileName)
	if err != nil {
		return err
	}
	cnf := new(T)
	err = json.Unmarshal([]byte(config), &cnf)
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cnf.Mysql.Username,
		cnf.Mysql.Password,
		cnf.Mysql.Host,
		cnf.Mysql.Port,
		cnf.Mysql.Database,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func WithTX(txFc func(tx *gorm.DB) error) {
	var err error
	tx := DB.Begin()
	err = txFc(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}
