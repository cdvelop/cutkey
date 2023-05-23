package cutkey

import (
	"log"
	"reflect"
	"testing"

	"github.com/cdvelop/model"
)

func TestDecodeEncodeBadData(t *testing.T) {
	cut := Add(&objects)

	requests := []*model.Response{

		{ // CASO 0: sin module y sin mensaje se espera que se iguale el nombre del module al del objeto
			Type:   "read",
			Object: "user",
			Data: []map[string]string{
				{"name": "bad"},
			},
		},
		{ // CASO 1: sin data
			Type:   "delete",
			Object: "product",
		},
	}

	data_decode, err := cut.EncodeResponses(requests)
	if err != nil {
		log.Fatalf("Error Encoding Packages: %v", err)
	}
	// fmt.Printf("%x\n", data_decode)

	responses, err := cut.DecodeResponses(data_decode)
	if err != nil {
		log.Fatalf("Error Decoding Packages: %v", err)
	}

	// CASO 0: agregamos al original el module para comparar
	requests[0].Module = "user"
	if !reflect.DeepEqual(responses[0], requests[0]) {
		log.Fatalf("Unexpected result:\n\n=>response: %v\n=>expected: %v\n", responses[0], requests[0])
	}

	// CASO 1: agregamos al original el module para comparar
	requests[1].Module = "product"
	if !reflect.DeepEqual(responses[1], requests[1]) {
		log.Fatalf("Unexpected result:\n\n=>response: %v\n=>expected: %v\n", responses[1], requests[1])
	}

	// fmt.Printf("result:\n\n=>response: %v\n=>expected: %v\n", responses[1], requests[1])

}

func TestDecodeEncodeBadNoData(t *testing.T) {
	cut := Add(&objects)

	requests := []*model.Response{

		{ // CASO 0: sin nada se espera error
			Type:   "",
			Object: "",
		},
	}

	_, err := cut.EncodeResponses(requests)
	if err == nil {
		log.Fatalf("Error Encoding Packages: %v", err)
	}

}
