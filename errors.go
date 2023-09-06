package cutkey

import (
	"encoding/json"

	"github.com/cdvelop/model"
)

func (Cut) encodeError(r *model.Response, error_message string) (out []byte, err error) {

	var space string
	if error_message != "" {
		space = " "
	}

	out, err = json.Marshal([]model.CutResponse{
		{
			CutOptions: []string{
				"error",
				r.Object,
				error_message + space + r.Message,
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
			Action:  "error",
			Object:  object,
			Message: message.Error(),
		},
	}

}
