package kcschema

import (
	"encoding/json"
	"log"
	"testing"
)

func TestValue_MarshalJSON(t *testing.T) {
	t1 := Value{
		Schema: Schema{
			Type:     "struct",
			Name:     "customer_order",
			Optional: false,
			Fields: []SchemaField{
				{Field: "id",
					Optional: false,
					Type:     "int32",
				},
			},
		},
		Payload: Payload(`{
			"id": 1
		}`),
	}

	b, err := json.Marshal(t1)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	log.Printf("Empty: %+v", string(b))
	sp, err := t1.Payload.ParseAsJSON()
	log.Printf("JSON: %+v", sp)
}

func TestValue_UnmarshalJSON(t *testing.T) {
	b := sampleRecordWithSchema()
	var val Value
	json.Unmarshal(b, &val)
	log.Printf("Schema: %+v", val.Schema)
	log.Printf("Payload: %+v", val.Payload)
}

func sampleRecordWithSchema() []byte {
	return []byte(`{
	"schema": {
		"type": "struct",
		"name": "customer_order",
		"fields": [{
			"field": "id",
			"type": "int32"
		}]
	},
	"payload": {
		"id": 1
	}
}`)
}
