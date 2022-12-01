package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	ws   *websocket.Conn // websocket连接
	sc   chan []byte
	data *Data
}

var wu = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin:     func(r *http.Request) bool { return true }}

// 处理当前连接
func myws(w http.ResponseWriter, r *http.Request) {
	// 将HTTP服务器连接升级到WebSocket协议
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	// 获取当前连接对象
	c := &connection{sc: make(chan []byte, 256), ws: ws, data: &Data{}}
	// 把连接放到h.r中(读 管道)
	h.r <- c
	// 开一个协程用来 写
	go c.writer()
	// 执行 读 逻辑
	c.reader()
	// 断开连接后执行
	defer func() {
		c.data.Type = "logout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		data_b, _ := json.Marshal(c.data)
		h.b <- data_b
		h.r <- c
	}()
}

// 写逻辑，单开协程来写（返回客户端）
func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

var user_list = []string{} // 当前用户列表

func (c *connection) reader() {
	for {
		// 从websocket中读消息
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		// 把读到的数据反序列化，并存入data中
		json.Unmarshal(message, &c.data)
		fmt.Println("reader c.data is", c.data)
		switch c.data.Type {
		case "login": // 用户登录
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "user": // 用户发送消息
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "logout": // 用户下线
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			h.r <- c
		default:
			fmt.Print("========default================")
		}
	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(n_slice)
	return n_slice
}
