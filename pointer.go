package plantuml

import "io"

type Pointer struct {
	from, to string
}

func NewPointer(from, to string) *Pointer {
	return &Pointer{from: from, to: to}
}

func (a *Pointer) Self(v **Pointer) *Pointer {
	*v = a
	return a
}

func (a *Pointer) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Printf("%s --> %s\n", a.from, a.to)

	return w.Err
}
