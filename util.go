package plantuml

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
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

// Render invokes a local plantuml command. fileType can be svg, or png etc.
func RenderLocal(fileType string, renderable Renderable) ([]byte, error) {
	cmd := exec.Command("plantuml", "-t"+fileType, "-p")
	cmd.Env = os.Environ()

	w, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	if _, err := w.Write([]byte(String(renderable))); err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return cmd.Output()
}
