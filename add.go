package cutkey

import (
	"github.com/cdvelop/model"
)

type cut struct {
	model.Logger
	model.ObjectsHandler
}

func AddDataConverter(h *model.Handlers) {
	c := &cut{
		Logger:         h,
		ObjectsHandler: h,
	}

	h.DataConverter = c
}
