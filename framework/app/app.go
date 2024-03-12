package app

import (
	"zg5/z311/framework/mysql"
)

func Init(fileName string, apps ...string) error {
	var err error
	for _, v := range apps {
		switch v {
		case "mysql":
			err = mysql.InitMysql(fileName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
