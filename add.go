package cutkey

import (
	"github.com/cdvelop/model"
)

type cut struct {
	model.Logger
	model.ObjectHandler
}

func AddDataConverter(h *model.MainHandler) {
	c := &cut{
		Logger:        h,
		ObjectHandler: h,
	}

	h.DataConverter = c
}
