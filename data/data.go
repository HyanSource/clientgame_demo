package data

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hyansource/20200212gameclient/pb"
)

type Message struct {
	Len   uint32
	MsgId uint32
	Data  []byte
}

//玩家的socket
type TcpClient struct {
	conn net.Conn
	id   int32
}

//建構子
func NewTcpClient(ip string, port int) *TcpClient {

	addr := fmt.Sprintf("%s:%d", ip, port)

	conn, err := net.Dial("tcp", addr)
	checkErr(err)

	client := &TcpClient{
		conn: conn,
	}
	return client
}

func (t *TcpClient) Start() {
	//下注業務
	go func() {
		for {
			time.Sleep(5 * time.Second)

			//下注
			p := &pb.Bet{
				PlayerID: t.id,
				Bet:      50,
			}

			t.SendMsg(100, p)

			// //廣播
			// talk := &pb.BroadCast{
			// 	PlayerID: t.id,
			// 	TalkMsg:  &pb.TalkMsg{Content: "你好"},
			// }

			// t.SendMsg(150, talk)
		}
	}()

	//接收訊息
	go func() {
		for {
			headdata := make([]byte, 8)
			if _, err := io.ReadFull(t.conn, headdata); err != nil {
				fmt.Println(err)
				return
			}

			pkghead, err := t.UnPack(headdata)

			if err != nil {
				return
			}

			if pkghead.Len > 0 {
				pkghead.Data = make([]byte, pkghead.Len)
				if _, err := io.ReadFull(t.conn, pkghead.Data); err != nil {
					return
				}
			}

			//處理業務 暫時先用顯示處理
			t.DoMsg(pkghead)
		}
	}()

	fmt.Println("OK")
}

//拆包 (這邊沒有寫入Message Data的動作 需要看一下有沒有BUG)
func (t *TcpClient) UnPack(headdata []byte) (head *Message, err error) {
	headbuf := bytes.NewReader(headdata)

	head = &Message{}

	//讀取Len
	if err = binary.Read(headbuf, binary.LittleEndian, &head.Len); err != nil {
		return nil, err
	}

	if err = binary.Read(headbuf, binary.LittleEndian, &head.MsgId); err != nil {
		return nil, err
	}

	return head, nil
}

//裝包
func (t *TcpClient) Pack(msgID uint32, dataBytes []byte) (out []byte, err error) {

	outbuff := bytes.NewBuffer([]byte{})
	//寫Len
	if err = binary.Write(outbuff, binary.LittleEndian, uint32(len(dataBytes))); err != nil {
		return
	}

	//寫msgid
	if err = binary.Write(outbuff, binary.LittleEndian, msgID); err != nil {
		return
	}

	//寫data
	if err = binary.Write(outbuff, binary.LittleEndian, dataBytes); err != nil {
		return
	}

	out = outbuff.Bytes()

	return
}

//傳送訊息
func (t *TcpClient) SendMsg(msgID uint32, data proto.Message) {

	//進行編碼
	binary_data, err := proto.Marshal(data)

	if err != nil {
		fmt.Println("marshaling err:", err)
		return
	}

	sendData, err := t.Pack(msgID, binary_data)

	if err != nil {
		fmt.Println(err)
	}

	t.conn.Write(sendData)
	return
}

//處理業務
func (t *TcpClient) DoMsg(msg *Message) {

	switch msg.MsgId {
	case 1: //登入業務
		{
			tempdata := &pb.PlayerData{}
			proto.Unmarshal(msg.Data, tempdata)
			t.id = tempdata.PlayerID
		}
		break
	case 100: //下注業務
		{
			tempdata := &pb.ReturnGameResult{}
			proto.Unmarshal(msg.Data, tempdata)
			fmt.Println(tempdata)
		}
		break
	case 150: //廣播業務
		{
			tempdata := &pb.BroadCast{}
			proto.Unmarshal(msg.Data, tempdata)
			fmt.Println(tempdata.TalkMsg.Content)
		}
		break
	case 200:
		{

		}
		break
	case 201: //玩家下線
		break
	default:
		break
	}

}

//check
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
