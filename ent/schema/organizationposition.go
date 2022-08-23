package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// OrganizationPosition holds the schema definition for the Organization entity.
type OrganizationPosition struct {
	ent.Schema
}

// Fields of the OrganizationPosition.
func (OrganizationPosition) Fields() []ent.Field {
	return []ent.Field{
		field.String("position_id").Unique(),
		field.String("address"),
		field.String("floor").Optional(),
		field.String("unit_no").Optional(),
		field.String("longitude_and_latitude"),
		field.String("summary").Optional(),
	}
}

// Edges of the OrganizationPosition.
func (OrganizationPosition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
		edge.To("person_charges", User.Type),
		edge.From("organization_tree", OrganizationTree.Type).
			Ref("organization_positions"),
	}
}

// Indexes ....
func (OrganizationPosition) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("position_id"),
	}
}

// Mixin ....
func (OrganizationPosition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
