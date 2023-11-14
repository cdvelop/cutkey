package cutkey

import (
	"github.com/cdvelop/model"
)

// soportados: []map[string]string,map[string]string
func (c cut) EncodeMaps(map_in any, object_name ...string) (out []byte, err error) {

	var name string
	for _, v := range object_name {
		name = v
	}

	o, err := c.GetObjectByName(name)
	if err != nil {
		return jsonEncode(map_in)
	}

	if len(o.Fields) == 0 { // objeto sin campos solo convertimos a json normal
		return jsonEncode(map_in)
	}

	// desde aca podemos recortar la información
	var cd []model.CutData

	switch data := map_in.(type) {
	case []map[string]string:
		cd, err = o.DataEncode(data...)

	case map[string]string:
		cd, err = o.DataEncode(data)

	default:
		return nil, model.Error("error EncodeMaps tipo de dato no soportado", data)
	}

	if err != nil {
		return nil, err
	}

	return jsonEncode(cd)

}
