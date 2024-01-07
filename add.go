package cutkey

import (
	"github.com/cdvelop/model"
)

type cut struct {
	model.Logger
	model.ObjectsHandlerAdapter
	object *model.Object
}

func AddDataConverter(h *model.MainHandler) {
	c := &cut{
		Logger:                h,
		ObjectsHandlerAdapter: h,
	}

	h.DataConverter = c
}
