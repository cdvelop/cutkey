package cutkey_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/cdvelop/cutkey"
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

	data_encode, err := cutkey.EncodeResponses(cutObjects, requests)
	if err != nil {
		log.Fatalf("Error Encoding Packages: %v", err)
	}

	// fmt.Printf("**DATA: %s\n", string(data_encode))

	responses, err := cutkey.DecodeResponses(cutObjects, data_encode)
	if err != nil {
		log.Fatalf("Error Decoding Packages: %v", err)
	}

	if !reflect.DeepEqual(responses, requests) {
		log.Fatalf("Unexpected result:\n\n=>response: %v\n=>expected: %v\n", responses, requests)
	}

}
