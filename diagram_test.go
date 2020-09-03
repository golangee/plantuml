package plantuml

import (
	"fmt"
	"testing"
)

func TestNewDiagram(t *testing.T) {
	fmt.Println(String(NewDiagram().
		Include(ThemeCerulean).
		Add(NewPackage("visual model").
			Add(NewClass("Dashboard").
				AddAttrs(
					Attr{Private, false, false, "id", "UUID"},
					Attr{PackagePrivate, true, false, "Stuff", "List<OtherStuff>"},
				),
			),
			NewPackage("domain model").
				Add(NewClass("A"),
					NewClass("B"),
					NewClass("SuperUser"),
					NewAbstractCLass("User").
						AddAttrs(
							Attr{Public, false, true, "Id", "UUID"},
							Attr{Protected, true, true, "Name", "string"},
							Attr{Public, false, false, "MyMethod(a,b string)", "(string, error)"},
						).
						Extends("SuperUser").
						Owns("A").
						Uses("B").
						NoteLeft(NewNote("Left\n<b>Note</b>")).
						NoteTop(NewNote("Top\n<b>Note</b>")).
						NoteRight(NewNote("Right\n<b>Note</b>")).
						NoteBottom(NewNote("Bottom\n<b>Note</b>")),
					NewInterface("Contract").
						AddAttrs(Attr{Public, true, false, "MyMethod()", ""}),
				),
			NewEnum("MetricType", "A", "Other", "Any", "of"),
		),
	))

	var actor *Actor
	var uc1, uc2 *UseCase

	fmt.Println(String(NewDiagram().
		Include(ThemeCerulean).
		Add(
			NewActor("loaner dude").Self(&actor),
			NewRectangle("book search").
				Add(
					NewUseCase("search for tags").Self(&uc1),
					NewUseCase("use autocomplete").Self(&uc2),
				),
			NewPointer(actor.Id(), uc1.Id()),
			NewPointer(actor.Id(), uc2.Id()),
		),
	))
}
