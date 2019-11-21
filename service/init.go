package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"hello/model"
	"log"
)

var DbEngin *xorm.Engine

func init()  {
	drivename := "mysql"
	DsName := "root:root@(127.0.0.01:3306)/chat?charset=utf8" //ユーザーネーム:パスワード@(port)/dbネーム?charset=utf8
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename,DsName)
	if err != nil && err.Error() != ""{
		log.Fatal(err.Error())
	}
	// 操作中にsql表示するかどうか
	DbEngin.ShowSQL(true)
	// データーベースのリンク数
	DbEngin.SetMaxOpenConns(2)

	//自動でテーブル作る
	DbEngin.Sync2(new(model.User))

	fmt.Println("init data base ok")
}