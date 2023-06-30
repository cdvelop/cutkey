package cutkey

import (
	"encoding/json"

	"github.com/cdvelop/model"

	"fmt"
)

func (c Cut) DecodeResponses(data []byte) (responses []model.Response) {

	var CutResponses []model.CutResponse

	// Decodificamos el array de bytes JSON en un slice de CutResponse
	err := json.Unmarshal(data, &CutResponses)
	if err != nil {
		return c.decodeError("error", err)
	}

	if len(CutResponses) > 0 {

		for i, cr := range CutResponses {

			// fmt.Printf("TAMAÑO CutOptions: %v\n", len(CutResponses[0].CutOptions))
			if len(CutResponses[i].CutOptions) < 2 || len(CutResponses[i].CutOptions) > 4 {
				return c.decodeError("error", fmt.Errorf("CutOptions incorrectas en DecodeResponses %s ", CutResponses[i].CutOptions))

			}

			if i >= len(CutResponses) {
				return c.decodeError("error", fmt.Errorf("índice fuera de rango en CutResponses: %d", i))
			}

			var object *model.Object

			for _, obj := range c.objects {
				if obj.Name == cr.CutOptions[1] {
					object = obj
					break
				}
			}

			if object == nil {

				if cr.CutOptions[1] == "" {
					return c.decodeError("error", fmt.Errorf("objeto no incluido en solicitud"))

				} else if cr.CutOptions[1] != "error" {

					return c.decodeError(cr.CutOptions[1], fmt.Errorf("objeto: %s no encontrado en el slice de objetos", cr.CutOptions[1]))

				} else {

					if len(cr.CutOptions) > 3 && cr.CutOptions[3] != "" { //Message
						// fmt.Println("contiene mensaje")
						return c.decodeError(cr.CutOptions[1], fmt.Errorf(cr.CutOptions[3]))
					} else {
						return c.decodeError(cr.CutOptions[1], fmt.Errorf("error"))
					}

				}

			}

			data, err := object.DataDecode(cr.CutData...)
			if err != nil {
				return c.decodeError(cr.CutOptions[1], err)
			}

			responses = append(responses, cr.CutResponseDecode(data))
		}
	}

	// fmt.Println("\n=> DATA:", responses)

	return
}
