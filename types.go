package gorand

import (
	"math/rand"
	"strings"
	"unicode/utf8"
)

func Rune() rune {
	for {
		r := rand.Int31n(utf8.MaxRune)
		if utf8.ValidRune(r) {
			return r
		}
	}
}

func String(len int) string {
	var str strings.Builder
	for i:=0; i< len ;i++{
		str.WriteRune(Rune())
	}
	return str.String()
}
