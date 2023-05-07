package cutkey

import (
	"fmt"

	model "github.com/cdvelop/go_model"
)

func dataDecode(o *model.Object, cut_data ...CutData) ([]map[string]string, error) {
	var decodedData []map[string]string

	for _, c := range cut_data {

		if len(c.Index) != len(c.Data) {
			return nil, fmt.Errorf("tamaño index %d en dataDecode es diferente a la data enviada: %d", len(c.Index), len(c.Data))
		}

		if len(c.Index) > len(o.Fields) {
			return nil, fmt.Errorf("tamaño index: %d en dataDecode es mayor al del modelo del objeto %v: %d", len(c.Index), o.Name, len(c.Data))
		}

		data_decoded := make(map[string]string)

		for field_index, field := range o.Fields {

			if value_position, exist := c.Index[uint8(field_index)]; exist {

				value := c.Data[value_position]

				data_decoded[field.Name] = value

			}
		}

		decodedData = append(decodedData, data_decoded)
	}

	return decodedData, nil
}
