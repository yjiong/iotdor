package schema

import "entgo.io/ent"

// Gateway holds the schema definition for the Gateway entity.
type Gateway struct {
	ent.Schema
}

// Fields of the Gateway.
func (Gateway) Fields() []ent.Field {
	return nil
}

// Edges of the Gateway.
func (Gateway) Edges() []ent.Edge {
	return nil
}
