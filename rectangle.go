package plantuml

import "io"

type Rectangle struct {
	name     string
	children []Renderable
}

func NewRectangle(name string) *Rectangle {
	return &Rectangle{name: name}
}

func (p *Rectangle) Add(r ...Renderable) *Rectangle {
	p.children = append(p.children, r...)
	return p
}

func (p *Rectangle) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Print("rectangle \"")
	w.Print(escapeP(p.name))
	w.Print("\" {\n")
	for _, child := range p.children {
		if err := child.Render(wr); err != nil {
			return err
		}
	}
	w.Print("}\n")

	return w.Err
}