package controller

import (
	"fmt"
	"hello/model"
	"hello/service"
	"hello/util"
	"math/rand"
	"net/http"
)

//curl http://127.0.0.1:8080/user/login -X POST -d "mobile=18600000000&passwd=123456"
//curl http://127.0.0.1:8080/user/login -X POST -d "mobile=18600000000&passwd=123456"

// ログイン
func UserLogin(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	//
	mobile := r.PostForm.Get("mobile")
	//
	plainpwd := r.PostForm.Get("passwd")
	loginok := false
	user,err := userService.Login(mobile, plainpwd)
	if err != nil{
	}else {
		loginok = true
	}
	if (loginok){
		util.RespOk(w,user, "ログインしました")
	}else {
		util.RespFail(w,"パスワードかアカウント違う")
	}
}

var userService service.UserService

// 新規アカウント
func UserRegister(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	//
	mobile := r.PostForm.Get("mobile")
	//
	plainpwd := r.PostForm.Get("passwd")
	//
	nickname := fmt.Sprintf("user%06d",rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW
	user,err := userService.Register(mobile,plainpwd,nickname,avatar,sex)
	if err!= nil {
		util.Resp(w,-1, nil, err.Error())
	}else {
		util.Resp(w,0, user, err.Error())
	}
}
