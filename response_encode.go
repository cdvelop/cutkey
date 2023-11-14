package cutkey

import (
	"encoding/json"

	"github.com/cdvelop/model"
)

func (c cut) EncodeResponses(requests []model.Response) ([]byte, error) {

	var CutResponses []model.CutResponse

	// Iteramos por cada Packages para generar un CutResponse para cada uno
	for _, data := range requests {

		object, err := c.GetObjectByName(data.Object)
		if err != nil {
			return nil, err
		}

		cut_data, err := object.DataEncode(data.Data...)
		if err != nil {
			return nil, err
		}
		// Generamos el CutResponse
		CutResponse := model.CutResponse{
			CutOptions: []string{data.Action, data.Object},
			CutData:    cut_data,
		}

		if data.Message != "" {
			CutResponse.CutOptions = append(CutResponse.CutOptions, data.Message)
		}

		CutResponses = append(CutResponses, CutResponse)

	}

	// fmt.Println("\n=> DATA ENCODE:", CutResponses)
	out, err := json.Marshal(CutResponses)
	if err != nil {
		return nil, model.Error("error json EncodeResponses: %v", err)
	}
	// Codificamos el resultado como un array de bytes JSON
	return out, nil
}
