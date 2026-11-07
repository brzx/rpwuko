package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestStartWithCR(t *testing.T) {
	in := strings.NewReader("5 + 10\r")
	var out bytes.Buffer

	Start(in, &out)

	output := out.String()
	if !strings.Contains(output, "{Type:INT Literal:5}") {
		t.Fatalf("expected token INT 5 in output, got: %q", output)
	}
	if !strings.Contains(output, "{Type:+ Literal:+}") {
		t.Fatalf("expected token + in output, got: %q", output)
	}
	if !strings.Contains(output, "{Type:INT Literal:10}") {
		t.Fatalf("expected token INT 10 in output, got: %q", output)
	}
	if !strings.HasSuffix(output, PROMPT) {
		t.Fatalf("expected output to end with prompt %q, got: %q", PROMPT, output)
	}
}

func TestStartWithLF(t *testing.T) {
	in := strings.NewReader("var x = 5;\n")
	var out bytes.Buffer

	Start(in, &out)

	output := out.String()
	if !strings.Contains(output, "{Type:IDENT Literal:var}") {
		t.Fatalf("expected token IDENT var in output, got: %q", output)
	}
	if !strings.Contains(output, "{Type:= Literal:=}") {
		t.Fatalf("expected token = in output, got: %q", output)
	}
}

func TestStartWithCRLF(t *testing.T) {
	in := strings.NewReader("x = 1;\r\n")
	var out bytes.Buffer

	Start(in, &out)

	output := out.String()
	if !strings.Contains(output, "{Type:IDENT Literal:x}") {
		t.Fatalf("expected token IDENT x in output, got: %q", output)
	}
	// Should only print two prompts: initial prompt and post-command prompt
	promptCount := strings.Count(output, PROMPT)
	if promptCount != 2 {
		t.Fatalf("expected exactly 2 prompts for CRLF, got %d. Output: %q", promptCount, output)
	}
}

func TestStartWithBackspace(t *testing.T) {
	// Type 'a', 'b', backspace, 'c', CR -> line should be "ac"
	in := strings.NewReader("ab\bc\r")
	var out bytes.Buffer

	Start(in, &out)

	output := out.String()
	if !strings.Contains(output, "{Type:IDENT Literal:ac}") {
		t.Fatalf("expected token IDENT ac after backspace, got: %q", output)
	}
	if strings.Contains(output, "abc") {
		t.Fatalf("output should not contain abc, got: %q", output)
	}
}
