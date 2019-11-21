package service

import (
	"errors"
	"fmt"
	"hello/model"
	"hello/util"
	"math/rand"
	"time"
)

type UserService struct {

}

//新規アカウント
func (s *UserService)Register(
	mobile, // 携帯
	plainpwd, // 明文パスワード
	nickname, // ユーザー名
	avatar, // アイコン
	sex string) (user model.User, err error) {
	//モバイル存在するか調べる
	tmp := model.User{}
    _, err = DbEngin.Where("mobile=?",mobile).Get(&tmp)
    if err != nil{
    	return tmp, err
	}
	//もし存在すれば,[すでにアカウント存在する]と返す
    if tmp.Id > 0 {
    	return tmp,errors.New("該当モバイル存在します")
	}
	//存在しなければ、新規アカウント作成
    tmp.Mobile = mobile
    tmp.Avatar = avatar
    tmp.Nickname = nickname
    tmp.Sex = sex

	// passwd =
	// md5暗号化
    tmp.Salt = fmt.Sprint("%06d", rand.Int31n(10000))
    tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
    tmp.Created_at = time.Now()
    //token
    tmp.Token = fmt.Sprintf("%08d",rand.Int31())

    //dbに挿入
    _,err = DbEngin.InsertOne(&tmp)
    if err != nil{
       return tmp, err
	}

	//新しいユーザー情報を返す
	return tmp, nil
}

//ログイン
func (s *UserService) Login (mobile,plainpwd string) (user model.User, err error) {
	return user, nil
}
