// Code generated by ent, DO NOT EDIT.

package systemupdate

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/doncicuto/openuem_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldLTE(FieldID, id))
}

// SystemUpdateStatus applies equality check predicate on the "system_update_status" field. It's identical to SystemUpdateStatusEQ.
func SystemUpdateStatus(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldSystemUpdateStatus, v))
}

// LastInstall applies equality check predicate on the "last_install" field. It's identical to LastInstallEQ.
func LastInstall(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldLastInstall, v))
}

// LastSearch applies equality check predicate on the "last_search" field. It's identical to LastSearchEQ.
func LastSearch(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldLastSearch, v))
}

// PendingUpdates applies equality check predicate on the "pending_updates" field. It's identical to PendingUpdatesEQ.
func PendingUpdates(v bool) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldPendingUpdates, v))
}

// SystemUpdateStatusEQ applies the EQ predicate on the "system_update_status" field.
func SystemUpdateStatusEQ(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusNEQ applies the NEQ predicate on the "system_update_status" field.
func SystemUpdateStatusNEQ(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldNEQ(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusIn applies the In predicate on the "system_update_status" field.
func SystemUpdateStatusIn(vs ...string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldIn(FieldSystemUpdateStatus, vs...))
}

// SystemUpdateStatusNotIn applies the NotIn predicate on the "system_update_status" field.
func SystemUpdateStatusNotIn(vs ...string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldNotIn(FieldSystemUpdateStatus, vs...))
}

// SystemUpdateStatusGT applies the GT predicate on the "system_update_status" field.
func SystemUpdateStatusGT(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldGT(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusGTE applies the GTE predicate on the "system_update_status" field.
func SystemUpdateStatusGTE(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldGTE(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusLT applies the LT predicate on the "system_update_status" field.
func SystemUpdateStatusLT(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldLT(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusLTE applies the LTE predicate on the "system_update_status" field.
func SystemUpdateStatusLTE(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldLTE(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusContains applies the Contains predicate on the "system_update_status" field.
func SystemUpdateStatusContains(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldContains(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusHasPrefix applies the HasPrefix predicate on the "system_update_status" field.
func SystemUpdateStatusHasPrefix(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldHasPrefix(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusHasSuffix applies the HasSuffix predicate on the "system_update_status" field.
func SystemUpdateStatusHasSuffix(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldHasSuffix(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusEqualFold applies the EqualFold predicate on the "system_update_status" field.
func SystemUpdateStatusEqualFold(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEqualFold(FieldSystemUpdateStatus, v))
}

// SystemUpdateStatusContainsFold applies the ContainsFold predicate on the "system_update_status" field.
func SystemUpdateStatusContainsFold(v string) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldContainsFold(FieldSystemUpdateStatus, v))
}

// LastInstallEQ applies the EQ predicate on the "last_install" field.
func LastInstallEQ(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldLastInstall, v))
}

// LastInstallNEQ applies the NEQ predicate on the "last_install" field.
func LastInstallNEQ(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldNEQ(FieldLastInstall, v))
}

// LastInstallIn applies the In predicate on the "last_install" field.
func LastInstallIn(vs ...time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldIn(FieldLastInstall, vs...))
}

// LastInstallNotIn applies the NotIn predicate on the "last_install" field.
func LastInstallNotIn(vs ...time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldNotIn(FieldLastInstall, vs...))
}

// LastInstallGT applies the GT predicate on the "last_install" field.
func LastInstallGT(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldGT(FieldLastInstall, v))
}

// LastInstallGTE applies the GTE predicate on the "last_install" field.
func LastInstallGTE(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldGTE(FieldLastInstall, v))
}

// LastInstallLT applies the LT predicate on the "last_install" field.
func LastInstallLT(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldLT(FieldLastInstall, v))
}

// LastInstallLTE applies the LTE predicate on the "last_install" field.
func LastInstallLTE(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldLTE(FieldLastInstall, v))
}

// LastSearchEQ applies the EQ predicate on the "last_search" field.
func LastSearchEQ(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldLastSearch, v))
}

// LastSearchNEQ applies the NEQ predicate on the "last_search" field.
func LastSearchNEQ(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldNEQ(FieldLastSearch, v))
}

// LastSearchIn applies the In predicate on the "last_search" field.
func LastSearchIn(vs ...time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldIn(FieldLastSearch, vs...))
}

// LastSearchNotIn applies the NotIn predicate on the "last_search" field.
func LastSearchNotIn(vs ...time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldNotIn(FieldLastSearch, vs...))
}

// LastSearchGT applies the GT predicate on the "last_search" field.
func LastSearchGT(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldGT(FieldLastSearch, v))
}

// LastSearchGTE applies the GTE predicate on the "last_search" field.
func LastSearchGTE(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldGTE(FieldLastSearch, v))
}

// LastSearchLT applies the LT predicate on the "last_search" field.
func LastSearchLT(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldLT(FieldLastSearch, v))
}

// LastSearchLTE applies the LTE predicate on the "last_search" field.
func LastSearchLTE(v time.Time) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldLTE(FieldLastSearch, v))
}

// PendingUpdatesEQ applies the EQ predicate on the "pending_updates" field.
func PendingUpdatesEQ(v bool) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldEQ(FieldPendingUpdates, v))
}

// PendingUpdatesNEQ applies the NEQ predicate on the "pending_updates" field.
func PendingUpdatesNEQ(v bool) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.FieldNEQ(FieldPendingUpdates, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.SystemUpdate {
	return predicate.SystemUpdate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.SystemUpdate {
	return predicate.SystemUpdate(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SystemUpdate) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SystemUpdate) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SystemUpdate) predicate.SystemUpdate {
	return predicate.SystemUpdate(sql.NotPredicates(p))
}
