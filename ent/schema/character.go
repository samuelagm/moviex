package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
		field.String("height"),
		field.String("mass"),
		field.String("hair_color"),
		field.String("skin_color"),
		field.String("eye_color"),
		field.String("birth_year"),
		field.String("gender"),
		field.Strings("films"),
		field.Time("created"),
		field.Time("edited"),
		field.String("url"),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("film", Movie.Type).Ref("people"),
	}
}
