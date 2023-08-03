package cutkey

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/cdvelop/model"
)

func Decode(in io.Reader, o *model.Object) ([]map[string]string, error) {

	if in != nil && o != nil {

		var cut_data []model.CutData
		err := json.NewDecoder(in).Decode(&cut_data)
		if err != nil {
			return nil, fmt.Errorf("error decode json %s", err)
		}

		data, err := o.DataDecode(cut_data...)
		if err != nil {
			return nil, err
		}

		return data, nil

	} else {

		return nil, fmt.Errorf("error objeto nulo al decodificar")
	}

}
