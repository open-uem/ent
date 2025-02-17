// Code generated by ent, DO NOT EDIT.

package operatingsystem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-uem/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldType, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldVersion, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldDescription, v))
}

// Edition applies equality check predicate on the "edition" field. It's identical to EditionEQ.
func Edition(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldEdition, v))
}

// InstallDate applies equality check predicate on the "install_date" field. It's identical to InstallDateEQ.
func InstallDate(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldInstallDate, v))
}

// Arch applies equality check predicate on the "arch" field. It's identical to ArchEQ.
func Arch(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldArch, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldUsername, v))
}

// LastBootupTime applies equality check predicate on the "last_bootup_time" field. It's identical to LastBootupTimeEQ.
func LastBootupTime(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldLastBootupTime, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContainsFold(FieldType, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContainsFold(FieldVersion, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContainsFold(FieldDescription, v))
}

// EditionEQ applies the EQ predicate on the "edition" field.
func EditionEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldEdition, v))
}

// EditionNEQ applies the NEQ predicate on the "edition" field.
func EditionNEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNEQ(FieldEdition, v))
}

// EditionIn applies the In predicate on the "edition" field.
func EditionIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIn(FieldEdition, vs...))
}

// EditionNotIn applies the NotIn predicate on the "edition" field.
func EditionNotIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotIn(FieldEdition, vs...))
}

// EditionGT applies the GT predicate on the "edition" field.
func EditionGT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGT(FieldEdition, v))
}

// EditionGTE applies the GTE predicate on the "edition" field.
func EditionGTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGTE(FieldEdition, v))
}

// EditionLT applies the LT predicate on the "edition" field.
func EditionLT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLT(FieldEdition, v))
}

// EditionLTE applies the LTE predicate on the "edition" field.
func EditionLTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLTE(FieldEdition, v))
}

// EditionContains applies the Contains predicate on the "edition" field.
func EditionContains(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContains(FieldEdition, v))
}

// EditionHasPrefix applies the HasPrefix predicate on the "edition" field.
func EditionHasPrefix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasPrefix(FieldEdition, v))
}

// EditionHasSuffix applies the HasSuffix predicate on the "edition" field.
func EditionHasSuffix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasSuffix(FieldEdition, v))
}

// EditionIsNil applies the IsNil predicate on the "edition" field.
func EditionIsNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIsNull(FieldEdition))
}

// EditionNotNil applies the NotNil predicate on the "edition" field.
func EditionNotNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotNull(FieldEdition))
}

// EditionEqualFold applies the EqualFold predicate on the "edition" field.
func EditionEqualFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEqualFold(FieldEdition, v))
}

// EditionContainsFold applies the ContainsFold predicate on the "edition" field.
func EditionContainsFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContainsFold(FieldEdition, v))
}

// InstallDateEQ applies the EQ predicate on the "install_date" field.
func InstallDateEQ(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldInstallDate, v))
}

// InstallDateNEQ applies the NEQ predicate on the "install_date" field.
func InstallDateNEQ(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNEQ(FieldInstallDate, v))
}

// InstallDateIn applies the In predicate on the "install_date" field.
func InstallDateIn(vs ...time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIn(FieldInstallDate, vs...))
}

// InstallDateNotIn applies the NotIn predicate on the "install_date" field.
func InstallDateNotIn(vs ...time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotIn(FieldInstallDate, vs...))
}

// InstallDateGT applies the GT predicate on the "install_date" field.
func InstallDateGT(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGT(FieldInstallDate, v))
}

// InstallDateGTE applies the GTE predicate on the "install_date" field.
func InstallDateGTE(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGTE(FieldInstallDate, v))
}

// InstallDateLT applies the LT predicate on the "install_date" field.
func InstallDateLT(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLT(FieldInstallDate, v))
}

// InstallDateLTE applies the LTE predicate on the "install_date" field.
func InstallDateLTE(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLTE(FieldInstallDate, v))
}

// InstallDateIsNil applies the IsNil predicate on the "install_date" field.
func InstallDateIsNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIsNull(FieldInstallDate))
}

// InstallDateNotNil applies the NotNil predicate on the "install_date" field.
func InstallDateNotNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotNull(FieldInstallDate))
}

// ArchEQ applies the EQ predicate on the "arch" field.
func ArchEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldArch, v))
}

// ArchNEQ applies the NEQ predicate on the "arch" field.
func ArchNEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNEQ(FieldArch, v))
}

// ArchIn applies the In predicate on the "arch" field.
func ArchIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIn(FieldArch, vs...))
}

// ArchNotIn applies the NotIn predicate on the "arch" field.
func ArchNotIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotIn(FieldArch, vs...))
}

// ArchGT applies the GT predicate on the "arch" field.
func ArchGT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGT(FieldArch, v))
}

// ArchGTE applies the GTE predicate on the "arch" field.
func ArchGTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGTE(FieldArch, v))
}

// ArchLT applies the LT predicate on the "arch" field.
func ArchLT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLT(FieldArch, v))
}

// ArchLTE applies the LTE predicate on the "arch" field.
func ArchLTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLTE(FieldArch, v))
}

// ArchContains applies the Contains predicate on the "arch" field.
func ArchContains(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContains(FieldArch, v))
}

// ArchHasPrefix applies the HasPrefix predicate on the "arch" field.
func ArchHasPrefix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasPrefix(FieldArch, v))
}

// ArchHasSuffix applies the HasSuffix predicate on the "arch" field.
func ArchHasSuffix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasSuffix(FieldArch, v))
}

// ArchIsNil applies the IsNil predicate on the "arch" field.
func ArchIsNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIsNull(FieldArch))
}

// ArchNotNil applies the NotNil predicate on the "arch" field.
func ArchNotNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotNull(FieldArch))
}

// ArchEqualFold applies the EqualFold predicate on the "arch" field.
func ArchEqualFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEqualFold(FieldArch, v))
}

// ArchContainsFold applies the ContainsFold predicate on the "arch" field.
func ArchContainsFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContainsFold(FieldArch, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldContainsFold(FieldUsername, v))
}

// LastBootupTimeEQ applies the EQ predicate on the "last_bootup_time" field.
func LastBootupTimeEQ(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldEQ(FieldLastBootupTime, v))
}

// LastBootupTimeNEQ applies the NEQ predicate on the "last_bootup_time" field.
func LastBootupTimeNEQ(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNEQ(FieldLastBootupTime, v))
}

// LastBootupTimeIn applies the In predicate on the "last_bootup_time" field.
func LastBootupTimeIn(vs ...time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIn(FieldLastBootupTime, vs...))
}

// LastBootupTimeNotIn applies the NotIn predicate on the "last_bootup_time" field.
func LastBootupTimeNotIn(vs ...time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotIn(FieldLastBootupTime, vs...))
}

// LastBootupTimeGT applies the GT predicate on the "last_bootup_time" field.
func LastBootupTimeGT(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGT(FieldLastBootupTime, v))
}

// LastBootupTimeGTE applies the GTE predicate on the "last_bootup_time" field.
func LastBootupTimeGTE(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldGTE(FieldLastBootupTime, v))
}

// LastBootupTimeLT applies the LT predicate on the "last_bootup_time" field.
func LastBootupTimeLT(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLT(FieldLastBootupTime, v))
}

// LastBootupTimeLTE applies the LTE predicate on the "last_bootup_time" field.
func LastBootupTimeLTE(v time.Time) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldLTE(FieldLastBootupTime, v))
}

// LastBootupTimeIsNil applies the IsNil predicate on the "last_bootup_time" field.
func LastBootupTimeIsNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldIsNull(FieldLastBootupTime))
}

// LastBootupTimeNotNil applies the NotNil predicate on the "last_bootup_time" field.
func LastBootupTimeNotNil() predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.FieldNotNull(FieldLastBootupTime))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.OperatingSystem {
	return predicate.OperatingSystem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.OperatingSystem {
	return predicate.OperatingSystem(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OperatingSystem) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OperatingSystem) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OperatingSystem) predicate.OperatingSystem {
	return predicate.OperatingSystem(sql.NotPredicates(p))
}
