package cutkey_test

import (
	"github.com/cdvelop/model"
)

var cutObjects = []model.Object{
	{
		Name: "user",
		Fields: []model.Field{
			{Name: "name"},
			{Name: "email"},
			{Name: "phone"},
		},
	},
	{
		Name: "product",
		Fields: []model.Field{
			{Name: "description"},
			{Name: "price"},
		},
	},
}
