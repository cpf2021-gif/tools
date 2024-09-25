package code

import "strings"

const Chars = "23456789ABCDEFGH"
const CodeLength = 6
const Salt = 233
const Prime = 7
const Prime2 = 5

const Base = uint64(len(Chars))

// GetInvatationCode 获取邀请码
func GetInvitationCode(uid uint64) string {
	var b [CodeLength]uint64
	var res string

	pid := uid*Prime + Salt
	b[0] = pid

	// 扩散
	for i := 0; i < CodeLength-1; i++ {
		// 为什么要 + uint64(i) * b[0] ?
		// 为了让个位的值影响到其他位，增加随机性，不然相邻uid的邀请码会很接近
		b[i] = (b[i] + uint64(i)*b[0]) % Base
		b[i+1] = b[i] / Base
	}

	for i := 0; i < CodeLength-1; i++ {
		b[CodeLength-1] += b[i]
	}

	b[CodeLength-1] = b[CodeLength-1] * Prime % CodeLength

	// 混淆
	for i := 0; i < CodeLength; i++ {
		res += string(Chars[b[(i*Prime2)%CodeLength]])
	}

	return res
}

func InvitationCode2Uid(code string) (uint64, bool) {
	if len(code) != CodeLength {
		return 0, false
	}

	var b [CodeLength]uint64

	// 反混淆
	for i := 0; i < CodeLength; i++ {
		b[(i*Prime2)%CodeLength] = uint64(i)
	}

	for i := 0; i < CodeLength; i++ {
		j := strings.Index(Chars, string(code[b[i]]))
		if j == -1 {
			return 0, false
		}
		b[i] = uint64(j)
	}

	// 检验
	var expected uint64
	for i := 0; i < CodeLength-1; i++ {
		expected += b[i]
	}
	expected = expected * Prime % CodeLength
	if b[CodeLength-1] != expected {
		return 0, false
	}

	// 反扩散
	for i := CodeLength - 2; i >= 0; i-- {
		b[i] = (b[i] - uint64(i)*(b[0]-Base)) % Base
	}

	var res uint64
	for i := CodeLength - 2; i > 0; i-- {
		res += b[i]
		res *= Base
	}

	res = (res + b[0] - Salt) / Prime

	return res, true
}
