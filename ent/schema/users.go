package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
    return []ent.Field{
        field.String("name"),
        field.String("email").Unique(),
        field.String("introduction"),
    }
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
