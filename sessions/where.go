// Code generated by ent, DO NOT EDIT.

package sessions

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-uem/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Sessions {
	return predicate.Sessions(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Sessions {
	return predicate.Sessions(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Sessions {
	return predicate.Sessions(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Sessions {
	return predicate.Sessions(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Sessions {
	return predicate.Sessions(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Sessions {
	return predicate.Sessions(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Sessions {
	return predicate.Sessions(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Sessions {
	return predicate.Sessions(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Sessions {
	return predicate.Sessions(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Sessions {
	return predicate.Sessions(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Sessions {
	return predicate.Sessions(sql.FieldContainsFold(FieldID, id))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v []byte) predicate.Sessions {
	return predicate.Sessions(sql.FieldEQ(FieldData, v))
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.Sessions {
	return predicate.Sessions(sql.FieldEQ(FieldExpiry, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v []byte) predicate.Sessions {
	return predicate.Sessions(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v []byte) predicate.Sessions {
	return predicate.Sessions(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...[]byte) predicate.Sessions {
	return predicate.Sessions(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...[]byte) predicate.Sessions {
	return predicate.Sessions(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v []byte) predicate.Sessions {
	return predicate.Sessions(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v []byte) predicate.Sessions {
	return predicate.Sessions(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v []byte) predicate.Sessions {
	return predicate.Sessions(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v []byte) predicate.Sessions {
	return predicate.Sessions(sql.FieldLTE(FieldData, v))
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.Sessions {
	return predicate.Sessions(sql.FieldEQ(FieldExpiry, v))
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.Sessions {
	return predicate.Sessions(sql.FieldNEQ(FieldExpiry, v))
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.Sessions {
	return predicate.Sessions(sql.FieldIn(FieldExpiry, vs...))
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.Sessions {
	return predicate.Sessions(sql.FieldNotIn(FieldExpiry, vs...))
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.Sessions {
	return predicate.Sessions(sql.FieldGT(FieldExpiry, v))
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.Sessions {
	return predicate.Sessions(sql.FieldGTE(FieldExpiry, v))
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.Sessions {
	return predicate.Sessions(sql.FieldLT(FieldExpiry, v))
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.Sessions {
	return predicate.Sessions(sql.FieldLTE(FieldExpiry, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Sessions {
	return predicate.Sessions(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Sessions {
	return predicate.Sessions(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Sessions) predicate.Sessions {
	return predicate.Sessions(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Sessions) predicate.Sessions {
	return predicate.Sessions(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Sessions) predicate.Sessions {
	return predicate.Sessions(sql.NotPredicates(p))
}
