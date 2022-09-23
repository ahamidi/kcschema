package kcschema

import (
	"log"
	"testing"
)

func TestPayload_KCSType(t *testing.T) {
	r := sampleRecordWithSchema()
	p := Payload(r)

	pt := p.Type()
	if pt != KCJSONWithSchemaType {
		t.Errorf("expected %s, got %s", KCJSONWithSchemaType, pt)
	}
}

func TestPayload_KCSParse(t *testing.T) {
	r := sampleRecordWithSchema()
	p := Payload(r)

	sp, err := Parse(p)
	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}
	log.Printf("sp: %+v", sp)
}

func TestStructuredPayload_AsKCSchemaJSON(t *testing.T) {
	sp := StructuredPayload{
		"id":    Field{Value: 1, Type: IntField},
		"email": Field{Value: "alice@example.com", Type: StringField},
	}

	j, err := sp.AsKCSchemaJSON("users")
	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}

	log.Printf("j: %+v", string(j))
}

func sampleRecordWithSchema() []byte {
	return []byte(`{
	"schema": {
		"type": "struct",
		"name": "customer_order",
		"fields": [
			{
				"field": "id",
				"type": "int32"
			},
			{
				"field": "related_item_ids",
				"type": "array"
			}
		]
	},
	"payload": {
		"id": 1,
		"related_item_ids": [83,287,903]
	}
}`)
}
