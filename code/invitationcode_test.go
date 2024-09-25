package code

import (
	"fmt"
	"testing"
)

func TestInvatitionCode(t *testing.T) {
	m := make(map[string]uint64)
	set := make(map[uint64]struct{})

	for i := 1; i <= 1000; i++ {
		uid := uint64(i)
		set[uid] = struct{}{}
		code := GetInvitationCode(uid)
		fmt.Println("uid:", uid, "code:", code)
		if v, ok := m[code]; ok {
			t.Errorf("old uid: %d, new uid: %d, have same code: %s", v, uid, code)
		}
	}

	fmt.Println("total uid:", len(set))
}

func TestInvationCode2Uid(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		uid := uint64(i)
		code := GetInvitationCode(uid)
		v, ok := InvitationCode2Uid(code)
		if !ok {
			t.Errorf("code: %s, uid: %d, decode failed", code, uid)
		}
		if v != uid {
			t.Errorf("code: %s, uid: %d, decode failed, got: %d", code, uid, v)
		}
	}
}
