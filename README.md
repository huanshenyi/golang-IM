# db

```text
go get github.com/go-xorm/xorm
```

```text
go get github.com/go-sql-driver/mysql
```


# goの db操作

## CURD
xorm.NewSession(driverName,dataSourceName)
## modelの定義 entity

```text
type User struct {
    Id         int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
    Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
    Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`   // 什么角色
    Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
    Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`   // 什么角色
    Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`   // 什么角色
    Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`   // 什么角色
    Online     int	`xorm:"int(10)" form:"online" json:"online"`   //是否在线
    Token      string	`xorm:"varchar(40)" form:"token" json:"token"`   // 什么角色
    Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`   // 什么角色
    Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   // 什么角色
}
```

Getは一つでデータ、Find は複数のデータ

## とあるユーザーを探す
```text
DbEngin.ID(userId).Get(&User)
```

## ある種の条件満足する
```text
result := make([]User,0)
DbEngin.where("a=? and b=? ...",a,b).Find(&result)
DbEngin.where("modile=?",moile).Get(&User)
```

## create
```text
DBengin.InsertOne(&User)
```
## update

```text
DBengin.ID(userId).Update(&User)
// update ... where id = xx
DBengin.Where("a=? and b=?",a,b).Update(&User)
DBengin.ID(userId).Cols("nick_name").Update(&User) //一つのカラムを変更
DBengin.Where("a=? and b=?",a,b).Cols("nick_name").Update(&User)
```

## delete
  DBengin.ID(userId).Delete(&User)

## MD5  

```go
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	
	return hex.EncodeToString(cipherStr)
}

func MD5Encode(data string) string{
	return strings.ToUpper(Md5Encode(data))
}

func ValidatePasswd(plainpwd,salt,passwd string) bool{
	return Md5Encode(plainpwd+salt)==passwd
}
func MakePasswd(plainpwd,salt string) string{
	return Md5Encode(plainpwd+salt)
}

```
```text
docker-machine ip default
```

# ユーザー追加/ グループ追加
```text
/contact/addfriend フレンド追加、パラメータ userid,dstid

ユーザー10000追加10086の場合,contactテーブルに二つのデータ追加
//ownerid=10000,distobj=10086
//ownerid=10086,dstobj=10000

/contact/loadfriend フレンド表示,パラメータuserid

/contact/createcommunity グループ追加，アイコンpic,名前name,
/contact/loadcommunity グループ全部表示,パラメータuserid

/contact/joincommunity グループに参加、パラメータuserid,dstid
```
流れ

# 追加/表示友達, 追加/表示グループ
```text
code fence: cgo /contact/addfri...
```

# テーブル同期

```text
DbEngin.Sync2(new(model.Contact), new(model.Community))
```