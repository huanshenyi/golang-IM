package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespFail(w http.ResponseWriter, msg string)  {
	Resp(w, -1, nil,msg)
}

func RespOk(w http.ResponseWriter, data interface{}, msg string)  {
	Resp(w, 0, data,msg)
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
