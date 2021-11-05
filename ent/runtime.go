// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/fitan/magic/ent/schema"
	"github.com/fitan/magic/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[6].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(int8)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[8].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(int32)
	// userDescOnlyOss is the schema descriptor for only_oss field.
	userDescOnlyOss := userFields[12].Descriptor()
	// user.DefaultOnlyOss holds the default value on creation for the only_oss field.
	user.DefaultOnlyOss = userDescOnlyOss.Default.(int8)
}