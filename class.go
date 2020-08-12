package plantuml

import (
	"io"
	"sort"
)

type Visibility string

const (
	Private        Visibility = "-"
	Protected      Visibility = "#"
	PackagePrivate Visibility = "~"
	Public         Visibility = "+"
)

type Attr struct {
	Visibility Visibility
	Abstract   bool
	Static     bool
	Name       string
	Type       string
}

type Class struct {
	typeName            string
	name                string
	attrs               []Attr
	notes               map[string][]*Note
	extends, uses, owns []string
}

func NewClass(name string) *Class {
	return &Class{typeName: "class", name: name, notes: map[string][]*Note{}}
}

func NewInterface(name string) *Class {
	return &Class{typeName: "interface", name: name, notes: map[string][]*Note{}}
}

func NewAbstractCLass(name string) *Class {
	return &Class{typeName: "abstract class", name: name, notes: map[string][]*Note{}}
}

func (d *Class) Name() string {
	return d.name
}

func (d *Class) AddAttrs(attrs ...Attr) *Class {
	d.attrs = append(d.attrs, attrs...)
	return d
}

func (d *Class) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Print(d.typeName)
	w.Print(" \"")
	w.Print(escapeP(d.name))
	w.Print("\" {\n")
	for _, attr := range d.attrs {
		if attr.Abstract {
			w.Print("{abstract} ")
		}

		if attr.Static {
			w.Print("{static} ")
		}

		w.Print(string(attr.Visibility))
		w.Print(attr.Name)
		w.Print(" : ")
		w.Print(attr.Type)
		w.Print("\n")

	}

	w.Print("}\n")

	var keys []string
	for key, _ := range d.notes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		for _, note := range d.notes[key] {
			if err := note.renderDirOf(key, escapeP(d.name), wr); err != nil {
				return err
			}
		}
	}

	for _, super := range d.extends {
		w.Print(escapeP(super))
		w.Print(" <|-- ")
		w.Print(escapeP(d.name))
		w.Print("\n")
	}

	for _, child := range d.owns {
		w.Print(escapeP(child))
		w.Print(" *-- ")
		w.Print(escapeP(d.name))
		w.Print("\n")
	}

	for _, child := range d.uses {
		w.Print(escapeP(child))
		w.Print(" o-- ")
		w.Print(escapeP(d.name))
		w.Print("\n")
	}
	return w.Err
}

func (d *Class) NoteLeft(n *Note) *Class {
	r := d.notes["left"]
	r = append(r, n)
	d.notes["left"] = r
	return d
}

func (d *Class) NoteRight(n *Note) *Class {
	r := d.notes["right"]
	r = append(r, n)
	d.notes["right"] = r
	return d
}

func (d *Class) NoteBottom(n *Note) *Class {
	r := d.notes["bottom"]
	r = append(r, n)
	d.notes["bottom"] = r
	return d
}

func (d *Class) NoteTop(n *Note) *Class {
	r := d.notes["top"]
	r = append(r, n)
	d.notes["top"] = r
	return d
}

func (d *Class) Extends(names ...string) *Class {
	d.extends = append(d.extends, names...)
	return d
}

// Owns is a composition, the given types cannot exist without this one.
func (d *Class) Owns(names ...string) *Class {
	d.owns = append(d.owns, names...)
	return d
}

// Uses is an aggregation, the given types do exist anyway.
func (d *Class) Uses(names ...string) *Class {
	d.uses = append(d.uses, names...)
	return d
}
