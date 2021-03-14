package port

import (
	"encoding/hex"
	"testing"
)

func TestWritePort(t *testing.T) {
	having := hex.EncodeToString(Write(2234))
	expecting := "0400000002000000"
	if having != expecting {
		t.Fail()
	}
}
