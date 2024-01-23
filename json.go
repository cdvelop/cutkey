package cutkey

import (
	"encoding/json"
)

func (cut) EncodeStruct(in any) (result []byte, err string) {
	out, e := json.Marshal(in)
	if e != nil {
		return nil, "EncodeStruct error " + e.Error()
	}
	return out, ""
}

func (cut) DecodeStruct(in []byte, out any) (err string) {

	e := json.Unmarshal(in, out)
	if e != nil {
		return "DecodeStruct error " + e.Error()
	}

	return ""
}
