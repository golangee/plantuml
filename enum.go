package plantuml

import (
	"io"
)

type Enum struct {
	name     string
	children []string
}

func NewEnum(name string, values ...string) *Enum {
	return &Enum{name: name, children: values}
}

func (p *Enum) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Print("enum \"")
	w.Print(escapeP(p.name))
	w.Print("\" {\n")
	for _, child := range p.children {
		w.Print(child)
		w.Print("\n")
	}
	w.Print("}\n")

	return w.Err
}
