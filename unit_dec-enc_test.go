package cutkey_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/cdvelop/model"
)

func TestDecodeEncode(t *testing.T) {

	requests := []model.Response{
		{
			Type:    "read",
			Object:  "user",
			Module:  "Users",
			Message: "ok",
			Data: []map[string]string{
				{"name": "John Doe", "email": "johndoe@example.com", "phone": ""},
				{"name": "Maria Salome", "email": "", "phone": "555"},
			},
		},
		{
			Type:    "create",
			Object:  "product",
			Module:  "Products",
			Message: "ok",
			Data: []map[string]string{
				{"description": "Manzanas", "price": "6000"},
				{"description": "Peras"},
				{"description": "Naranjas", "price": "2000"},
			},
		},
	}

	data_encode := cut.EncodeResponses(requests)

	responses := cut.DecodeResponses(data_encode)

	if !reflect.DeepEqual(responses, requests) {
		log.Fatalf("Unexpected result:\n\n=>response: %v\n=>expected: %v\n", responses, requests)
	}

}
