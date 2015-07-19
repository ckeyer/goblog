package controllers

import (
	"bufio"
	"container/list"
	"github.com/ckeyer/goblog/models"
	"golang.org/x/net/websocket"
	"io"
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
	log.Info("Connected ")
	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		log.Info("Received: " + string(data))
		rmsg, _ := models.DecodeJson(data)
		if rmsg.Code == "close" {
			break
		} else {
			w.controlMsg(rmsg)
		}
	}
	log.Info("Listenning Over")
}
func (wc *WSChatClient) controlMsg(rmsg *models.Message) {
	msg := models.NewMsg()
	switch rmsg.Code {
	case "conn":
		wc.SendWelcomeMessage()
	case "msg":
		msg = rmsg
		wc.SendMessage(wc.item_ws, msg.ToBase64String())
	}
}
func (w *WSChatClient) SendCloseMessage() {
	msg := models.NewMsg()
	msg.Code = "msg"
	msg.Data = "One go away"
	w.SendMessage(w.item_ws, msg.ToBase64String())
	msg.Code = "online_user_count"
	msg.Data = strconv.Itoa(WSChatClients.Len() - 1)
	w.SendMessage(w.item_ws, msg.ToBase64String())
}
func (w *WSChatClient) SendWelcomeMessage() {
	msg := models.NewMsg()
	msg.Code = "msg"
	msg.Data = "One joined"
	w.SendMessage(nil, msg.ToBase64String())
	msg.Code = "online_user_count"
	msg.Data = strconv.Itoa(WSChatClients.Len())
	w.SendMessage(nil, msg.ToBase64String())
}
