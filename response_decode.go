package cutkey

import (
	json "github.com/fxamacker/cbor/v2"

	"fmt"

	"github.com/cdvelop/model"
)

func (c Cut) DecodeResponses(data []byte) (responses []*model.Response, err error) {

	var cutResponses []cutResponse
	// Decodificamos el array de bytes JSON en un slice de CutResponse
	err = json.Unmarshal(data, &cutResponses)
	if err != nil {
		return nil, err
	}

	if len(cutResponses) > 0 {

		for i, cr := range cutResponses {
			// fmt.Printf("TAMAÑO CutOptions: %v\n", len(cutResponses[0].CutOptions))
			if len(cutResponses[i].CutOptions) < 2 || len(cutResponses[i].CutOptions) > 4 {
				return nil, fmt.Errorf("CutOptions incorrectas en DecodeResponses %s ", cutResponses[i].CutOptions)
			}

			if i >= len(cutResponses) {
				return nil, fmt.Errorf("índice fuera de rango en cutResponses: %d", i)
			}

			// Obtenemos el objeto correspondiente a partir del mapa de objetos
			obj, ok := (*c.models)[cr.CutOptions[1]]
			if !ok {
				return nil, fmt.Errorf("objeto %s no encontrado en el mapa de objetos", cr.CutOptions[1])
			}

			data, err := dataDecode(&obj, cr.CutData...)
			if err != nil {
				return nil, err
			}

			var response = model.Response{
				Module:  "",
				Message: "",
				Data:    data,
			}

			response.Type = cr.CutOptions[0]
			response.Object = cr.CutOptions[1]

			if len(cr.CutOptions) > 2 {
				// fmt.Println("si contiene module")
				response.Module = cr.CutOptions[2]
			} else {
				// fmt.Println("no contiene module copiamos el objeto")
				response.Module = response.Object
			}

			if len(cr.CutOptions) > 3 && cr.CutOptions[3] != "" {
				// fmt.Println("contiene mensaje")
				response.Message = cr.CutOptions[3]
			}

			responses = append(responses, &response)
		}
	}

	// fmt.Println("\n=> DATA:", responses)

	return responses, nil
}
