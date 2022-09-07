package kcschema

import "time"

type Operation string

const (
	Read   Operation = "r"
	Create           = "c"
	Update           = "u"
	Delete           = "d"
)

type DBZValue struct {
	Before      Payload     `json:"before,omitempty"`
	After       Payload     `json:"after,omitempty"`
	Op          Operation   `json:"op,omitempty"`
	Patch       string      `json:"patch,omitempty"`
	Filter      string      `json:"filter,omitempty"`
	Source      interface{} `json:"source,omitempty"`
	Transaction interface{} `json:"transaction,omitempty"`
	TimestampMS time.Time   `json:"ts_ms"`
}
