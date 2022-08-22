package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("address"),
		field.String("floor").Optional(),
		field.String("unitNo").Optional(),
		field.String("longitudeAndLatituude"),
		field.String("personCharge"),
		field.String("summary").Optional(),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
	}
}

// Indexes ....
func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

// Mixin ....
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
