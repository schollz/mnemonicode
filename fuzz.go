// For use with go-fuzz, "github.com/dvyukov/go-fuzz"
//
// +build gofuzz

package mnemonicode

import (
	"bytes"
	"fmt"

	"golang.org/x/text/transform"
)

var (
	tenc    = NewEncodeTransformer(nil)
	tdec    = NewDecodeTransformer()
	tencdec = transform.Chain(tenc, tdec)
)

// For use with go-fuzz, "github.com/dvyukov/go-fuzz"
func Fuzz(data []byte) int {
	words := EncodeWordList(nil, data)
	if len(words) != WordsRequired(len(data)) {
		panic("bad WordsRequired result")
	}
	data2, err := DecodeWordList(nil, words)
	if err != nil {
		fmt.Println("words:", words)
		panic(err)
	}
	if !bytes.Equal(data, data2) {
		fmt.Println("words:", words)
		panic("data != data2")
	}

	tencdec.Reset()
	data3, _, err := transform.Bytes(tencdec, data)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(data, data3) {
		fmt.Println("words:", words)
		panic("data != data3")
	}

	if len(data) == 0 {
		return 0
	}
	return 1
}
