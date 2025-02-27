// Code generated by ent, DO NOT EDIT.

package metadata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-uem/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Metadata {
	return predicate.Metadata(sql.FieldLTE(FieldID, id))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldValue, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Metadata {
	return predicate.Metadata(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Metadata {
	return predicate.Metadata(sql.FieldContainsFold(FieldValue, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrg applies the HasEdge predicate on the "org" edge.
func HasOrg() predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrgTable, OrgColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgWith applies the HasEdge predicate on the "org" edge with a given conditions (other predicates).
func HasOrgWith(preds ...predicate.OrgMetadata) predicate.Metadata {
	return predicate.Metadata(func(s *sql.Selector) {
		step := newOrgStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Metadata) predicate.Metadata {
	return predicate.Metadata(sql.NotPredicates(p))
}
