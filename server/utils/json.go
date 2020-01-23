package utils

import (
	"bufio"
	"fmt"
)

//ReadJSON reads JSON from r and return []byte
func ReadJSON(r *bufio.Reader) (data []byte, err error) {
	data = []byte{}
	balance := 0
	inString := false
	passByte := false
	b, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	if b != byte('{') {
		return nil, fmt.Errorf("JSON must start from '{', not from '%v'", rune(b))
	}

	data = append(data, b)
	balance++
	for balance > 0 {
		b, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		data = append(data, b)

		if passByte {
			passByte = false
		} else {
			switch b {
			case byte('\\'):
				passByte = true
			case byte('"'):
				inString = !inString
			}
			if !inString {
				switch b {
				case byte('{'):
					balance++
				case byte('}'):
					balance--
				}
			}
		}
	}
	return data, nil
}
