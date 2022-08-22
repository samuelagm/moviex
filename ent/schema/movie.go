package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Default("unknown"),
		field.Int("episode_id").
			Positive(),
		field.String("opening_crawl").Default(""),
		field.String("director"),
		field.String("producer"),
		field.Time("release_date"),
		field.Strings("characters"),
		field.Time("created"),
		field.Time("edited"),
		field.String("url"),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("people", Character.Type),
		edge.To("comments", Comment.Type),
	}
}
