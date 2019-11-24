package model

import "time"
// 友達とグループ保存用
// 二つ分けてもいい
type Contact struct {
	Id         int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Ownerid       int64	`xorm:"bigint(20)" form:"ownerid" json:"ownerid"`   // 誰の
	Dstobj       int64	`xorm:"bigint(20)" form:"dstobj" json:"dstobj"`   // 相手は誰
	Cate      int	`xorm:"int(11)" form:"cate" json:"cate"`   // 友達追加かグループ追加なのか
	Memo    string	`xorm:"varchar(120)" form:"memo" json:"memo"`   // 什么角色
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   // 什么角色
}

const (
	CONCAT_CATE_USER = 0x01
	CONCAT_CATE_COMUNITY = 0x02
)
