package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {

	return []ent.Field{

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),

		field.String("displayname").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),

		field.String("password").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),

		field.String("email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),

		field.String("phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),

		field.Int32("login_frequency").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional(),

		field.Int8("active").SchemaType(map[string]string{
			dialect.MySQL: "tinyint", // Override MySQL.
		}).Optional().Default(1),

		field.String("api_token").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),

		field.Int32("role").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Default(1),

		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional(),

		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional(),

		field.Int8("auth_type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint", // Override MySQL.
		}).Optional(),

		field.Int8("only_oss").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Optional().Default(0),
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
