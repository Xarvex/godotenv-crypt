package godotenvcrypt

import (
	"bytes"
	"fmt"
	"unicode"
)

func nextStatement(src []byte) (int, int) {
	// single round of parsing, O(n)
	var comment bool
	index := bytes.IndexFunc(src, func(r rune) bool {
		if comment {
			comment = r != '\n' // end of line comment
		} else {
			comment = r == '#' // start of line comment
			return !comment && !unicode.IsSpace(r)
		}

		// even if comment is unset, would be the end of the line
		return false
	})
	if index == -1 {
		return -1, -1
	}

	return index, bytes.IndexAny(src[index:], "\r\n") // size of line
}

func eachStatement(src []byte, callback func([]byte)) {
	var size int
	for ; size != -1; src = src[size:] {
		var index int
		index, size = nextStatement(src)
		if index == -1 {
			return
		}

		src = src[index:]
		if size == -1 {
			callback(src)
		} else {
			callback(src[:size])
		}
	}
}

func environmentPair(src []byte) (string, string, error) {
	index := bytes.IndexRune(src, '=')

	key := bytes.TrimLeftFunc(bytes.TrimPrefix(bytes.TrimSpace(src[:index]), []byte("export")), unicode.IsSpace)
	for _, b := range key {
		r := rune(b)
		switch r {
		case '.':
		case '_':
		default:
			if !(unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r)) {
				return "", "", fmt.Errorf("Invalid character %q in key %q", string(r), string(key))
			}
		}
	}

	// TODO: more validation on value needed
	value := bytes.TrimSpace(src[index+1:])

	return string(key), string(value), nil
}
