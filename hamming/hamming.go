package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	var n int = 0
	if len(a) != len(b) {
		return -1, errors.New("error")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				n++
			}
		}
		return n, nil
	}
}
