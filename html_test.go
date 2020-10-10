package ctgov

import (
	"bytes"
	"testing"
)

func TestWriteOpenTag(t *testing.T) {
	tests := []struct {
		el       string
		testOpen string
	}{
		{"ol", "<ol>"},
		{"ul", "<ul>"},
		{"li", "<li>"},
	}
	for _, test := range tests {
		var buffer bytes.Buffer
		WriteOpenTag(test.el, &buffer)
		if buffer.String() != test.testOpen {
			t.Errorf("Bad writing to buffer, got: %s, should be: %s.", buffer.String(), test.testOpen)
		}
	}
}

func TestWriteCloseTag(t *testing.T) {
	tests := []struct {
		el        string
		testClose string
	}{
		{"ol", "</ol>"},
		{"ul", "</ul>"},
		{"li", "</li>"},
	}
	for _, test := range tests {
		var buffer bytes.Buffer
		WriteCloseTag(test.el, &buffer)
		if buffer.String() != test.testClose {
			t.Errorf("Bad writing to buffer, got: %s, should be: %s.", buffer.String(), test.testClose)
		}
	}
}
