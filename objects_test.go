package cutkey

import (
	model "github.com/cdvelop/go_model"
)

var objects = map[string]model.Object{
	"user": {
		Name: "Usuario",
		Fields: []model.Field{
			{Name: "name"},
			{Name: "email"},
			{Name: "phone"},
		},
	},
	"product": {
		Name: "Producto",
		Fields: []model.Field{
			{Name: "description"},
			{Name: "price"},
		},
	},
}
