package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("devID"),
		field.String("devType"),
		field.String("devAddr"),
		field.String("conn"),
		field.String("name").
			Optional(),
		field.Bool("idDelete"),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gateway", Gateway.Type).
			Ref("devices").
			Unique(),
	}
}

// Mixin of the Dev
func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Indexes of the Device
func (Device) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("devID"),
	}
}
