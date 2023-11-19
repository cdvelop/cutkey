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

	out, err := handler.EncodeMaps(expected)
	if err != nil {
		t.Fatal("no se esperaba error en cut.EncodeMaps y se obtuvo", err, out)
	}

	result, err := handler.DecodeMaps(out)
	if err != nil {
		t.Fatal("no se esperaba error en cut.DecodeMaps y se obtuvo", err, result)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result:\n\n=>response:\n%v\n\n=>expected:\n%v\n", result, expected)
	}

}

func TestWhitObjectWhitOutFieldsAndSliceMaps(t *testing.T) {
	// agregamos objeto sin capos y slice de maps

	object_name := "object_without_fields"

	object_without_fields := model.Object{
		ObjectName: object_name,
	}

	handler := model.Handlers{}
	handler.AddObjects(&object_without_fields)
	cutkey.AddDataConverter(&handler)

	expected := []map[string]string{
		{"description": "Manzanas", "price": "6000"},
		{"description": "Peras"},
	}

	out, err := handler.EncodeMaps(expected, object_name)
	if err != nil {
		t.Fatal("no se esperaba error en cut.EncodeMaps y se obtuvo", err, out)
	}

	result, err := handler.DecodeMaps(out)
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
		ObjectName: object_name,
	}

	handler := model.Handlers{}
	handler.AddObjects(&object_without_fields)
	cutkey.AddDataConverter(&handler)

	expected := map[string]string{"description": "Manzanas", "price": "6000"}

	out, err := handler.EncodeMaps(expected, object_name)
	if err != nil {
		t.Fatal("no se esperaba error en cut.EncodeMaps y se obtuvo", err, out)
	}

	result, err := handler.DecodeMaps(out)
	if err != nil {
		t.Fatal("no se esperaba error en cut.DecodeMaps y se obtuvo", err, result)
	}

	expect := []map[string]string{expected}

	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("result:\n\n=>response:\n%v\n\n=>expected:\n%v\n", result, expect)
	}

}

func TestObjectWhitFieldAndOneMapIN(t *testing.T) {
	// agregamos objeto  con  campos y enviamos solo mapa

	object_name := "product"

	object_without_fields := model.Object{
		ObjectName: object_name,
		Fields: []model.Field{
			{Name: "description"},
			{Name: "price"},
		},
	}

	handler := model.Handlers{}
	handler.AddObjects(&object_without_fields)
	cutkey.AddDataConverter(&handler)

	source_data := map[string]string{"description": "Manzanas", "price": "6000"}

	out, err := handler.EncodeMaps(source_data, object_name)
	if err != nil {
		t.Fatal("no se esperaba error en cut.EncodeMaps y se obtuvo", err, out)
	}

	result, err := handler.DecodeMaps(out, object_name)
	if err != nil {
		t.Fatal("no se esperaba error en cut.DecodeMaps y se obtuvo", err, result)
	}

	expected := []map[string]string{source_data}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result:\n\n=>response:\n%v\n\n=>expected:\n%v\n", result, expected)
	}

}

func TestObjectWhitFieldAndTwoMapsIN(t *testing.T) {
	// agregamos objeto  con  campos y enviamos 2 mapas

	object_name := "product"

	object_without_fields := model.Object{
		ObjectName: object_name,
		Fields: []model.Field{
			{Name: "description"},
			{Name: "price"},
		},
	}

	handler := model.Handlers{}
	handler.AddObjects(&object_without_fields)
	cutkey.AddDataConverter(&handler)

	expected := []map[string]string{
		{"description": "Manzanas", "price": "6000"},
		{"description": "Peras"},
	}

	out, err := handler.EncodeMaps(expected, object_name)
	if err != nil {
		t.Fatal("no se esperaba error en cut.EncodeMaps y se obtuvo", err, out)
	}

	result, err := handler.DecodeMaps(out, object_name)
	if err != nil {
		t.Fatal("no se esperaba error en cut.DecodeMaps y se obtuvo", err, result)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result:\n\n=>response:\n%v\n\n=>expected:\n%v\n", result, expected)
	}

}
