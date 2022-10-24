package accessor

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

type Extension struct {
	entc.DefaultExtension
}

func (*Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("greet").ParseFiles("template/greet.tmpl")),
	}
}
