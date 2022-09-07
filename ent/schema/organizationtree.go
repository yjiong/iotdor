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
		field.Int("parent_id"),
		field.Int("left"),
		field.Int("right"),
		field.Int("level"),
	}
}

// Edges of the OrganizationTree.
func (OrganizationTree) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organization_positions", OrganizationPosition.Type),
	}
}

// Mixin ....
func (OrganizationTree) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
