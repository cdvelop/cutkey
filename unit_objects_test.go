package cutkey_test

import (
	"github.com/cdvelop/model"
)

var cutObjects = []*model.Object{
	{
		ApiHandler: model.ApiHandler{
			Name: "user",
		},
		Fields: []model.Field{
			{Name: "name"},
			{Name: "email"},
			{Name: "phone"},
		},
	},
	{
		ApiHandler: model.ApiHandler{
			Name: "product",
		},
		Fields: []model.Field{
			{Name: "description"},
			{Name: "price"},
		},
	},
}
