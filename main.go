package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hello/model"
	"hello/service"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)
//curl http://127.0.0.1:8080/user/login -X POST -d "mobile=18600000000&passwd=123456"
//curl http://127.0.0.1:8080/user/login -X POST -d "mobile=18600000000&passwd=123456"

// ログイン
func userLogin(w http.ResponseWriter, r *http.Request)  {
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
		Resp(w, 0, user, "ログインしました")
	}else {
		Resp(w, -1, nil, "パスワードかアカウント違う")
	}
}

var userService service.UserService

// 新規アカウント
func userRegister(w http.ResponseWriter, r *http.Request)  {
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
		Resp(w,-1, nil, err.Error())
	}else {
		Resp(w,0, user, err.Error())
	}
}

type H struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string)  {
	// headerをapplication/json  デフォルト:text/html
	// set header
	w.Header().Set("Content-Type", "application/json")
	// set status=200
	w.WriteHeader(http.StatusOK)
	// structを定義
	h := H{
      Code:code,
      Msg:msg,
      Data:data,
	}
	// structをjsonにする
	jsonData, err := json.Marshal(h)
	if err != nil{
		log.Println(err)
		return
	}else {
		// 出力
		w.Write(jsonData)
	}
}

func RegisterView(mux *http.ServeMux) {
	tpl,err := template.ParseGlob("view/**/*")
	//エラー出たら実行停止
	if err != nil {
		//printして終わり
		log.Fatal(err.Error())
	}
	for _,v := range tpl.Templates(){
	   tplname := v.Name()
		mux.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
          tpl.ExecuteTemplate(writer, tplname, nil)
	   })
	}
}

func main()  {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/login", userLogin)
	mux.HandleFunc("/user/register", userRegister)

	// 1.startファイルのディレクトリアクセス許可
	mux.Handle("/asset/",http.FileServer(http.Dir(".")))

	//mux.HandleFunc("/user/login.shtml", func(writer http.ResponseWriter, request *http.Request) {
	//	tpl, err := template.ParseFiles("view/user/login.html")
	//	if err!= nil{
	//		log.Fatal(err.Error())
	//	}
	//	tpl.ExecuteTemplate(writer,"/user/login.shtml","")
	//})

	RegisterView(mux)
	server := http.Server{
		Addr:"127.0.0.1:8080",
		Handler:mux,
	}
	server.ListenAndServe()

}