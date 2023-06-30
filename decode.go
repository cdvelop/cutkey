package cutkey

import (
	"encoding/json"
	"io"

	"github.com/cdvelop/model"
)

func (c Cut) CutDataDecode(in io.Reader, h *model.Object) ([]map[string]string, error) {

	var cut_data model.CutData
	err := json.NewDecoder(in).Decode(&cut_data)
	if err != nil {
		return nil, err
	}

	data, err := h.DataDecode(cut_data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
