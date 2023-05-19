package snbt

import (
	"github.com/mitchellh/mapstructure"
)

// Decode parses SNBT data to the provided value of a pointer s.
func Decode(r []byte, s interface{}) error {
	//re := regexp.MustCompile(`^#.+|[ \n\t\r]#.*[^"][\n]+|#$`)
	////re := regexp.MustCompile(`^#.*\n`)
	//r = re.ReplaceAll(r, []byte(""))
	lxr := newLexer(r)
	if err := lxr.tokenize(); err != nil {
		return err
	}
	//for _, tk := range lxr.tks {
	//	fmt.Println("------------------------------------------")
	//	fmt.Println(tk.t, tk.v)
	//	fmt.Println("------------------------------------------")
	//}
	psr := newParser(lxr.tokens())
	dmap, err := psr.parse()
	if err != nil && err != ErrExhaustedAllTokens {
		return err
	}
	return mapstructure.Decode(dmap, s)
}
