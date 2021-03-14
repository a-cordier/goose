package proto

import (
	"bytes"
	"encoding/binary"
	"io"
)

type (
	Int     int32
	UInt    uint32
	Long    int64
	ULong   uint64
	String  []byte
	Boolean uint8
)

const (
	MajorVersion UInt = 160
	MinorVersion UInt = 1
)

func Pack(data []byte) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, UInt(len(data)))
	binary.Write(buf, binary.LittleEndian, data)
	return buf.Bytes()
}

func ReadUInt(reader io.Reader) UInt {
	var val UInt
	binary.Read(reader, binary.LittleEndian, &val)
	return val
}

func WriteUInt(buf *bytes.Buffer, val UInt) {
	binary.Write(buf, binary.LittleEndian, val)
}

func NewString(content string) String {
	buf := new(bytes.Buffer)
	buf.WriteString(content)
	return Pack(buf.Bytes())
}

func ReadString(reader io.Reader) string {
	size := ReadUInt(reader)
	buf := make([]byte, size)
	io.ReadFull(reader, buf)
	return string(buf)
}

func ReadBool(reader io.Reader) bool {
	var val Boolean
	binary.Read(reader, binary.LittleEndian, &val)
	return val == 1
}
