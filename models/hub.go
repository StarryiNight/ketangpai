package models

type Hub struct {
	Rooms     map[string]map[*Connection]string
	Broadcast chan Message
	Join      chan Message
	Quit      chan Message
	Eager     chan Message
	Choose    chan Message
	Answer    chan Message
}

var H = Hub{
	Rooms:     make(map[string]map[*Connection]string),
	Broadcast: make(chan Message),
	Join:      make(chan Message),
	Quit:      make(chan Message),
	Eager:     make(chan Message),
	Choose:    make(chan Message),
	Answer:    make(chan Message),
}

func (h *Hub) Run() {
	for true {
		select {
		//广播消息
		case m := <-h.Broadcast:
			h.sendAll(m)

		//进入聊天室
		case m := <-h.Join:
			conns := h.Rooms[m.Roomid]
			if conns == nil {
				conns = make(map[*Connection]string)
				h.Rooms[m.Roomid] = conns
			}
			h.Rooms[m.Roomid][m.Conn] = m.Username

			for con := range conns {
				str := "欢迎" + m.Username + "加入" + m.Roomid + "聊天室"
				msg := []byte(str)
				select {
				case con.Send <- msg:
				}

			}
			//退出聊天室
		case m := <-h.Quit:
			conns := h.Rooms[m.Roomid]
			if conns != nil {
				if _, ok := conns[m.Conn]; ok {
					delete(conns, m.Conn)
					close(m.Conn.Send)
					for con := range conns {
						str := m.Username + "离开了" + m.Roomid + "聊天室"
						msg := []byte(str)
						select {
						case con.Send <- msg:
						}
						if len(conns) == 0 {
							delete(h.Rooms, m.Roomid)
						}
					}
				}

			}
			//发出抢答问题
		case m := <-h.Eager:
			for true {
				h.sendAll(m)
				ans := <-H.Answer
				h.sendAll(ans)
				q := <-H.Eager
				str := string(q.Msg)
				if str == "正确" {
					h.sendAll(Message{
						Msg:      []byte("回答正确"),
						Roomid:   m.Roomid,
						Username: m.Username,
						Conn:     m.Conn,
					})
					H.Eager = make(chan Message)
					H.Answer = make(chan Message)
					break
				}
			}
		//抽问
		case m := <-h.Choose:
			//获取要抽问的人名字user
			user := string(m.Msg)
			//发送消息 提示被抽到的同学回答问题
			msg := "请" + user + "回答问题"
			m.Msg = []byte(msg)
			h.sendAll(m)
			//等待被抽问人回答
			for true {
				q := <-H.Answer
				if q.Username == user {
					h.sendAll(q)
					//获取老师的选择
					o := <-h.Choose
					option := string(o.Msg)
					h.sendAll(o)
					if option == "正确" || option == "取消" {
						break
					}
				}
			}

		}

	}
}

func (h *Hub) sendAll(m Message) {
	pre := []byte(m.Username + ":")
	m.Msg = append(pre, m.Msg...)
	conns := h.Rooms[m.Roomid]
	for con := range conns {
		if con == m.Conn {
			continue
		}
		select {
		case con.Send <- m.Msg:
		default:
			close(con.Send)
			delete(conns, con)
			if len(conns) == 0 {
				delete(h.Rooms, m.Roomid)
			}

		}
	}
}
