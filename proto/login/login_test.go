package login

import (
	"encoding/hex"
	"testing"
)

func TestLoginMessage(t *testing.T) {
	having := hex.EncodeToString(Write("username", "password"))
	expecting := "480000000100000008000000757365726e616d650800000070617373776f7264a000000020000000643531633961376539333533373436613630323066393630326434353239323901000000"
	if having != expecting {
		t.Fail()
	}
}
