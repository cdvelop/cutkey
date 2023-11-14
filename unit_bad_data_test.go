package cutkey_test

import (
	"reflect"
	"testing"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func TestDecodeEncodeBadData(t *testing.T) {
	handler := model.Handlers{}
	handler.AddObjects(cutObjects...)
	cutkey.AddDataConverter(&handler)

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

	data_decode, err := handler.EncodeResponses(requests)
	if err != nil {
		t.Fatal(err)
	}

	responses, err := handler.DecodeResponses(data_decode)
	if err != nil {
		t.Fatal(err)
	}

	// CASO 0:
	if !reflect.DeepEqual(responses[0], requests[0]) {
		t.Fatalf("Unexpected result:\n\n=>response: %v\n=>expected: %v\n", responses[0], requests[0])
	}

	// CASO 1:
	if !reflect.DeepEqual(responses[1], requests[1]) {
		t.Fatalf("Unexpected result:\n\n=>response: %v\n=>expected: %v\n", responses[1], requests[1])
	}

	// fmt.Printf("result:\n\n=>response: %v\n=>expected: %v\n", responses[1], requests[1])

}

func TestDecodeEncodeBadNoData(t *testing.T) {

	handler := model.Handlers{}
	// handler.AddObjects(cutObjects...)
	cutkey.AddDataConverter(&handler)

	requests := []model.Response{

		{ // CASO 0: sin nada se espera error
			Action: "",
			Object: "",
		},
	}

	data, err := handler.EncodeResponses(requests)
	if err == nil {
		t.Fatal("se esperaba error EncodeResponses y no se obtuvo", err)
	}

	// fmt.Printf("|||-%s-|||\n", data)

	resp, err := handler.DecodeResponses(data)
	if err == nil {
		t.Fatal("se esperaba error DecodeResponses y no se obtuvo", err, resp)
	}

}
