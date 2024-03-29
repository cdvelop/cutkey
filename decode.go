package cutkey

import (
	"github.com/cdvelop/model"
)

func (c *cut) DecodeMaps(in []byte, object_name ...string) (data []map[string]string, err string) {

	var name string
	for _, v := range object_name {
		name = v
	}

	c.object, err = c.GetObjectBY(name, "")
	if err != "" {
		return c.decodeMaps(in)
	}

	if len(c.object.Fields) == 0 { // objeto sin campos salida normal
		return c.decodeMaps(in)
	}

	var cut_data []model.CutData
	err = c.DecodeStruct(in, &cut_data)
	if err != "" {
		return nil, "DecodeMaps error " + err
	}

	return c.object.DataDecode(cut_data...)

}

func (c cut) decodeMaps(in []byte) (result []map[string]string, err string) {
	const e = "decodeMaps "
	var message = e + "tipo de dato no soportado:"

	if len(in) == 0 {
		return []map[string]string{}, ""
	}

	// fmt.Println("DATA DE ENTRADA", in)

	var data interface{}

	err = c.DecodeStruct(in, &data)
	if err != "" {
		return nil, e + err
	}

	// fmt.Println("DATA DE SALIDA ANY", data)

	switch items := data.(type) {
	case []interface{}:
		result = make([]map[string]string, len(items))
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
				// fmt.Printf(message+" %t",item)
				c.Log(e+message, "data (%T): %v", items, items)
				return nil, message
			}
		}
		return result, ""
	case map[string]interface{}:
		return []map[string]string{convertMap(items)}, ""

	case map[string]string:
		// c.Log(e + "EN UN MAPA STRING")
		return []map[string]string{items}, ""

	case nil:
		return []map[string]string{}, ""

	default:
		// fmt.Printf("data (%T): %v\n", items, items)
		c.Log(e+message, "data (%T): %v", items, items)
		return nil, message
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
