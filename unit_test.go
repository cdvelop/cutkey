package cutkey_test

import (
	"reflect"
	"testing"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func TestWhitOutObject(t *testing.T) {
	handler := model.Handlers{}
	cutkey.AddDataConverter(&handler)

	expected := []map[string]string{
		{"description": "Manzanas", "price": "6000"},
		{"description": "Peras"},
	}

	out, err := cut.EncodeMaps(expected)
	if err != nil {
		t.Fatal("no se esperaba error en cut.EncodeMaps y se obtuvo", err, out)
	}

	result, err := cut.DecodeMaps(out)
	if err != nil {
		t.Fatal("no se esperaba error en cut.DecodeMaps y se obtuvo", err, result)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result:\n\n=>response:\n%v\n\n=>expected:\n%v\n", result, expected)
	}

}

func TestWhitObject(t *testing.T) {
	// agregamos objeto sin capos

	object_name := "object_without_fields"

	object_without_fields := model.Object{
		Name: object_name,
	}

	handler := model.Handlers{}
	handler.AddObjects(&object_without_fields)
	cutkey.AddDataConverter(&handler)

	expected := []map[string]string{
		{"description": "Manzanas", "price": "6000"},
		{"description": "Peras"},
	}

	out, err := cut.EncodeMaps(expected, object_name)
	if err != nil {
		t.Fatal("no se esperaba error en cut.EncodeMaps y se obtuvo", err, out)
	}

	result, err := cut.DecodeMaps(out)
	if err != nil {
		t.Fatal("no se esperaba error en cut.DecodeMaps y se obtuvo", err, result)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result:\n\n=>response:\n%v\n\n=>expected:\n%v\n", result, expected)
	}

}

func TestWhitObjectAndOneMapIN(t *testing.T) {
	// agregamos objeto sin capos y enviamos solo mapa

	object_name := "object_without_fields"

	object_without_fields := model.Object{
		Name: object_name,
	}

	handler := model.Handlers{}
	handler.AddObjects(&object_without_fields)
	cutkey.AddDataConverter(&handler)

	expected := map[string]string{"description": "Manzanas", "price": "6000"}

	out, err := cut.EncodeMaps(expected, object_name)
	if err != nil {
		t.Fatal("no se esperaba error en cut.EncodeMaps y se obtuvo", err, out)
	}

	result, err := cut.DecodeMaps(out)
	if err != nil {
		t.Fatal("no se esperaba error en cut.DecodeMaps y se obtuvo", err, result)
	}

	expect := []map[string]string{expected}

	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("result:\n\n=>response:\n%v\n\n=>expected:\n%v\n", result, expect)
	}

}
