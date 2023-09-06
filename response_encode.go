package cutkey

import (
	"fmt"

	"encoding/json"

	"github.com/cdvelop/model"
)

func (c Cut) EncodeResponses(requests []model.Response) ([]byte, error) {
	var CutResponses []model.CutResponse

	// Iteramos por cada Packages para generar un CutResponse para cada uno
	for _, data := range requests {

		var object *model.Object
		for _, obj := range c.objects {
			if obj.Name == data.Object {
				object = obj
				break
			}
		}

		if object == nil {
			return c.encodeError(&data, "objeto "+data.Object+" no registrado en cutkey")
		}

		// Generamos el CutResponse
		CutResponse := model.CutResponse{
			CutOptions: []string{data.Action, data.Object},
			CutData:    object.DataEncode(data.Data...),
		}

		if data.Message != "" {
			CutResponse.CutOptions = append(CutResponse.CutOptions, data.Message)
		}

		CutResponses = append(CutResponses, CutResponse)

	}

	// fmt.Println("\n=> DATA ENCODE:", CutResponses)
	out, err := json.Marshal(CutResponses)
	if err != nil {
		return nil, fmt.Errorf("error json EncodeResponses: %v", err)
	}
	// Codificamos el resultado como un array de bytes JSON
	return out, nil
}
