package cutkey

import (
	"github.com/cdvelop/model"
)

func (c cut) EncodeResponses(requests ...model.Response) (result []byte, err string) {
	const this = "EncodeResponses "

	responses := model.Responses{
		NoCut: []model.Response{},
		Cut:   []model.CutResponse{},
	}

	// Iteramos por cada Packages para generar un CutResponse para cada uno
	for _, data := range requests {

		object, err := c.GetObjectBY(data.Object, "")

		// fmt.Println("data object:", data.Object, " OBJETO OBTENIDO:", object)

		if err != "" {

			if data.Action == "" && data.Object == "" {
				return nil, this + "error action y objeto no declarado"
			}

			responses.NoCut = append(responses.NoCut, data)
			continue
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

		responses.Cut = append(responses.Cut, CutResponse)

	}

	result, err = c.EncodeStruct(responses)
	if err != "" {
		return nil, this + err
	}

	// fmt.Println("\n=> DATA ENCODE:", CutResponses)
	// Codificamos el resultado como un array de bytes JSON
	return
}
