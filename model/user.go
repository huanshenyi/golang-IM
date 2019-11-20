package model


import "time"

//model.SEX_WOMEN
const (
	SEX_WOMEN="W"
	SEX_MAN="M"
	SEX_UNKNOW="U"
)

type User struct {
	Id         int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"` //携帯番号
	Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`   // ユーザーパスワード(plainpwd+salt),MD5
	Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"` //logo
	Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`   //
	Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`   // ユーザーネーム
	Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`   // 暗号化用の要素,ランダム数
	Online     int	`xorm:"int(10)" form:"online" json:"online"`   //
	Token      string	`xorm:"varchar(40)" form:"token" json:"token"`   //token chat?id=1&token=x
	Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`   //
	Created_at   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   //新規ユーザーの量を統計用
	Updated_at   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   //
}