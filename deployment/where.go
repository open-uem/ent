// Code generated by ent, DO NOT EDIT.

package deployment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-uem/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Deployment {
	return predicate.Deployment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Deployment {
	return predicate.Deployment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Deployment {
	return predicate.Deployment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Deployment {
	return predicate.Deployment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Deployment {
	return predicate.Deployment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Deployment {
	return predicate.Deployment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Deployment {
	return predicate.Deployment(sql.FieldLTE(FieldID, id))
}

// PackageID applies equality check predicate on the "package_id" field. It's identical to PackageIDEQ.
func PackageID(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldPackageID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldName, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldVersion, v))
}

// Installed applies equality check predicate on the "installed" field. It's identical to InstalledEQ.
func Installed(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldInstalled, v))
}

// Updated applies equality check predicate on the "updated" field. It's identical to UpdatedEQ.
func Updated(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldUpdated, v))
}

// Failed applies equality check predicate on the "failed" field. It's identical to FailedEQ.
func Failed(v bool) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldFailed, v))
}

// ByProfile applies equality check predicate on the "by_profile" field. It's identical to ByProfileEQ.
func ByProfile(v bool) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldByProfile, v))
}

// PackageIDEQ applies the EQ predicate on the "package_id" field.
func PackageIDEQ(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldPackageID, v))
}

// PackageIDNEQ applies the NEQ predicate on the "package_id" field.
func PackageIDNEQ(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldNEQ(FieldPackageID, v))
}

// PackageIDIn applies the In predicate on the "package_id" field.
func PackageIDIn(vs ...string) predicate.Deployment {
	return predicate.Deployment(sql.FieldIn(FieldPackageID, vs...))
}

// PackageIDNotIn applies the NotIn predicate on the "package_id" field.
func PackageIDNotIn(vs ...string) predicate.Deployment {
	return predicate.Deployment(sql.FieldNotIn(FieldPackageID, vs...))
}

// PackageIDGT applies the GT predicate on the "package_id" field.
func PackageIDGT(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldGT(FieldPackageID, v))
}

// PackageIDGTE applies the GTE predicate on the "package_id" field.
func PackageIDGTE(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldGTE(FieldPackageID, v))
}

// PackageIDLT applies the LT predicate on the "package_id" field.
func PackageIDLT(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldLT(FieldPackageID, v))
}

// PackageIDLTE applies the LTE predicate on the "package_id" field.
func PackageIDLTE(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldLTE(FieldPackageID, v))
}

// PackageIDContains applies the Contains predicate on the "package_id" field.
func PackageIDContains(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldContains(FieldPackageID, v))
}

// PackageIDHasPrefix applies the HasPrefix predicate on the "package_id" field.
func PackageIDHasPrefix(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldHasPrefix(FieldPackageID, v))
}

// PackageIDHasSuffix applies the HasSuffix predicate on the "package_id" field.
func PackageIDHasSuffix(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldHasSuffix(FieldPackageID, v))
}

// PackageIDEqualFold applies the EqualFold predicate on the "package_id" field.
func PackageIDEqualFold(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldEqualFold(FieldPackageID, v))
}

// PackageIDContainsFold applies the ContainsFold predicate on the "package_id" field.
func PackageIDContainsFold(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldContainsFold(FieldPackageID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Deployment {
	return predicate.Deployment(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Deployment {
	return predicate.Deployment(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldContainsFold(FieldName, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.Deployment {
	return predicate.Deployment(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.Deployment {
	return predicate.Deployment(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionIsNil applies the IsNil predicate on the "version" field.
func VersionIsNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldIsNull(FieldVersion))
}

// VersionNotNil applies the NotNil predicate on the "version" field.
func VersionNotNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldNotNull(FieldVersion))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.Deployment {
	return predicate.Deployment(sql.FieldContainsFold(FieldVersion, v))
}

// InstalledEQ applies the EQ predicate on the "installed" field.
func InstalledEQ(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldInstalled, v))
}

// InstalledNEQ applies the NEQ predicate on the "installed" field.
func InstalledNEQ(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldNEQ(FieldInstalled, v))
}

// InstalledIn applies the In predicate on the "installed" field.
func InstalledIn(vs ...time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldIn(FieldInstalled, vs...))
}

// InstalledNotIn applies the NotIn predicate on the "installed" field.
func InstalledNotIn(vs ...time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldNotIn(FieldInstalled, vs...))
}

// InstalledGT applies the GT predicate on the "installed" field.
func InstalledGT(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldGT(FieldInstalled, v))
}

// InstalledGTE applies the GTE predicate on the "installed" field.
func InstalledGTE(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldGTE(FieldInstalled, v))
}

// InstalledLT applies the LT predicate on the "installed" field.
func InstalledLT(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldLT(FieldInstalled, v))
}

// InstalledLTE applies the LTE predicate on the "installed" field.
func InstalledLTE(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldLTE(FieldInstalled, v))
}

// InstalledIsNil applies the IsNil predicate on the "installed" field.
func InstalledIsNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldIsNull(FieldInstalled))
}

// InstalledNotNil applies the NotNil predicate on the "installed" field.
func InstalledNotNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldNotNull(FieldInstalled))
}

// UpdatedEQ applies the EQ predicate on the "updated" field.
func UpdatedEQ(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldUpdated, v))
}

// UpdatedNEQ applies the NEQ predicate on the "updated" field.
func UpdatedNEQ(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldNEQ(FieldUpdated, v))
}

// UpdatedIn applies the In predicate on the "updated" field.
func UpdatedIn(vs ...time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldIn(FieldUpdated, vs...))
}

// UpdatedNotIn applies the NotIn predicate on the "updated" field.
func UpdatedNotIn(vs ...time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldNotIn(FieldUpdated, vs...))
}

// UpdatedGT applies the GT predicate on the "updated" field.
func UpdatedGT(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldGT(FieldUpdated, v))
}

// UpdatedGTE applies the GTE predicate on the "updated" field.
func UpdatedGTE(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldGTE(FieldUpdated, v))
}

// UpdatedLT applies the LT predicate on the "updated" field.
func UpdatedLT(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldLT(FieldUpdated, v))
}

// UpdatedLTE applies the LTE predicate on the "updated" field.
func UpdatedLTE(v time.Time) predicate.Deployment {
	return predicate.Deployment(sql.FieldLTE(FieldUpdated, v))
}

// UpdatedIsNil applies the IsNil predicate on the "updated" field.
func UpdatedIsNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldIsNull(FieldUpdated))
}

// UpdatedNotNil applies the NotNil predicate on the "updated" field.
func UpdatedNotNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldNotNull(FieldUpdated))
}

// FailedEQ applies the EQ predicate on the "failed" field.
func FailedEQ(v bool) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldFailed, v))
}

// FailedNEQ applies the NEQ predicate on the "failed" field.
func FailedNEQ(v bool) predicate.Deployment {
	return predicate.Deployment(sql.FieldNEQ(FieldFailed, v))
}

// FailedIsNil applies the IsNil predicate on the "failed" field.
func FailedIsNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldIsNull(FieldFailed))
}

// FailedNotNil applies the NotNil predicate on the "failed" field.
func FailedNotNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldNotNull(FieldFailed))
}

// ByProfileEQ applies the EQ predicate on the "by_profile" field.
func ByProfileEQ(v bool) predicate.Deployment {
	return predicate.Deployment(sql.FieldEQ(FieldByProfile, v))
}

// ByProfileNEQ applies the NEQ predicate on the "by_profile" field.
func ByProfileNEQ(v bool) predicate.Deployment {
	return predicate.Deployment(sql.FieldNEQ(FieldByProfile, v))
}

// ByProfileIsNil applies the IsNil predicate on the "by_profile" field.
func ByProfileIsNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldIsNull(FieldByProfile))
}

// ByProfileNotNil applies the NotNil predicate on the "by_profile" field.
func ByProfileNotNil() predicate.Deployment {
	return predicate.Deployment(sql.FieldNotNull(FieldByProfile))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Deployment) predicate.Deployment {
	return predicate.Deployment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Deployment) predicate.Deployment {
	return predicate.Deployment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Deployment) predicate.Deployment {
	return predicate.Deployment(sql.NotPredicates(p))
}
