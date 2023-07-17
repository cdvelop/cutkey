package cutkey

import (
	"encoding/json"

	"github.com/cdvelop/model"
)

func (Cut) encodeError(r *model.Response) (out []byte, err error) {

	out, err = json.Marshal([]model.CutResponse{
		{
			CutOptions: []string{
				"error",
				r.Object,
				r.Module,
				r.Message,
			},
			CutData: []model.CutData{},
		},
	})

	if err != nil {
		return
	}

	return out, nil
}

func (Cut) decodeError(object string, message error) []model.Response {

	return []model.Response{
		{
			Action: "error",
			Object: object,
			// Module:  "error",
			Message: message.Error(),
			// Data:    []map[string]string{},
		},
	}

}
