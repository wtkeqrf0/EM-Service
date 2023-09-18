package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// EnrichedFio holds the schema definition for the Author entity.
type EnrichedFio struct {
	ent.Schema
}

// Fields of the EnrichedFio.
func (EnrichedFio) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),

		field.String("surname").NotEmpty(),

		field.String("patronymic").Optional().Nillable().NotEmpty(),

		field.Int("age").NonNegative(),

		field.String("gender").NotEmpty(),

		field.String("country").NotEmpty(),
	}
}

func (EnrichedFio) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
