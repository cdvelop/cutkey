package cutkey

import (
	"github.com/cdvelop/model"
)

func (c cut) DecodeMaps(in []byte, object_name ...string) ([]map[string]string, error) {

	var name string
	for _, v := range object_name {
		name = v
	}

	o, err := c.GetObjectByName(name)
	if err != nil {
		return c.decodeMaps(in)
	}

	if len(o.Fields) == 0 { // objeto sin campos salida normal
		return c.decodeMaps(in)
	}

	var cut_data []model.CutData
	err = jsonDecode(in, &cut_data)
	if err != nil {
		return nil, err
	}

	data, err := o.DataDecode(cut_data...)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (c cut) decodeMaps(in []byte) ([]map[string]string, error) {

	const message = "tipo de dato no soportado:"

	var data interface{}

	err := jsonDecode(in, &data)
	if err != nil {
		return nil, err
	}

	switch items := data.(type) {
	case []interface{}:
		result := make([]map[string]string, len(items))
		for i, item := range items {
			if itemData, ok := item.(map[string]interface{}); ok {
				stringMap := make(map[string]string)
				for key, value := range itemData {
					if str, ok := value.(string); ok {
						stringMap[key] = str
					}
				}
				result[i] = stringMap
			} else {
				return nil, model.Error(message, item)
			}
		}
		return result, nil
	case map[string]interface{}:
		return []map[string]string{convertMap(items)}, nil
	default:
		c.Log(message, "decodeMaps data (%T): %v", items, items)
		return nil, model.Error(message, data)
	}
}

func convertMap(input map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for key, value := range input {
		if str, ok := value.(string); ok {
			result[key] = str
		}
	}
	return result
}
