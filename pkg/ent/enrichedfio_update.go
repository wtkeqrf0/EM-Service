// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/restService/pkg/ent/enrichedfio"
	"github.com/wtkeqrf0/restService/pkg/ent/predicate"
)

// EnrichedFioUpdate is the builder for updating EnrichedFio entities.
type EnrichedFioUpdate struct {
	config
	hooks    []Hook
	mutation *EnrichedFioMutation
}

// Where appends a list predicates to the EnrichedFioUpdate builder.
func (efu *EnrichedFioUpdate) Where(ps ...predicate.EnrichedFio) *EnrichedFioUpdate {
	efu.mutation.Where(ps...)
	return efu
}

// SetUpdateTime sets the "update_time" field.
func (efu *EnrichedFioUpdate) SetUpdateTime(t time.Time) *EnrichedFioUpdate {
	efu.mutation.SetUpdateTime(t)
	return efu
}

// SetName sets the "name" field.
func (efu *EnrichedFioUpdate) SetName(s string) *EnrichedFioUpdate {
	efu.mutation.SetName(s)
	return efu
}

// SetSurname sets the "surname" field.
func (efu *EnrichedFioUpdate) SetSurname(s string) *EnrichedFioUpdate {
	efu.mutation.SetSurname(s)
	return efu
}

// SetPatronymic sets the "patronymic" field.
func (efu *EnrichedFioUpdate) SetPatronymic(s string) *EnrichedFioUpdate {
	efu.mutation.SetPatronymic(s)
	return efu
}

// SetNillablePatronymic sets the "patronymic" field if the given value is not nil.
func (efu *EnrichedFioUpdate) SetNillablePatronymic(s *string) *EnrichedFioUpdate {
	if s != nil {
		efu.SetPatronymic(*s)
	}
	return efu
}

// ClearPatronymic clears the value of the "patronymic" field.
func (efu *EnrichedFioUpdate) ClearPatronymic() *EnrichedFioUpdate {
	efu.mutation.ClearPatronymic()
	return efu
}

// SetAge sets the "age" field.
func (efu *EnrichedFioUpdate) SetAge(i int) *EnrichedFioUpdate {
	efu.mutation.ResetAge()
	efu.mutation.SetAge(i)
	return efu
}

// AddAge adds i to the "age" field.
func (efu *EnrichedFioUpdate) AddAge(i int) *EnrichedFioUpdate {
	efu.mutation.AddAge(i)
	return efu
}

// SetGender sets the "gender" field.
func (efu *EnrichedFioUpdate) SetGender(s string) *EnrichedFioUpdate {
	efu.mutation.SetGender(s)
	return efu
}

// SetCountry sets the "country" field.
func (efu *EnrichedFioUpdate) SetCountry(s string) *EnrichedFioUpdate {
	efu.mutation.SetCountry(s)
	return efu
}

// Mutation returns the EnrichedFioMutation object of the builder.
func (efu *EnrichedFioUpdate) Mutation() *EnrichedFioMutation {
	return efu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (efu *EnrichedFioUpdate) Save(ctx context.Context) (int, error) {
	efu.defaults()
	return withHooks(ctx, efu.sqlSave, efu.mutation, efu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (efu *EnrichedFioUpdate) SaveX(ctx context.Context) int {
	affected, err := efu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (efu *EnrichedFioUpdate) Exec(ctx context.Context) error {
	_, err := efu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efu *EnrichedFioUpdate) ExecX(ctx context.Context) {
	if err := efu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (efu *EnrichedFioUpdate) defaults() {
	if _, ok := efu.mutation.UpdateTime(); !ok {
		v := enrichedfio.UpdateDefaultUpdateTime()
		efu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (efu *EnrichedFioUpdate) check() error {
	if v, ok := efu.mutation.Name(); ok {
		if err := enrichedfio.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.name": %w`, err)}
		}
	}
	if v, ok := efu.mutation.Surname(); ok {
		if err := enrichedfio.SurnameValidator(v); err != nil {
			return &ValidationError{Name: "surname", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.surname": %w`, err)}
		}
	}
	if v, ok := efu.mutation.Patronymic(); ok {
		if err := enrichedfio.PatronymicValidator(v); err != nil {
			return &ValidationError{Name: "patronymic", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.patronymic": %w`, err)}
		}
	}
	if v, ok := efu.mutation.Age(); ok {
		if err := enrichedfio.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.age": %w`, err)}
		}
	}
	if v, ok := efu.mutation.Gender(); ok {
		if err := enrichedfio.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.gender": %w`, err)}
		}
	}
	if v, ok := efu.mutation.Country(); ok {
		if err := enrichedfio.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.country": %w`, err)}
		}
	}
	return nil
}

func (efu *EnrichedFioUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := efu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(enrichedfio.Table, enrichedfio.Columns, sqlgraph.NewFieldSpec(enrichedfio.FieldID, field.TypeInt))
	if ps := efu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := efu.mutation.UpdateTime(); ok {
		_spec.SetField(enrichedfio.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := efu.mutation.Name(); ok {
		_spec.SetField(enrichedfio.FieldName, field.TypeString, value)
	}
	if value, ok := efu.mutation.Surname(); ok {
		_spec.SetField(enrichedfio.FieldSurname, field.TypeString, value)
	}
	if value, ok := efu.mutation.Patronymic(); ok {
		_spec.SetField(enrichedfio.FieldPatronymic, field.TypeString, value)
	}
	if efu.mutation.PatronymicCleared() {
		_spec.ClearField(enrichedfio.FieldPatronymic, field.TypeString)
	}
	if value, ok := efu.mutation.Age(); ok {
		_spec.SetField(enrichedfio.FieldAge, field.TypeInt, value)
	}
	if value, ok := efu.mutation.AddedAge(); ok {
		_spec.AddField(enrichedfio.FieldAge, field.TypeInt, value)
	}
	if value, ok := efu.mutation.Gender(); ok {
		_spec.SetField(enrichedfio.FieldGender, field.TypeString, value)
	}
	if value, ok := efu.mutation.Country(); ok {
		_spec.SetField(enrichedfio.FieldCountry, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, efu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enrichedfio.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	efu.mutation.done = true
	return n, nil
}

// EnrichedFioUpdateOne is the builder for updating a single EnrichedFio entity.
type EnrichedFioUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EnrichedFioMutation
}

// SetUpdateTime sets the "update_time" field.
func (efuo *EnrichedFioUpdateOne) SetUpdateTime(t time.Time) *EnrichedFioUpdateOne {
	efuo.mutation.SetUpdateTime(t)
	return efuo
}

// SetName sets the "name" field.
func (efuo *EnrichedFioUpdateOne) SetName(s string) *EnrichedFioUpdateOne {
	efuo.mutation.SetName(s)
	return efuo
}

// SetSurname sets the "surname" field.
func (efuo *EnrichedFioUpdateOne) SetSurname(s string) *EnrichedFioUpdateOne {
	efuo.mutation.SetSurname(s)
	return efuo
}

// SetPatronymic sets the "patronymic" field.
func (efuo *EnrichedFioUpdateOne) SetPatronymic(s string) *EnrichedFioUpdateOne {
	efuo.mutation.SetPatronymic(s)
	return efuo
}

// SetNillablePatronymic sets the "patronymic" field if the given value is not nil.
func (efuo *EnrichedFioUpdateOne) SetNillablePatronymic(s *string) *EnrichedFioUpdateOne {
	if s != nil {
		efuo.SetPatronymic(*s)
	}
	return efuo
}

// ClearPatronymic clears the value of the "patronymic" field.
func (efuo *EnrichedFioUpdateOne) ClearPatronymic() *EnrichedFioUpdateOne {
	efuo.mutation.ClearPatronymic()
	return efuo
}

// SetAge sets the "age" field.
func (efuo *EnrichedFioUpdateOne) SetAge(i int) *EnrichedFioUpdateOne {
	efuo.mutation.ResetAge()
	efuo.mutation.SetAge(i)
	return efuo
}

// AddAge adds i to the "age" field.
func (efuo *EnrichedFioUpdateOne) AddAge(i int) *EnrichedFioUpdateOne {
	efuo.mutation.AddAge(i)
	return efuo
}

// SetGender sets the "gender" field.
func (efuo *EnrichedFioUpdateOne) SetGender(s string) *EnrichedFioUpdateOne {
	efuo.mutation.SetGender(s)
	return efuo
}

// SetCountry sets the "country" field.
func (efuo *EnrichedFioUpdateOne) SetCountry(s string) *EnrichedFioUpdateOne {
	efuo.mutation.SetCountry(s)
	return efuo
}

// Mutation returns the EnrichedFioMutation object of the builder.
func (efuo *EnrichedFioUpdateOne) Mutation() *EnrichedFioMutation {
	return efuo.mutation
}

// Where appends a list predicates to the EnrichedFioUpdate builder.
func (efuo *EnrichedFioUpdateOne) Where(ps ...predicate.EnrichedFio) *EnrichedFioUpdateOne {
	efuo.mutation.Where(ps...)
	return efuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (efuo *EnrichedFioUpdateOne) Select(field string, fields ...string) *EnrichedFioUpdateOne {
	efuo.fields = append([]string{field}, fields...)
	return efuo
}

// Save executes the query and returns the updated EnrichedFio entity.
func (efuo *EnrichedFioUpdateOne) Save(ctx context.Context) (*EnrichedFio, error) {
	efuo.defaults()
	return withHooks(ctx, efuo.sqlSave, efuo.mutation, efuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (efuo *EnrichedFioUpdateOne) SaveX(ctx context.Context) *EnrichedFio {
	node, err := efuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (efuo *EnrichedFioUpdateOne) Exec(ctx context.Context) error {
	_, err := efuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efuo *EnrichedFioUpdateOne) ExecX(ctx context.Context) {
	if err := efuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (efuo *EnrichedFioUpdateOne) defaults() {
	if _, ok := efuo.mutation.UpdateTime(); !ok {
		v := enrichedfio.UpdateDefaultUpdateTime()
		efuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (efuo *EnrichedFioUpdateOne) check() error {
	if v, ok := efuo.mutation.Name(); ok {
		if err := enrichedfio.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.name": %w`, err)}
		}
	}
	if v, ok := efuo.mutation.Surname(); ok {
		if err := enrichedfio.SurnameValidator(v); err != nil {
			return &ValidationError{Name: "surname", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.surname": %w`, err)}
		}
	}
	if v, ok := efuo.mutation.Patronymic(); ok {
		if err := enrichedfio.PatronymicValidator(v); err != nil {
			return &ValidationError{Name: "patronymic", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.patronymic": %w`, err)}
		}
	}
	if v, ok := efuo.mutation.Age(); ok {
		if err := enrichedfio.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.age": %w`, err)}
		}
	}
	if v, ok := efuo.mutation.Gender(); ok {
		if err := enrichedfio.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.gender": %w`, err)}
		}
	}
	if v, ok := efuo.mutation.Country(); ok {
		if err := enrichedfio.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "EnrichedFio.country": %w`, err)}
		}
	}
	return nil
}

func (efuo *EnrichedFioUpdateOne) sqlSave(ctx context.Context) (_node *EnrichedFio, err error) {
	if err := efuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(enrichedfio.Table, enrichedfio.Columns, sqlgraph.NewFieldSpec(enrichedfio.FieldID, field.TypeInt))
	id, ok := efuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EnrichedFio.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := efuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enrichedfio.FieldID)
		for _, f := range fields {
			if !enrichedfio.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != enrichedfio.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := efuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := efuo.mutation.UpdateTime(); ok {
		_spec.SetField(enrichedfio.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := efuo.mutation.Name(); ok {
		_spec.SetField(enrichedfio.FieldName, field.TypeString, value)
	}
	if value, ok := efuo.mutation.Surname(); ok {
		_spec.SetField(enrichedfio.FieldSurname, field.TypeString, value)
	}
	if value, ok := efuo.mutation.Patronymic(); ok {
		_spec.SetField(enrichedfio.FieldPatronymic, field.TypeString, value)
	}
	if efuo.mutation.PatronymicCleared() {
		_spec.ClearField(enrichedfio.FieldPatronymic, field.TypeString)
	}
	if value, ok := efuo.mutation.Age(); ok {
		_spec.SetField(enrichedfio.FieldAge, field.TypeInt, value)
	}
	if value, ok := efuo.mutation.AddedAge(); ok {
		_spec.AddField(enrichedfio.FieldAge, field.TypeInt, value)
	}
	if value, ok := efuo.mutation.Gender(); ok {
		_spec.SetField(enrichedfio.FieldGender, field.TypeString, value)
	}
	if value, ok := efuo.mutation.Country(); ok {
		_spec.SetField(enrichedfio.FieldCountry, field.TypeString, value)
	}
	_node = &EnrichedFio{config: efuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, efuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enrichedfio.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	efuo.mutation.done = true
	return _node, nil
}
