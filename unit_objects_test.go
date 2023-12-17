package cutkey_test

import (
	"github.com/cdvelop/model"
)

var cutModule = &model.Module{
	ModuleName: "cutkey.module.test",
	Objects: []*model.Object{
		{
			ObjectName: "user",
			Fields: []model.Field{
				{Name: "name"},
				{Name: "email"},
				{Name: "phone"},
			},
		},
		{
			ObjectName: "product",
			Fields: []model.Field{
				{Name: "description"},
				{Name: "price"},
			},
		},
	},
}
