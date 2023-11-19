package cutkey_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func TestDecodeEncode(t *testing.T) {
	handler := model.Handlers{}
	handler.AddObjects(cutObjects...)
	cutkey.AddDataConverter(&handler)

	requests := []model.Response{
		{
			Action:  "read",
			Object:  "user",
			Message: "ok",
			Data: []map[string]string{
				{"name": "John Doe", "email": "johndoe@example.com", "phone": ""},
				{"name": "Maria Salome", "email": "", "phone": "555"},
			},
		},
		{
			Action:  "create",
			Object:  "product",
			Message: "ok",
			Data: []map[string]string{
				{"description": "Manzanas", "price": "6000"},
				{"description": "Peras"},
				{"description": "Naranjas", "price": "2000"},
			},
		},
	}

	data_encode, err := handler.EncodeResponses(requests...)
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Printf("|||-%s-|||\n", data_encode)

	responses, err := handler.DecodeResponses(data_encode)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(responses, requests) {
		log.Fatalf("Unexpected result:\n\n=>response:\n%v\n\n=>expected:\n%v\n", responses, requests)
	}

}
