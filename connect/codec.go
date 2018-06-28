package connect

import (
	"encoding/binary"
	"net"
)

// CodecConfig
type CodecFactory struct {
	TypeLen   int // 消息type字节数组的长度
	LenLen    int // 消息length字节数组的长度
	HeadLen   int // 消息type字节数组的长度 + 消息length字节数组的长度
	BufferLen int // buffer大小,建议不小于最大消息体的字节长度
}

func NewCodecFactory(typeLen, lenLen, bufferLen int) *CodecFactory {
	return &CodecFactory{
		TypeLen:   typeLen,
		LenLen:    lenLen,
		HeadLen:   typeLen + lenLen,
		BufferLen: bufferLen,
	}
}

type Codec struct {
	Factory *CodecFactory
	conn    net.Conn
	buffer  buffer // 接收消息的buffer
	//typeBuf    []byte // 数据类型存储字节
	//lenBuf     []byte // 数据所占字节数存储字节
	// valueType  int    // 数据类型
	// valueLen   int    // 数据内容长度
	// valueBuf   []byte // 数据内容buffer
	// typeLen    int    // 数据类型长度
	// lenLen     int    // 数据长度字节长度
	// typeLenLen int    // 数据类型+数据长度 长度
}

// newCodec 创建一个解码器
func (c *CodecFactory) NewCodec(conn net.Conn) *Codec {
	return &Codec{
		buffer:  newBuffer(conn, c.BufferLen),
		conn:    conn,
		Factory: c,
	}
}

// Read 从conn里面读取数据，当conn发生阻塞，这个方法也会阻塞
func (c *Codec) Read() (int, error) {
	return c.buffer.readFromReader()
}

// Decode 解码数据
func (c *Codec) Decode() (*Message, bool) {
	var err error
	// 读取数据类型
	typeBuf, err := c.buffer.seek(0, c.Factory.TypeLen)
	if err != nil {
		return nil, false
	}

	// 读取数据长度
	lenBuf, err := c.buffer.seek(c.Factory.TypeLen, c.Factory.LenLen)
	if err != nil {
		return nil, false
	}

	// 读取数据内容
	valueType := int(binary.BigEndian.Uint64(typeBuf))
	valueLen := int(binary.BigEndian.Uint64(lenBuf))

	valueBuf, err := c.buffer.read(c.Factory.HeadLen, valueLen)
	if err != nil {
		message := Message{Code: valueType, Content: valueBuf}
		return &message, true
	}
	return nil, false
}

// Eecode 编码数据
func (c *Codec) Eecode(message Message) error {
	contentLen := len(message.Content)
	buf := make([]byte, c.Factory.HeadLen+contentLen)

	binary.BigEndian.PutUint64(buf[0:c.Factory.TypeLen], uint64(message.Code))
	binary.BigEndian.PutUint64(buf[c.Factory.LenLen:c.Factory.HeadLen], uint64(message.Code))
	copy(buf[c.Factory.HeadLen:], message.Content)

	_, err := c.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
