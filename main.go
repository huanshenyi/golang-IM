package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"html/template"
	"log"
	"net/http"
)
//curl http://127.0.0.1:8080/user/login -X POST -d "mobile=18600000000&passwd=123456"

var DbEngin *xorm.Engine

func init()  {
    drivename := "mysql"
    DsName := "root:root@(127.0.0.01:3306)/chat?charset=utf8" //ユーザーネーム:パスワード@(port)/dbネーム?charset=utf8
	DbEngin, err := xorm.NewEngine(drivename,DsName)
	if err != nil{
		log.Fatal(err.Error())
	}
	// 操作中にsql表示するかどうか
	DbEngin.ShowSQL(true)
	// データーベースのリンク数
	DbEngin.SetMaxOpenConns(2)

	//自動でテーブル作る
	//DbEngin.Sync2(new(User))

	fmt.Println("init data base ok")
}


func userLogin(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	moblie := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")
	loginok := false
	if (moblie == "18600000000" && passwd=="123456"){
		loginok = true
	}
	if (loginok){
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(w, 0, data, "")
	}else {
		Resp(w, -1, nil, "パスワード違う")
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