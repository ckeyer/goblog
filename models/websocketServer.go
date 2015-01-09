package models

import (
	"bufio"
	"container/list"
	"io"
	"log"
	"net/websocket"
	"strconv"
)

var WSChatClients list.List

type WSChatClient struct {
	ws      *websocket.Conn
	item_ws *list.Element
}

func NewWSChatClient(ws *websocket.Conn) (w *WSChatClient) {
	w = &WSChatClient{}
	w.ws = ws
	w.item_ws = WSChatClients.PushBack(w)
	return
}
func (w *WSChatClient) Close() {
	w.SendCloseMessage()
	WSChatClients.Remove(w.item_ws)
	w.ws.Close()
}
func (w *WSChatClient) SendMessage(self *list.Element, data string) {
	for item := WSChatClients.Front(); item != nil; item = item.Next() {
		wc, ok := item.Value.(*WSChatClient)
		if !ok {
			panic("item not *websocket.Conn")
		}
		if item == self {
			continue
		}
		io.WriteString(wc.ws, data)
	}
}

func ChatroomServer(ws *websocket.Conn) {
	w := NewWSChatClient(ws)
	defer w.Close()

	r := bufio.NewReader(ws)
	log.Println("Connected ")
	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		log.Println("Received: " + string(data))
		rmsg := DecodeJson(string(data))
		if rmsg.Code == "close" {
			break
		} else {
			w.controlMsg(rmsg)
		}

		// msg := NewMsg()
		// msg.Code = "HeroServer"
		// log.Println("Send Will Start")

		// log.Println(WSChatClients.Len())
		// w.SendMessage(nil, msg.ToBase64String())
		// log.Println("Send Will Over")
		// log.Println(`{"Code":"online_user_list","Data": "` + ch.OnlineListToBase64() + `"}`)
		// var msg Msg
		// err = json.Unmarshal(data, &msg)
		// if err != nil {
		// 	SendMessage(item_ws, `{"Code":"error","Data":"JSON Error","Desc":"Nothing"}`)
		// }
	}
	log.Println("Listenning Over")
}
func (wc *WSChatClient) controlMsg(rmsg *Message) {
	msg := NewMsg()
	switch rmsg.Code {
	case "conn":
		wc.SendWelcomeMessage()
	case "msg":
		msg = rmsg
		wc.SendMessage(wc.item_ws, msg.ToBase64String())
	}
}
func (w *WSChatClient) SendCloseMessage() {
	msg := NewMsg()
	msg.Code = "msg"
	msg.Data = "One go away"
	w.SendMessage(w.item_ws, msg.ToBase64String())
	msg.Code = "online_user_count"
	msg.Data = strconv.Itoa(WSChatClients.Len() - 1)
	w.SendMessage(w.item_ws, msg.ToBase64String())
}
func (w *WSChatClient) SendWelcomeMessage() {
	msg := NewMsg()
	msg.Code = "msg"
	msg.Data = "One joined"
	w.SendMessage(nil, msg.ToBase64String())
	msg.Code = "online_user_count"
	msg.Data = strconv.Itoa(WSChatClients.Len())
	w.SendMessage(nil, msg.ToBase64String())
}
