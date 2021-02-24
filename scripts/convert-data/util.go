package main

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tealeg/xlsx"
	"golang.org/x/text/unicode/norm"
)

// extract root directory of this project
var _, _path, _, _ = runtime.Caller(0)
var rootPath = filepath.Join(filepath.Dir(_path), "../..")

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type Rows []*xlsx.Row

func (rs Rows) Index(i int) *xlsx.Row {
	if i >= len(rs) {
		return nil
	}
	return rs[i]
}

type Cells []*xlsx.Cell

func (cs Cells) Index(i int) *xlsx.Cell {
	if i >= len(cs) {
		return nil
	}
	return cs[i]
}

func (cs Cells) ByValue(v string) *xlsx.Cell {
	for _, c := range cs {
		if strings.TrimSpace(c.Value) == v {
			return c
		}
	}
	return nil
}

const (
	vneseChars = "đĐ" +
		"àáạảãâầấậẩẫăằắặẳẵ" +
		"ÀÁẠẢÃÂẦẤẬẨẪĂẰẮẶẲẴ" +
		"èéẹẻẽêềếệểễ" +
		"ÈÉẸẺẼÊỀẾỆỂỄ" +
		"òóọỏõôồốộổỗơờớợởỡ" +
		"ÒÓỌỎÕÔỒỐỘỔỖƠỜỚỢỞỠ" +
		"ùúụủũưừứựửữ" +
		"ÙÚỤỦŨƯỪỨỰỬỮ" +
		"ìíịỉĩ" + "ỳýỵỷỹ" +
		"ÌÍỊỈĨ" + "ỲÝỴỶỸ"

	numChars   = "0123456789"
	upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars = "abcdefghijklmnopqrstuvwxyz"
	signChars  = ` .,/\"'_-+=@#%*()[]{}<>!?$&`
	nameChars  = signChars + numChars + upperChars + lowerChars + vneseChars
)

var (
	vneseMap map[rune]byte

	specialChars = []bool{
		'(': true,
		')': true,
		'[': true,
		']': true,
		'{': true,
		'}': true,
		'<': true,
		'>': true,
		'/': true,

		'!': true,
		'@': true,
		'#': true,
		'$': true,
		'%': true,
		'^': true,
		'&': true,
		'*': true,
		'-': true,
		'_': true,
		'+': true,
		'=': true,

		'.': true,
		',': true,
		':': true,
		';': true,
		'?': true,
		'|': true,
	}
)

func init() {
	vneseMap = make(map[rune]byte)
	initVneseMap("đĐ", 'd')
	initVneseMap("àáạảãâầấậẩẫăằắặẳẵ", 'a')
	initVneseMap("ÀÁẠẢÃÂẦẤẬẨẪĂẰẮẶẲẴ", 'a')
	initVneseMap("èéẹẻẽêềếệểễ", 'e')
	initVneseMap("ÈÉẸẺẼÊỀẾỆỂỄ", 'e')
	initVneseMap("òóọỏõôồốộổỗơờớợởỡ", 'o')
	initVneseMap("ÒÓỌỎÕÔỒỐỘỔỖƠỜỚỢỞỠ", 'o')
	initVneseMap("ùúụủũưừứựửữ", 'u')
	initVneseMap("ÙÚỤỦŨƯỪỨỰỬỮ", 'u')
	initVneseMap("ìíịỉĩÌÍỊỈĨ", 'i')
	initVneseMap("ỳýỵỷỹỲÝỴỶỸ", 'y')
}

func initVneseMap(s string, c byte) {
	for _, src := range s {
		vneseMap[src] = c
	}
}

// Keep alphanumeric and some special characters while ignoring the rest.
//
//    hello@world -> hello @ world
//    hello #@@@ world -> hello # @ @@@ world
//    hello(1) -> hello ( 1 )
//    hello.world -> hello . world
func normalizeSearch(s string, space string, quote bool, lower bool) string {
	var lastChar rune
	lastGroup := 0 // space
	b := make([]byte, 0, len(s))
	for _, c := range s {
		switch {
		case c >= '0' && c <= '9':
			if lastGroup == 2 {
				b = append(b, space...)
			}
			b = append(b, byte(c))
			lastGroup = 1 // numeric

		case c >= 'A' && c <= 'Z':
			if lastGroup == 1 {
				b = append(b, space...)
			}
			if lower {
				b = append(b, byte(c)+'a'-'A')
			} else {
				b = append(b, byte(c))
			}
			lastGroup = 2 // alpha

		case c >= 'a' && c <= 'z':
			if lastGroup == 1 {
				b = append(b, space...)
			}
			b = append(b, byte(c))
			lastGroup = 2 // alpha

		case vneseMap[c] != 0:
			if lastGroup == 1 {
				b = append(b, space...)
			}
			b = append(b, vneseMap[c])
			lastGroup = 2 // alpha

		default:
			if lastGroup != 0 {
				lastGroup = 0 // space
				b = append(b, space...)
			}
			if IsAcceptedSpecialChar(c) {
				if c == lastChar {
					continue
				}
				if quote {
					b = append(b, '\'', byte(c), '\'')
				} else {
					b = append(b, byte(c))
				}
				b = append(b, space...)
				lastGroup = 0 // space
			}
		}
		lastChar = c
	}
	if lastGroup == 0 && len(b) != 0 {
		b = b[:len(b)-len(space)]
	}
	return string(b)
}

func IsAcceptedSpecialChar(c rune) bool {
	return false
	// return int(c) < len(specialChars) && specialChars[c]
}

func NormalizeUnaccent(s string) string {
	s = norm.NFC.String(s)
	return normalizeSearch(s, "-", false, true)
}
