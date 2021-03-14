package login

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"io"
	"net"

	"github.com/a-cordier/goose/proto"
)

const (
	Code proto.UInt = 1
)

type Response interface {
	OK() bool
}

type Success struct {
	Greet string
	IP    net.IP
	Sum   string
}

func readIp(val proto.UInt) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, uint32(val))
	return ip
}

func ReadSuccess(reader io.Reader) Success {
	greet := proto.ReadString(reader)
	ip := readIp(proto.ReadUInt(reader))
	sum := proto.ReadString(reader)
	return Success{greet, ip, sum}
}

func (response Success) OK() bool {
	return true
}

type Failure struct {
	Reason string
}

func ReadFailure(reader io.Reader) Failure {
	reason := proto.ReadString(reader)
	return Failure{reason}
}

func (response Failure) OK() bool {
	return false
}

func Write(username string, password string) []byte {
	buf := new(bytes.Buffer)
	proto.WriteUInt(buf, Code)
	binary.Write(buf, binary.LittleEndian, proto.NewString(username))
	binary.Write(buf, binary.LittleEndian, proto.NewString(password))
	binary.Write(buf, binary.LittleEndian, proto.MajorVersion)
	binary.Write(buf, binary.LittleEndian, sum(username, password))
	binary.Write(buf, binary.LittleEndian, proto.MinorVersion)
	return proto.Pack(buf.Bytes())
}

func sum(username string, password string) proto.String {
	sum := md5.Sum([]byte(username + password))
	return proto.NewString(hex.EncodeToString(sum[:]))
}

func Read(reader io.Reader) Response {
	proto.ReadUInt(reader) // size
	proto.ReadUInt(reader) // code 1
	success := proto.ReadBool(reader)
	if success {
		return ReadSuccess(reader)
	}
	return ReadFailure(reader)
}
