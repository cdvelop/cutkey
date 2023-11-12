package cutkey

import (
	"github.com/cdvelop/model"
)

type cut struct {
	model.ObjectsHandler
}

func AddDataConverter(h *model.Handlers) {
	c := &cut{
		ObjectsHandler: h,
	}

	h.DataConverter = c
}
