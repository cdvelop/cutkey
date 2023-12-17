package cutkey

import (
	"github.com/cdvelop/model"
)

type cut struct {
	model.Logger
	model.ObjectsHandler
}

func AddDataConverter(h *model.MainHandler) {
	c := &cut{
		Logger:         h,
		ObjectsHandler: h,
	}

	h.DataConverter = c
}
