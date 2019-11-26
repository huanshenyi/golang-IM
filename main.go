package main

import (
	_ "github.com/go-sql-driver/mysql"
	"hello/controller"
	"html/template"
	"log"
	"net/http"
)


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

	mux.HandleFunc("/user/login", controller.UserLogin)
	mux.HandleFunc("/user/register", controller.UserRegister)
	mux.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	mux.HandleFunc("/contact/loadfriend", controller.LoadFriend)
	mux.HandleFunc("/contact/joincommunity", controller.JoinCommunity)
	//http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	mux.HandleFunc("/contact/addfriend", controller.Addfriend)
	mux.HandleFunc("/chat", controller.Chat)

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