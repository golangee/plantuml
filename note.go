package plantuml

import (
	"io"
	"strconv"
	"sync"
)

var noteNum int
var mutex sync.Mutex

type Note struct {
	id   string
	text string
}

func NewNote(text string) *Note {
	mutex.Lock()
	defer mutex.Unlock()
	noteNum++

	return &Note{text: text, id: "N" + strconv.Itoa(noteNum)}
}

func (p *Note) Render(wr io.Writer) error {
	return p.renderUnconnected(wr)
}

func (p *Note) renderDirOf(pos, name string, wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Print("note ")
	w.Print(pos)
	w.Print(" of ")
	w.Print(name)
	w.Print("\n")
	w.Print(p.text)
	w.Print("\n")
	w.Print("end note\n")

	return w.Err
}

func (p *Note) renderUnconnected(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Print("note as ")
	w.Print(p.id)
	w.Print("\n")
	w.Print(p.text)
	w.Print("end note\n")

	return w.Err
}
