package plantuml

import (
	"fmt"
	"testing"
)

func TestRenderLocal(t *testing.T) {
	buf, err := RenderLocal("svg", NewDiagram().Add(NewClass("hello").NoteTop(NewNote("world"))))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(buf))
}
