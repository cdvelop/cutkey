package cutkey

import (
	"encoding/json"
)

func jsonEncode(in any) (result []byte, err string) {
	out, e := json.Marshal(in)
	if e != nil {
		return nil, "error json encode" + e.Error()
	}
	return out, ""
}

func jsonDecode(in []byte, out any) (err string) {

	e := json.Unmarshal(in, &out)
	if e != nil {
		return "error json decode " + e.Error()
	}

	return ""
}
