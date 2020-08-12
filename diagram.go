package plantuml

import "io"

const ThemeCerulean = "https://raw.githubusercontent.com/bschwarz/puml-themes/master/themes/cerulean/puml-theme-cerulean.puml"

type Diagram struct {
	includes    []string
	renderables []Renderable
}

func NewDiagram() *Diagram {
	d := &Diagram{}
	return d
}

func (d *Diagram) Add(r ...Renderable) *Diagram {
	d.renderables = append(d.renderables, r...)
	return d
}

func (d *Diagram) Include(inc ...string) *Diagram {
	d.includes = append(d.includes, inc...)
	return d
}

func (d *Diagram) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Print("@startuml\n")
	for _, include := range d.includes {
		w.Print("!include ")
		w.Print(include)
		w.Print("\n")
	}

	for _, renderable := range d.renderables {
		if err := renderable.Render(wr); err != nil {
			return err
		}
	}

	w.Print("@enduml\n")
	return w.Err
}
