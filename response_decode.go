package cutkey

import (
	"encoding/json"

	"github.com/cdvelop/model"
)

func (c cut) DecodeResponses(data []byte) (responses []model.Response, err error) {

	var CutResponses []model.CutResponse

	// Decodificamos el array de bytes JSON en un slice de CutResponse
	err = json.Unmarshal(data, &CutResponses)
	if err != nil {
		return nil, model.Error("error json DecodeResponses", err)
	}

	if len(CutResponses) > 0 {

		for i, cr := range CutResponses {

			// fmt.Printf("TAMAÑO CutOptions: %v\n", len(CutResponses[0].CutOptions))
			if len(CutResponses[i].CutOptions) < 2 || len(CutResponses[i].CutOptions) > 4 {
				return nil, model.Error("CutOptions incorrectas en DecodeResponses %s ", CutResponses[i].CutOptions)

			}

			if i >= len(CutResponses) {
				return nil, model.Error("índice", i, "fuera de rango en CutResponses")
			}

			object, err := c.GetObjectByName(cr.CutOptions[1])
			if err != nil {
				return nil, err
			}

			data, err := object.DataDecode(cr.CutData...)
			if err != nil {
				return nil, err
			}

			responses = append(responses, cr.CutResponseDecode(data))
		}
	}

	// fmt.Println("\n=> DATA:", responses)

	return
}
