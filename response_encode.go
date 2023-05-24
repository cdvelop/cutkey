package cutkey

import (
	"fmt"

	"github.com/cdvelop/model"
	json "github.com/fxamacker/cbor/v2"
)

func EncodeResponses(objects []model.Object, requests []model.Response) ([]byte, error) {
	var CutResponses []model.CutResponse

	// Iteramos por cada Packages para generar un CutResponse para cada uno
	for i, data := range requests {

		var object model.Object
		var found_object bool

		for _, obj := range objects {
			if obj.Name == data.Object {
				object = obj
				found_object = true
				break
			}
		}

		if !found_object {
			return nil, fmt.Errorf("objeto %s no encontrado en el slice de objetos", data.Object)
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
			CutOptions: []string{data.Type, data.Object},
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

	// Codificamos el resultado como un array de bytes JSON
	return json.Marshal(CutResponses)
}
