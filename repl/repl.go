package repl

import (
	"fmt"
	"io"

	"github.com/brzx/rpwuko/lexer"
	"github.com/brzx/rpwuko/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	var buf [1]byte
	var lineBuf [256]byte
	lineLen := 0
	lastWasCR := false

	fmt.Fprintf(out, PROMPT)

	for {
		n, err := in.Read(buf[:])
		if err == io.EOF {
			return
		}
		if err != nil || n == 0 {
			continue
		}

		c := buf[0]

		switch c {
		case '\r':
			lastWasCR = true
			fmt.Fprintf(out, "\r\n")
			if lineLen > 0 {
				line := string(lineBuf[:lineLen])
				l := lexer.New(line)

				for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
					fmt.Fprintf(out, "%+v\r\n", tok)
				}
				lineLen = 0
			}
			fmt.Fprintf(out, PROMPT)

		case '\n':
			if lastWasCR {
				// Skip LF immediately following CR (\r\n pair)
				lastWasCR = false
				continue
			}
			lastWasCR = false
			fmt.Fprintf(out, "\r\n")
			if lineLen > 0 {
				line := string(lineBuf[:lineLen])
				l := lexer.New(line)

				for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
					fmt.Fprintf(out, "%+v\r\n", tok)
				}
				lineLen = 0
			}
			fmt.Fprintf(out, PROMPT)

		case '\b', 0x7F:
			lastWasCR = false
			if lineLen > 0 {
				lineLen--
				out.Write([]byte{'\b', ' ', '\b'})
			}

		default:
			lastWasCR = false
			if c >= 32 && c < 127 {
				if lineLen < len(lineBuf)-1 {
					lineBuf[lineLen] = c
					lineLen++
					out.Write([]byte{c})
				}
			}
		}
	}
}

