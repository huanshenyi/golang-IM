package controller

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//本核心在于形成userid和Node的映射关系
type Node struct {
	Conn *websocket.Conn
	//并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}

//映射关系表
var clientMap map[int64]*Node = make(map[int64]*Node,0)
//读写锁
var rwlocker sync.RWMutex

// ws://127.0.0.1/chat?id=1&token=xxxx
func Chat(w http.ResponseWriter, r *http.Request)  {
	// アクセスToken合法なのか
	// checkToken(userId int64,token string)
	query := r.URL.Query() //?id=1&token=xxxx を取得
	id := query.Get("id")
	token := query.Get("token")
	userId, err := strconv.ParseInt(id,10,64)
	if err != nil{
		log.Fatal(err.Error())
	}
	// tokenの有効性を判断
	isvalida := checkToken(userId,token)
	// if isvalida == true
	conn, err := (&websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(w,r,nil)

	if err != nil{
		log.Println(err.Error())
		return
	}

	//conn取得
	node := &Node{
		Conn:conn,
		DataQueue:make(chan []byte, 50),
		GroupSets:set.New(set.ThreadSafe),
	}

	// userid と node 関連つける
	rwlocker.Lock()
	clientMap[userId]=node
	rwlocker.Unlock()
	// todo 送信,con
	go sendproc(node)
	// todo 受け取り
	go recvproc(node)
	sendMsg(userId, []byte("hello,world"))

}

//送信並行処理
func sendproc(node *Node)  {
	for {
		select {
		case data := <- node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage,data)
			if err != nil{
				log.Println(err.Error())
				return
			}
		}
	}
}

//受け取り並行処理
func recvproc(node *Node)  {
	for{
		_, data, err := node.Conn.ReadMessage()
		if err != nil{
			log.Println(err.Error())
			return
		}
		//todo dataに対して更に処理する
		fmt.Printf("recv<=%s",data)
	}
}

//送信処理   userid=>誰に送る  msg=>何を送る
func sendMsg(userid int64, msg []byte)  {
	rwlocker.RLock()
	node, ok := clientMap[userid]
	rwlocker.RUnlock()
	if ok{
		node.DataQueue <- msg
	}
}

// tokenを検証
func checkToken(userId int64, token string) bool  {
	//dbから調べて比較する
	user := userService.Find(userId)
	return user.Token == token
}