package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Match(regexp.MustCompile("[a-zA-Z_]+$")),
		field.Text("summary").
			Optional(),
		//Validate(func(s string) error {
		//if strings.ToLower(s) == s {
		//return errors.New("group name must begin with uppercase")
		//}
		//return nil
		//}),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("admins", User.Type),
		edge.To("gateways", Gateway.Type),
	}
}
