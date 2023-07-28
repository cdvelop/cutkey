package cutkey

import (
	"encoding/json"

	"github.com/cdvelop/model"
)

func Encode(o *model.Object, data ...map[string]string) (out []byte, err error) {

	if o != nil && data != nil {

		cd := o.DataEncode(data...)

		out, err = json.Marshal(cd)
		if err != nil {
			return nil, err
		}
	}

	return out, nil

}
