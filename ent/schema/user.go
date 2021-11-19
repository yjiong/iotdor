package schema

import (
	"entgo.io/ent"
	//"entgo.io/ent/dialect/entsql"
	//"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("passwd").
			Sensitive(),
		field.String("phone").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).
			Ref("users"),
		edge.From("manage", Group.Type).
			Ref("admin"),
	}
}

// Mixin of the User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Annotations ...
//func (User) Annotations() []schema.Annotation {
//return []schema.Annotation{
//entsql.Annotation{Table: "Users"},
//}
//}
