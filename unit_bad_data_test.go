package cutkey_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

var cut = cutkey.Add(cutObjects...)

func TestDecodeEncodeBadData(t *testing.T) {

	requests := []model.Response{

		{ // CASO 0: sin module y sin mensaje se espera que se iguale el nombre del module al del objeto
			Action: "read",
			Object: "user",
			Data: []map[string]string{
				{"name": "bad"},
			},
		},
		{ // CASO 1: sin data
			Action: "delete",
			Object: "product",
		},
	}

	data_decode := cut.EncodeResponses(requests)

	responses := cut.DecodeResponses(data_decode)

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

	requests := []model.Response{

		{ // CASO 0: sin nada se espera error
			Action: "",
			Object: "",
		},
	}

	data := cut.EncodeResponses(requests)

	resp := cut.DecodeResponses(data)

	if resp[0].Action != "error" {
		log.Fatalln("Se esperaba: error en Action se obtuvo:", resp[0].Action)
	}

	if resp[0].Message != "objeto no incluido en solicitud" {
		log.Fatalln("Se esperaba: objeto no incluido en solicitud se obtuvo:", resp[0].Message)
	}

	fmt.Printf("result:\n\n=>response: %v\n=>expected: %v\n", resp, requests[0])

}
