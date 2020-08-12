package plantuml

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type strWriter struct {
	Writer io.Writer
	Err    error
}

func (w strWriter) Print(str string) {
	if w.Err != nil {
		return
	}

	_, err := w.Writer.Write([]byte(str))
	if err != nil {
		w.Err = err
	}
}

func (w strWriter) Printf(format string, args ...interface{}) {
	if w.Err != nil {
		return
	}

	_, err := w.Writer.Write([]byte(fmt.Sprintf(format, args...)))
	if err != nil {
		w.Err = err
	}
}

func escapeP(str string) string {
	return strings.ReplaceAll(str, `"`, "<U+0022>")
}

type Renderable interface {
	Render(wr io.Writer) error
}

func String(r Renderable) string {
	sb := &bytes.Buffer{}
	if err := r.Render(sb); err != nil {
		return err.Error()
	}

	return sb.String()
}
