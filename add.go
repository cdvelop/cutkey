package cutkey

import (
	"github.com/cdvelop/model"
)

type Cut struct {
	objects []model.Object
}

func Add(objects ...model.Object) *Cut {
	c := Cut{}
	for _, o := range objects {
		c.objects = append(c.objects, o)
	}
	return &c
}
