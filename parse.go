package godotenvcrypt

import (
	"bytes"
	"errors"
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

func eachStatement(src []byte, callback func([]byte) error) error {
	var size int
	for ; size != -1; src = src[size:] {
		var index int
		index, size = nextStatement(src)
		if index == -1 {
			return nil
		}

		src = src[index:]
		var err error
		if size == -1 {
			err = callback(src)
		} else {
			err = callback(src[:size])
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func environmentPair(src []byte) (string, string, error) {
	index := bytes.IndexRune(src, '=')
	if index == -1 {
		return "", "", errors.New("No valid separator ('=') found")
	}

	key := bytes.TrimSpace(src[:index])
	// export needs to have a space separating afterward to be a valid keyword
	if bytes.HasPrefix(key, []byte("export")) && unicode.IsSpace(rune(key[len([]byte("export"))])) {
		key = bytes.TrimLeftFunc(bytes.TrimPrefix(key, []byte("export")), unicode.IsSpace)
	}
	for _, b := range key {
		r := rune(b)
		switch {
		case unicode.IsLetter(r):
		case unicode.IsDigit(r):
		case unicode.IsSpace(r):
		default:
			switch r {
			case '.':
			case '_':
			default:
				return "", "", fmt.Errorf("Invalid character %q in key %q", string(r), string(key))
			}
		}
	}

	// TODO: more validation on value needed
	value := bytes.TrimLeftFunc(src[index+1:], unicode.IsSpace)
	// trim inline comment
	if index := bytes.LastIndexByte(value, byte('#')); index != -1 {
		value = value[:index]
	}
	value = bytes.TrimRightFunc(value, unicode.IsSpace)

	return string(key), string(value), nil
}
