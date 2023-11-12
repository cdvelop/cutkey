package cutkey

import (
	"encoding/json"

	"github.com/cdvelop/model"
)

func jsonEncode(in any) ([]byte, error) {
	out, err := json.Marshal(in)
	if err != nil {
		return nil, model.Error("error json encode", err)
	}
	return out, nil
}

func jsonDecode(in []byte, out any) error {

	err := json.Unmarshal(in, &out)
	if err != nil {
		return model.Error("error json decode", err)
	}

	return nil
}
