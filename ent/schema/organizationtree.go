package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// OrganizationTree holds the schema definition for the OrganizationTree entity.
type OrganizationTree struct {
	ent.Schema
}

// Fields of the OrganizationTree.
func (OrganizationTree) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Int("parent_id").Unique(),
		field.Int("left").Unique(),
		field.Int("right").Unique(),
		field.Int("level").Unique(),
	}
}

// Edges of the OrganizationTree.
func (OrganizationTree) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organization_positions", OrganizationPosition.Type).
			Unique().Required(),
	}
}

// Mixin ....
func (OrganizationTree) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
