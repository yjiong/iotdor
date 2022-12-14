package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Gateway holds the schema definition for the Gateway entity.
type Gateway struct {
	ent.Schema
}

// Fields of the Gateway.
func (Gateway) Fields() []ent.Field {
	return []ent.Field{
		field.String("gwid").
			Unique(),
		field.String("svrid"),
		field.String("broker"),
		field.String("installation_location").
			Optional(),
		field.String("stat").
			Optional(),
		field.Bool("delete_flag").
			Optional(),
		field.Int("up_interval").
			Default(60),
		field.String("version").
			Optional(),
		field.String("summary").
			Optional(),
	}
}

// Edges of the Gateway.
func (Gateway) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
		edge.From("group", Group.Type).
			Ref("gateways").
			Unique(),
	}
}

// Indexes of the GW
func (Gateway) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		//index.Fields("field1", "field2"),
		// 唯一约束索引
		index.Fields("gwid").
			Unique(),
	}
}

// Mixin of the GW
func (Gateway) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
