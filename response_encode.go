package cutkey

import (
	"encoding/json"

	"github.com/cdvelop/model"
)

func (c cut) EncodeResponses(requests ...model.Response) (result []byte, err string) {
	const this = "EncodeResponses error "
	var CutResponses []model.CutResponse

	// Iteramos por cada Packages para generar un CutResponse para cada uno
	for _, data := range requests {

		object, err := c.GetObjectByName(data.Object)
		if err != "" {
			return nil, this + err
		}

		cut_data, err := object.DataEncode(data.Data...)
		if err != "" {
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
	out, e := json.Marshal(CutResponses)
	if e != nil {
		return nil, this + e.Error()
	}
	// Codificamos el resultado como un array de bytes JSON
	return out, ""
}
