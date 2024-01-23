package cutkey

import (
	"github.com/cdvelop/model"
)

// soportados: []map[string]string,map[string]string
func (c *cut) EncodeMaps(map_in any, object_name ...string) (out []byte, err string) {

	var name string
	for _, v := range object_name {
		name = v
	}

	c.object, err = c.GetObjectBY(name, "")
	if err != "" {
		return c.EncodeStruct(map_in)
	}

	if len(c.object.Fields) == 0 { // objeto sin campos solo convertimos a json normal
		return c.EncodeStruct(map_in)
	}

	// fmt.Println("CAMPOS OBJETO FILE:", o.Name, o.Fields)

	// desde aca podemos recortar la información
	var cd []model.CutData

	switch data := map_in.(type) {
	case []map[string]string:
		cd, err = c.object.DataEncode(data...)

	case map[string]string:
		cd, err = c.object.DataEncode(data)

	default:
		const msg = "EncodeMaps error tipo de dato no soportado"
		c.Log(msg, data)
		return nil, msg
	}

	if err != "" {
		return nil, err
	}

	return c.EncodeStruct(cd)

}
