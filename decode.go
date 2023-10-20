package cutkey

import (
	"encoding/json"
	"io"

	"github.com/cdvelop/model"
)

func Decode(in io.Reader, o *model.Object) ([]map[string]string, error) {

	if in != nil && o != nil {

		var cut_data []model.CutData
		err := json.NewDecoder(in).Decode(&cut_data)
		if err != nil {
			return nil, model.Error("error decode json %s", err)
		}

		data, err := o.DataDecode(cut_data...)
		if err != nil {
			return nil, err
		}

		return data, nil

	} else {

		return nil, model.Error("error objeto nulo al decodificar")
	}

}
