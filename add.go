package cutkey

import (
	"github.com/cdvelop/model"
)

type cut struct {
	objects []*model.Object
}

func Add(objects ...*model.Object) *cut {
	c := cut{}
	for _, o := range objects {
		if o != nil {
			c.objects = append(c.objects, o)
		}
	}

	return &c
}
