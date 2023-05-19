package snbt

import (
	"fmt"
	"regexp"
)

type tokenType int

const (
	tkIllegal tokenType = iota
	tkStr
	tkInt
	tkFloat
	tkArrStart
	tkArrEnd
	tkMapStart
	tkMapEnd
	tkBool // false true
)

var displayTokenType = []string{
	"ILL",
	"STR",
	"INT",
	"FLO",
	"ARRS",
	"ARRE",
	"MAPS",
	"MAPE",
}

var endSigBytes = []byte{
	':',
	',',
	'b',
	'd',
	'f',
	's',
	'\n',
}

var specialEndSigBytes = []byte{
	']',
	'}',
}

var emptyBytes = []byte{
	' ',
	',',
	'\n',
	'\t',
	'\r',
	':',
	']',
	'}',
	';',
}
var emptyBytes2 = []byte{
	' ',
	',',
	'\n',
	'\t',
	'\r',
}

// 匹配空字符
var emptyBytesRe = regexp.MustCompile(`(?m)\n|\t| |\r|L|,`)

func bytesContain(bs []byte, tb byte) bool {
	for _, b := range bs {
		if b == tb {
			return true
		}
	}
	return false
}

type token struct {
	t tokenType
	v interface{}
}

func (t token) string() string {
	if t.v == nil {
		return displayTokenType[t.t]
	}
	return fmt.Sprintf("%s:%v", displayTokenType[t.t], t.v)
}
