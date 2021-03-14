package port

import (
	"bytes"

	"github.com/a-cordier/goose/proto"
)

const Code proto.UInt = 2

func Write(port proto.UInt) []byte {
	buf := new(bytes.Buffer)
	proto.WriteUInt(buf, Code)
	return proto.Pack(buf.Bytes())
}
