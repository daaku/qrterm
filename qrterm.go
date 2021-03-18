// Package qrterm prints QR codes to the terminal.
package qrterm

import (
	"io"

	"rsc.io/qr"
)

// WriteBig writes a large QR code using color escape codes.
func WriteBig(w io.Writer, code *qr.Code) error {
	const white = "\033[47m  \033[0m"
	const black = "\033[40m  \033[0m"

	for i := 0; i <= code.Size; i++ {
		for j := 0; j <= code.Size; j++ {
			c := white
			if code.Black(j, i) {
				c = black
			}
			if _, err := io.WriteString(w, c); err != nil {
				return err
			}
		}
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}
	return nil
}

// WriteSmall writes a small QR code using UTF8 characters.
func WriteSmall(w io.Writer, code *qr.Code) error {
	const bw = "\u2584"
	const bb = " "
	const wb = "\u2580"
	const ww = "\u2588"

	for i := 0; i <= code.Size; i += 2 {
		for j := 0; j <= code.Size; j++ {
			currB := code.Black(j, i)
			nextB := i+1 < code.Size && code.Black(j, i+1)

			c := wb
			if currB && nextB {
				c = bb
			} else if currB && !nextB {
				c = bw
			} else if !currB && !nextB {
				c = ww
			}
			if _, err := io.WriteString(w, c); err != nil {
				return err
			}
		}
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}
	return nil
}
