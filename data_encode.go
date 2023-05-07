package cutkey

import (
	model "github.com/cdvelop/go_model"
)

// dataEncode quita los nombres de los campos de la data según modelo del objeto
func dataEncode(o *model.Object, data *map[string]string) cutData {

	cut_data := cutData{
		Index: make(map[uint8]uint8),
		Data:  []string{},
	}

	for field_index, f := range o.Fields {
		if value, exist := (*data)[f.Name]; exist && value != "" {

			cut_data.Data = append(cut_data.Data, value)
			// indice del campo + posición en el array de data
			cut_data.Index[uint8(field_index)] = uint8(len(cut_data.Data) - 1)
		} else {
			// eliminamos el campo del objeto original
			delete(*data, f.Name)
		}
	}

	return cut_data
}
