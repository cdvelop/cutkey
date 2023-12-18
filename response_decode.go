package cutkey

import (
	"strconv"

	"github.com/cdvelop/model"
)

func (c cut) DecodeResponses(data []byte) (responses []model.Response, err string) {
	const this = "DecodeResponses error "

	var resp model.Responses
	// Decodificamos el array de bytes JSON en un slice de Responses
	err = c.DecodeStruct(data, &resp)
	if err != "" {
		return
	}

	responses = append(responses, resp.NoCut...)

	if len(resp.Cut) > 0 {

		for i, cr := range resp.Cut {

			// fmt.Printf("TAMAÑO CutOptions: %v\n", len(resp.Cut[0].CutOptions))
			if len(resp.Cut[i].CutOptions) < 2 || len(resp.Cut[i].CutOptions) > 4 {
				// return nil, "CutOptions incorrectas en DecodeResponses "+ resp.Cut[i].CutOptions
				return nil, this + "CutOptions incorrectas"

			}

			if i >= len(resp.Cut) {
				return nil, this + "índice " + strconv.Itoa(i) + " fuera de rango en resp.Cut"
			}

			object, err := c.MainHandlerGetObjectByName(cr.CutOptions[1])
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
