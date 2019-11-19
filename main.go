package main

import (
	"encoding/json"
	"log"
	"net/http"
)
//curl http://127.0.0.1:8080/user/login -X POST -d "mobile=18600000000&passwd=123456"

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


func main()  {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/login", userLogin)

	server := http.Server{
		Addr:"127.0.0.1:8080",
		Handler:mux,
	}
	server.ListenAndServe()

}