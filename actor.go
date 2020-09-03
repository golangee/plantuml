package plantuml

import "io"

type Actor struct {
	name     string
	id       string
}

func NewActor(name string) *Actor {
	return &Actor{name: name, id: "ac" + nextId()}
}

func (a *Actor) Self(v **Actor) *Actor {
	*v = a
	return a
}

func (a *Actor) Id() string {
	return a.id
}

func (a *Actor) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Printf("actor \"%s\" as %s\n", escapeP(a.name), a.id)

	return w.Err
}

