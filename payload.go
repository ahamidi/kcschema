package kcschema

import (
	"encoding/json"
	"log"
	"reflect"
)

type FieldType string

const (
	StringField FieldType = "string"
	IntField    FieldType = "int"
	FloatField  FieldType = "float"
	MapField    FieldType = "map"
)

type FieldInterface interface {
	Get() interface{}
	Set(value FieldInterface) error
	Type() FieldType
}

type PayloadInterface interface {
	Get(field string) FieldInterface
	Set(field FieldInterface) error
	Type(field string) FieldType
}

type Payload []byte

type Field struct {
	Type  FieldType
	Value interface{}
}

type StructuredPayload map[string]Field

func (p Payload) ParseAsJSON() (StructuredPayload, error) {
	sp := make(map[string]Field)
	var m map[string]interface{}
	err := json.Unmarshal(p, &m)

	for f, v := range m {
		log.Printf("f: %s, v: %+v", f, v)
		log.Printf("parsed: %+v", parseField(v))
		sp[f] = *parseField(v)
	}
	return sp, err
}

func parseField(v interface{}) *Field {
	switch v.(type) {
	case map[string]interface{}:
		return &Field{
			Type:  MapField,
			Value: parseField(v),
		}
	case string:
		return &Field{
			Type:  StringField,
			Value: v.(string),
		}
	case int, int32, int64:
		return &Field{
			Type:  IntField,
			Value: v.(int),
		}
	case float32, float64:
		return &Field{
			Type:  FloatField,
			Value: v.(float64),
		}
	default:
		log.Printf("v: %+v, type: %s", v, reflect.TypeOf(v).String())
	}
	return nil
}
