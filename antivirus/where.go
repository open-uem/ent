// Code generated by ent, DO NOT EDIT.

package antivirus

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-uem/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldEQ(FieldName, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldEQ(FieldIsActive, v))
}

// IsUpdated applies equality check predicate on the "is_updated" field. It's identical to IsUpdatedEQ.
func IsUpdated(v bool) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldEQ(FieldIsUpdated, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldContainsFold(FieldName, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldNEQ(FieldIsActive, v))
}

// IsUpdatedEQ applies the EQ predicate on the "is_updated" field.
func IsUpdatedEQ(v bool) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldEQ(FieldIsUpdated, v))
}

// IsUpdatedNEQ applies the NEQ predicate on the "is_updated" field.
func IsUpdatedNEQ(v bool) predicate.Antivirus {
	return predicate.Antivirus(sql.FieldNEQ(FieldIsUpdated, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Antivirus {
	return predicate.Antivirus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.Antivirus {
	return predicate.Antivirus(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Antivirus) predicate.Antivirus {
	return predicate.Antivirus(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Antivirus) predicate.Antivirus {
	return predicate.Antivirus(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Antivirus) predicate.Antivirus {
	return predicate.Antivirus(sql.NotPredicates(p))
}
