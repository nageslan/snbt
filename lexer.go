package snbt

import (
	"bytes"
	"errors"
	"fmt"
	"unicode"
)

var (
	ErrIllegalChar = errors.New("illegal character")
)

type lexer struct {
	r   []byte
	cur int
	tks []token
}

func newLexer(r []byte) *lexer {
	return &lexer{
		r:   r,
		cur: -1,
		tks: []token{},
	}
}

func (lxr *lexer) next() bool {
	if lxr.cur+1 < len(lxr.r) {
		lxr.cur++
		return true
	}
	return false
}

func (lxr *lexer) curr() byte {
	return lxr.r[lxr.cur]
}

func (lxr *lexer) tokens() []token {
	return lxr.tks
}

func (lxr *lexer) tokenize() error {
	for lxr.next() {

		b := lxr.curr()
		if bytesContain(emptyBytes2, b) {
			continue
		}
		switch b {
		case '#':
			for lxr.next() {
				b2 := lxr.curr()
				if b2 == '\r' || b2 == '\n' {
					break
				}
			}
		case '{':
			lxr.appendToken(token{t: tkMapStart, v: "{"})
		case '}':
			lxr.appendToken(token{t: tkMapEnd, v: "}"})
		case '[':
			lxr.appendToken(token{t: tkArrStart, v: "["})
		case ']':
			lxr.appendToken(token{t: tkArrEnd, v: "]"})
		case ':':
		case ';':
		case '"':
			s, err := lxr.buildStr("", "\"")
			if err != nil {
				return err
			}
			lxr.appendToken(token{t: tkStr, v: s})
		case '\'':
			s, err := lxr.buildStr("", "'")
			if err != nil {
				return err
			}
			lxr.appendToken(token{t: tkStr, v: s})
		case '-':
			n, t, err := lxr.buildNum("", true)
			if err != nil {
				return err
			}
			lxr.appendToken(token{t: t, v: n})
		default:
			if b == 'f' {
				v, err := lxr.buildBool("f")
				if err == nil {
					lxr.next()
					b = lxr.curr()
					lxr.appendToken(token{t: tkBool, v: v})
				}
			}
			if b == 't' {
				v, err := lxr.buildBool("t")
				if err == nil {
					lxr.next()
					b = lxr.curr()
					lxr.appendToken(token{t: tkBool, v: v})
				}
			}
			if bytesContain(emptyBytes2, b) {
				continue
			}
			if b >= '0' && b <= '9' {
				n, t, err := lxr.buildNum(string(b), false)
				if err != nil {
					return err
				}
				lxr.appendToken(token{t: t, v: n})
				continue
			}
			if b <= unicode.MaxASCII {
				s, err := lxr.buildStr(string(b), ":")
				if err != nil {
					return err
				}
				lxr.appendToken(token{t: tkStr, v: s})
				continue
			}

			return fmt.Errorf("%s: %b", ErrIllegalChar, b)
		}
	}
	return nil
}

func (lxr *lexer) appendToken(t token) {
	lxr.tks = append(lxr.tks, t)
}

func (lxr *lexer) buildBool(start string) (bool, error) {
	str := []byte(start)
	tempIndex := lxr.cur
	for lxr.next() {
		b := lxr.curr()
		if emptyBytesRe.Match([]byte{b}) {
			break
		}
		str = append(str, b)
	}
	if string(str) == "true" {
		return true, nil
	} else if string(str) == "false" {
		return false, nil
	}
	// back
	lxr.cur = tempIndex
	return false, errors.New("not is bool")
}

var escapeCharacters = []byte{
	'n',
	'r',
	't',
	'"',
	'\'',
	'\\',
}

func (lxr *lexer) buildStr(start string, end string) (string, error) {
	str := []byte(start)
	for lxr.next() {
		b := lxr.curr()

		// The raw data for the player UUID is as follows: { UUID: [I; INT, INT, INT, INT] }
		// From the doc: "UUID of owner, stored as four ints.". NBT, why???
		if end == "\"" || end == "'" {
			if string(b) == "\\" {
				str = append(str, b)
				lxr.next()
				str = append(str, lxr.curr())
				continue
			}
			if string(b) == end {
				b2 := lxr.r[lxr.cur+1]
				if bytes.Contains(emptyBytes, []byte{b2}) {
					return string(str), nil
				}
			}
		} else {
			if string(b) == end || string(b) == ";" {
				b2 := lxr.r[lxr.cur+1]
				if bytes.Contains(emptyBytes, []byte{b2}) {
					return string(str), nil
				}

			}
		}
		str = append(str, b)

	}
	return string(str), nil

}

func (lxr *lexer) buildNum(start string, isSigned bool) (string, tokenType, error) {
	num := []byte(start)
	if isSigned {
		num = append([]byte{'-'}, num...)
	}
	t := tkInt
	for lxr.next() {
		b := lxr.curr()
		if bytesContain(specialEndSigBytes, b) {
			lxr.cur--
			break
		}
		if bytesContain(endSigBytes, b) {
			break
		}
		if b == '.' {
			if t == tkFloat {
				t = tkIllegal
			} else {
				t = tkFloat
			}
		}
		num = append(num, b)
	}
	num = emptyBytesRe.ReplaceAll(num, []byte(""))
	return string(num), t, nil
}
