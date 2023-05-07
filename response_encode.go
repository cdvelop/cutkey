package cutkey

import (
	"fmt"

	json "github.com/fxamacker/cbor/v2"

	model "github.com/cdvelop/go_model"
)

func (c Cut) EncodeResponses(requests []model.Response) ([]byte, error) {
	var cutResponses []cutResponse

	// Iteramos por cada Responses para generar un CutResponse para cada uno
	for i, data := range requests {
		// Obtenemos el objeto correspondiente a partir del mapa de objetos
		obj, ok := (*c.models)[data.Object]
		if !ok {
			return nil, fmt.Errorf("objeto %s no encontrado en el mapa de objetos", data.Object)
		}

		// Generamos los Cut_data a partir de la data de la respuesta
		var cut_data []cutData
		for _, m := range data.Data {

			cut_data = append(cut_data, dataEncode(&obj, &m))

		}
		//actualizamos la data original
		requests[i] = data

		// Generamos el CutResponse
		cutResponse := cutResponse{
			CutOptions: []string{data.Type, data.Object},
			CutData:    cut_data,
		}

		if data.Module != "" {
			cutResponse.CutOptions = append(cutResponse.CutOptions, data.Module)
		}

		if data.Message != "" {
			cutResponse.CutOptions = append(cutResponse.CutOptions, data.Message)
		}

		cutResponses = append(cutResponses, cutResponse)

	}

	// fmt.Println("\n=> DATA ENCODE:", cutResponses)

	// Codificamos el resultado como un array de bytes JSON
	return json.Marshal(cutResponses)
}
