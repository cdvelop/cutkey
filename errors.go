package cutkey

import (
	"log"

	"encoding/json"

	"github.com/cdvelop/model"
)

func (Cut) encodeError(r *model.Response) []byte {

	out, err := json.Marshal([]model.CutResponse{
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
		log.Println("cutkey encodeError: ", err)
	}

	return out
}

func (Cut) decodeError(object string, message error) []model.Response {

	return []model.Response{
		{
			Action:  "error",
			Object:  object,
			Module:  "error",
			Message: message.Error(),
			Data:    []map[string]string{},
		},
	}

}
