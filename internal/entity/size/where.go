// Code generated by ent, DO NOT EDIT.

package size

import (
	"entgo.io/ent/dialect/sql"
	"github.com/dsha256/packer-pro/internal/entity/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Size {
	return predicate.Size(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Size {
	return predicate.Size(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Size {
	return predicate.Size(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Size {
	return predicate.Size(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Size {
	return predicate.Size(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Size {
	return predicate.Size(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Size {
	return predicate.Size(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Size {
	return predicate.Size(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Size {
	return predicate.Size(sql.FieldLTE(FieldID, id))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.Size {
	return predicate.Size(sql.FieldEQ(FieldSize, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.Size {
	return predicate.Size(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.Size {
	return predicate.Size(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.Size {
	return predicate.Size(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.Size {
	return predicate.Size(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.Size {
	return predicate.Size(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.Size {
	return predicate.Size(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.Size {
	return predicate.Size(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.Size {
	return predicate.Size(sql.FieldLTE(FieldSize, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Size) predicate.Size {
	return predicate.Size(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Size) predicate.Size {
	return predicate.Size(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Size) predicate.Size {
	return predicate.Size(sql.NotPredicates(p))
}
