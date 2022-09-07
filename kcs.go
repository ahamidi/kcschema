package kcschema

import "encoding/json"

type Value struct {
	Schema  Schema  `json:"schema"`
	Payload Payload `json:"payload"`
}

type SchemaField struct {
	Field    string `json:"field,omitempty"`
	Type     string `json:"type,omitempty"`
	Optional bool   `json:"optional,omitempty"`
}

type Schema struct {
	Type     string        `json:"type,omitempty"`
	Name     string        `json:"name,omitempty"`
	Optional bool          `json:"optional,omitempty"`
	Fields   []SchemaField `json:"fields,omitempty"`
}

func (v *Value) MarshalJSON() ([]byte, error) {
	type ValAlias Value
	return json.Marshal(&struct {
		*ValAlias
	}{
		ValAlias: (*ValAlias)(v),
	})
}
