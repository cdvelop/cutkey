package cutkey

import (
	"encoding/json"
	"strconv"

	"github.com/cdvelop/model"
)

func (c cut) DecodeResponses(data []byte) (responses []model.Response, err string) {
	const this = "DecodeResponses error "
	var CutResponses []model.CutResponse

	// Decodificamos el array de bytes JSON en un slice de CutResponse
	e := json.Unmarshal(data, &CutResponses)
	if e != nil {
		return nil, this + e.Error()
	}

	if len(CutResponses) > 0 {

		for i, cr := range CutResponses {

			// fmt.Printf("TAMAÑO CutOptions: %v\n", len(CutResponses[0].CutOptions))
			if len(CutResponses[i].CutOptions) < 2 || len(CutResponses[i].CutOptions) > 4 {
				// return nil, "CutOptions incorrectas en DecodeResponses "+ CutResponses[i].CutOptions
				return nil, this + "CutOptions incorrectas"

			}

			if i >= len(CutResponses) {
				return nil, this + "índice " + strconv.Itoa(i) + " fuera de rango en CutResponses"
			}

			object, err := c.GetObjectByName(cr.CutOptions[1])
			if err != "" {
				return nil, this + err
			}

			data, err := object.DataDecode(cr.CutData...)
			if err != "" {
				return nil, this + err
			}

			responses = append(responses, cr.CutResponseDecode(data))
		}
	}

	// fmt.Println("\n=> DATA:", responses)

	return
}
