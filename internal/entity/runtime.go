// Code generated by ent, DO NOT EDIT.

package entity

import (
	"github.com/dsha256/packer-pro/internal/entity/schema"
	"github.com/dsha256/packer-pro/internal/entity/size"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sizeFields := schema.Size{}.Fields()
	_ = sizeFields
	// sizeDescSize is the schema descriptor for size field.
	sizeDescSize := sizeFields[1].Descriptor()
	// size.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	size.SizeValidator = sizeDescSize.Validators[0].(func(int) error)
}
