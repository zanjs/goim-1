package connect

import (
	"net"
	"encoding/binary"
)

type Codec struct {
	buffer     buffer // 接收消息的buffer
	typeBuf    []byte // 数据类型存储字节
	lenBuf     []byte // 数据库所占字节数存储字节
	valueType  int    // 数据类型
	valueLen   int    // 数据内容长度
	valueBuf   []byte // 数据内容buffer
	typeLen    int    // 数据类型长度
	lenLen     int    // 数据长度字节长度
	typeLenLen int    // 数据类型+数据长度 长度
}

// newCodec 创建一个解码器
func NewCodec(conn net.Conn, bufferLen int, typeLen int, lenLen int) *Codec {
	c := Codec{
		buffer:     newBuffer(conn, bufferLen),
		typeLenLen: typeLen + lenLen,
	}
	return &c
}

// read 从conn里面读取数据，当conn发生阻塞，这个方法也会阻塞
func (c *Codec) Read() (int, error) {
	return c.buffer.readFromReader()
}

// decode 解码数据
func (c *Codec) Decode() (*Message, bool) {
	var err error
	// 读取数据类型
	c.typeBuf, err = c.buffer.seek(0, c.typeLen)
	if err != nil {
		return nil, false
	}

	// 读取数据长度
	c.lenBuf, err = c.buffer.seek(c.typeLen, c.lenLen)
	if err != nil {
		return nil, false
	}

	// 读取数据内容
	c.valueType = int(binary.BigEndian.Uint16(c.typeBuf))
	c.valueLen = int(binary.BigEndian.Uint16(c.lenBuf))

	c.valueBuf, err = c.buffer.read(c.typeLenLen, c.valueLen)
	if err != nil {
		message := Message{Code: c.valueType, Content: c.valueBuf}
		return &message, true
	}
	return nil, false
}

func (c *Codec)Eecode(message Message)error{
	sendBuf:=make([]byte,c.typeLenLen+len(message.Content))



}
