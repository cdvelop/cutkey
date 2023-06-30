package cutkey

import (
	"github.com/cdvelop/model"
)

type Cut struct {
	objects []*model.Object
}

func Add(objects ...*model.Object) *Cut {
	c := Cut{}
	c.objects = append(c.objects, objects...)
	return &c
}
