package data

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/hyansource/20200212gameclient/pb"
)

func TestA(t *testing.T) {

	//序列化

	p := &pb.Data01{
		Name:  "ABCDE",
		Money: 10000,
	}

	binary_data, err := proto.Marshal(p)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p)
	fmt.Println(binary_data)

	//反序列化
	headbuf := bytes.NewReader(binary_data)

	head := &Message{}

	if err = binary.Read(headbuf, binary.LittleEndian, &head.Len); err != nil {
		return
	}

	if err = binary.Read(headbuf, binary.LittleEndian, &head.MsgId); err != nil {
		return
	}

	//head.Data = binary_data

	fmt.Println(head)
}
