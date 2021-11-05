// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fitan/magic/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetDisplayname sets the "displayname" field.
func (uc *UserCreate) SetDisplayname(s string) *UserCreate {
	uc.mutation.SetDisplayname(s)
	return uc
}

// SetNillableDisplayname sets the "displayname" field if the given value is not nil.
func (uc *UserCreate) SetNillableDisplayname(s *string) *UserCreate {
	if s != nil {
		uc.SetDisplayname(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetLoginFrequency sets the "login_frequency" field.
func (uc *UserCreate) SetLoginFrequency(i int32) *UserCreate {
	uc.mutation.SetLoginFrequency(i)
	return uc
}

// SetNillableLoginFrequency sets the "login_frequency" field if the given value is not nil.
func (uc *UserCreate) SetNillableLoginFrequency(i *int32) *UserCreate {
	if i != nil {
		uc.SetLoginFrequency(*i)
	}
	return uc
}

// SetActive sets the "active" field.
func (uc *UserCreate) SetActive(i int8) *UserCreate {
	uc.mutation.SetActive(i)
	return uc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uc *UserCreate) SetNillableActive(i *int8) *UserCreate {
	if i != nil {
		uc.SetActive(*i)
	}
	return uc
}

// SetAPIToken sets the "api_token" field.
func (uc *UserCreate) SetAPIToken(s string) *UserCreate {
	uc.mutation.SetAPIToken(s)
	return uc
}

// SetNillableAPIToken sets the "api_token" field if the given value is not nil.
func (uc *UserCreate) SetNillableAPIToken(s *string) *UserCreate {
	if s != nil {
		uc.SetAPIToken(*s)
	}
	return uc
}

// SetRole sets the "role" field.
func (uc *UserCreate) SetRole(i int32) *UserCreate {
	uc.mutation.SetRole(i)
	return uc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uc *UserCreate) SetNillableRole(i *int32) *UserCreate {
	if i != nil {
		uc.SetRole(*i)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetAuthType sets the "auth_type" field.
func (uc *UserCreate) SetAuthType(i int8) *UserCreate {
	uc.mutation.SetAuthType(i)
	return uc
}

// SetNillableAuthType sets the "auth_type" field if the given value is not nil.
func (uc *UserCreate) SetNillableAuthType(i *int8) *UserCreate {
	if i != nil {
		uc.SetAuthType(*i)
	}
	return uc
}

// SetOnlyOss sets the "only_oss" field.
func (uc *UserCreate) SetOnlyOss(i int8) *UserCreate {
	uc.mutation.SetOnlyOss(i)
	return uc
}

// SetNillableOnlyOss sets the "only_oss" field if the given value is not nil.
func (uc *UserCreate) SetNillableOnlyOss(i *int8) *UserCreate {
	if i != nil {
		uc.SetOnlyOss(*i)
	}
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Active(); !ok {
		v := user.DefaultActive
		uc.mutation.SetActive(v)
	}
	if _, ok := uc.mutation.Role(); !ok {
		v := user.DefaultRole
		uc.mutation.SetRole(v)
	}
	if _, ok := uc.mutation.OnlyOss(); !ok {
		v := user.DefaultOnlyOss
		uc.mutation.SetOnlyOss(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		_node.Name = value
	}
	if value, ok := uc.mutation.Displayname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDisplayname,
		})
		_node.Displayname = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := uc.mutation.LoginFrequency(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldLoginFrequency,
		})
		_node.LoginFrequency = value
	}
	if value, ok := uc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: user.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := uc.mutation.APIToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAPIToken,
		})
		_node.APIToken = value
	}
	if value, ok := uc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldRole,
		})
		_node.Role = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.AuthType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: user.FieldAuthType,
		})
		_node.AuthType = value
	}
	if value, ok := uc.mutation.OnlyOss(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: user.FieldOnlyOss,
		})
		_node.OnlyOss = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
