package cutkey

import (
	"log"

	"encoding/json"

	"github.com/cdvelop/model"
)

func (c Cut) EncodeResponses(requests []model.Response) []byte {
	var CutResponses []model.CutResponse

	// Iteramos por cada Packages para generar un CutResponse para cada uno
	for i, data := range requests {

		var object *model.Object
		for _, obj := range c.objects {
			if obj.Name == data.Object {
				object = obj
				break
			}
		}

		if object == nil {
			return c.encodeError(&data)
			// return nil, fmt.Errorf("objeto: %s no encontrado en el slice de objetos", data.Object)
		}

		// Generamos los Cut_data a partir de la data de la respuesta
		var cut_data []model.CutData
		for _, m := range data.Data {
			cut_data = append(cut_data, object.DataEncode(m))
		}
		//actualizamos la data original
		requests[i] = data

		// Generamos el CutResponse
		CutResponse := model.CutResponse{
			CutOptions: []string{data.Action, data.Object},
			CutData:    cut_data,
		}

		if data.Module != "" {
			CutResponse.CutOptions = append(CutResponse.CutOptions, data.Module)
		}

		if data.Message != "" {
			CutResponse.CutOptions = append(CutResponse.CutOptions, data.Message)
		}

		CutResponses = append(CutResponses, CutResponse)

	}

	// fmt.Println("\n=> DATA ENCODE:", CutResponses)
	out, err := json.Marshal(CutResponses)
	if err != nil {
		log.Println("Error json EncodeResponses: ", err)
	}
	// Codificamos el resultado como un array de bytes JSON
	return out
}
