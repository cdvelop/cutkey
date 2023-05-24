package cutkey

import (
	"github.com/cdvelop/model"
	json "github.com/fxamacker/cbor/v2"

	"fmt"
)

func (a Add) DecodeResponses(data []byte) (responses []model.Response, err error) {

	var CutResponses []model.CutResponse
	// Decodificamos el array de bytes JSON en un slice de CutResponse
	err = json.Unmarshal(data, &CutResponses)
	if err != nil {
		return nil, err
	}

	if len(CutResponses) > 0 {

		for i, cr := range CutResponses {
			// fmt.Printf("TAMAÑO CutOptions: %v\n", len(CutResponses[0].CutOptions))
			if len(CutResponses[i].CutOptions) < 2 || len(CutResponses[i].CutOptions) > 4 {
				return nil, fmt.Errorf("CutOptions incorrectas en DecodeResponses %s ", CutResponses[i].CutOptions)
			}

			if i >= len(CutResponses) {
				return nil, fmt.Errorf("índice fuera de rango en CutResponses: %d", i)
			}

			var object model.Object
			var found_object bool

			for _, obj := range *a.Objects {
				if obj.Name == cr.CutOptions[1] {
					object = obj
					found_object = true
					break
				}
			}

			if !found_object {
				return nil, fmt.Errorf("objeto %s no encontrado en el slice de objetos", cr.CutOptions[1])
			}

			data, err := object.DataDecode(cr.CutData...)
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

			responses = append(responses, response)
		}
	}

	// fmt.Println("\n=> DATA:", responses)

	return responses, nil
}
