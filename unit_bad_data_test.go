package cutkey_test

import (
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

	data_decode, err := cut.EncodeResponses(requests)
	if err != nil {
		log.Fatal(err)
	}

	responses := cut.DecodeResponses(data_decode)

	// CASO 0:
	if !reflect.DeepEqual(responses[0], requests[0]) {
		log.Fatalf("Unexpected result:\n\n=>response: %v\n=>expected: %v\n", responses[0], requests[0])
	}

	// CASO 1:
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

	data, err := cut.EncodeResponses(requests)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("|||-%s-|||\n", data)

	resp := cut.DecodeResponses(data)

	// fmt.Printf("resp|||-%s-|||\n", resp)

	if resp[0].Action != "error" {
		log.Fatalln("Se esperaba: error en Action se obtuvo:", resp[0].Action)
	}

	if resp[0].Object != "" {
		t.Fatalf("Se esperaba objeto vació pero obtuvo:%v", resp[0].Object)
	}

	if resp[0].Message != "objeto no incluido en solicitud" {
		log.Fatalln("Se esperaba: objeto no incluido en solicitud se obtuvo:", resp[0].Message)
	}

	// fmt.Printf("result:\n\n=>response: %v\n=>expected: %v\n", resp, requests[0])

}
