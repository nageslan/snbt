package snbt

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"os"
)

// Decode parses SNBT data to the provided value of a pointer s.
func Decode(r []byte, s interface{}) error {
	lxr := newLexer(r)
	if err := lxr.tokenize(); err != nil {
		return err
	}
	psr := newParser(lxr.tokens())
	dmap, err := psr.parse()
	if err != nil && err != ErrExhaustedAllTokens {
		return err
	}
	return mapstructure.Decode(dmap, s)
}

// DecodeByString parses SNBT data to the provided value of a pointer s.
func DecodeByString(str string, s interface{}) error {
	return Decode([]byte(str), s)
}

// DecodeByFile parses SNBT data to the provided value of a pointer s.
func DecodeByFile(filepath string, s interface{}) error {
	context, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	return Decode(context, s)
}

// DecodeToMap parses SNBT to map
func DecodeToMap(r []byte) (map[string]interface{}, error) {
	var m1 = map[string]interface{}{}
	err := Decode(r, &m1)
	if err != nil {
		return nil, err
	}
	return m1, nil
}

// DecodeToJson parses SNBT to json
func DecodeToJson(r []byte) (string, error) {
	var m1 = map[string]interface{}{}
	err := Decode(r, &m1)
	if err != nil {
		return "", err
	}
	marshal, err := json.Marshal(m1)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
