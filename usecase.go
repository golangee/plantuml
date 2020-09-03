package plantuml

import "io"

type UseCase struct {
	name string
	id   string
}

func NewUseCase(name string) *UseCase {
	return &UseCase{name: name, id: "uc" + nextId()}
}

func (a *UseCase) Self(v **UseCase) *UseCase {
	*v = a
	return a
}

func (a *UseCase) Id() string {
	return a.id
}

func (a *UseCase) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Printf("usecase \"%s\" as %s\n", escapeP(a.name), a.id)

	return w.Err
}
