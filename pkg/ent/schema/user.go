package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the Author entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),

		field.String("surname").NotEmpty(),

		field.String("patronymic").Optional().Nillable().NotEmpty(),

		field.Int("age").NonNegative(),

		field.String("gender").NotEmpty(),

		field.String("country").NotEmpty(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
