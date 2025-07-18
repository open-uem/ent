// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-uem/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCountry, v))
}

// EmailVerified applies equality check predicate on the "email_verified" field. It's identical to EmailVerifiedEQ.
func EmailVerified(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailVerified, v))
}

// Register applies equality check predicate on the "register" field. It's identical to RegisterEQ.
func Register(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRegister, v))
}

// CertClearPassword applies equality check predicate on the "cert_clear_password" field. It's identical to CertClearPasswordEQ.
func CertClearPassword(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCertClearPassword, v))
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldExpiry, v))
}

// Openid applies equality check predicate on the "openid" field. It's identical to OpenidEQ.
func Openid(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOpenid, v))
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreated, v))
}

// Modified applies equality check predicate on the "modified" field. It's identical to ModifiedEQ.
func Modified(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldModified, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhone, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryIsNil applies the IsNil predicate on the "country" field.
func CountryIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCountry))
}

// CountryNotNil applies the NotNil predicate on the "country" field.
func CountryNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCountry))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldCountry, v))
}

// EmailVerifiedEQ applies the EQ predicate on the "email_verified" field.
func EmailVerifiedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailVerified, v))
}

// EmailVerifiedNEQ applies the NEQ predicate on the "email_verified" field.
func EmailVerifiedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmailVerified, v))
}

// RegisterEQ applies the EQ predicate on the "register" field.
func RegisterEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRegister, v))
}

// RegisterNEQ applies the NEQ predicate on the "register" field.
func RegisterNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRegister, v))
}

// RegisterIn applies the In predicate on the "register" field.
func RegisterIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldRegister, vs...))
}

// RegisterNotIn applies the NotIn predicate on the "register" field.
func RegisterNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRegister, vs...))
}

// RegisterGT applies the GT predicate on the "register" field.
func RegisterGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldRegister, v))
}

// RegisterGTE applies the GTE predicate on the "register" field.
func RegisterGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRegister, v))
}

// RegisterLT applies the LT predicate on the "register" field.
func RegisterLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldRegister, v))
}

// RegisterLTE applies the LTE predicate on the "register" field.
func RegisterLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRegister, v))
}

// RegisterContains applies the Contains predicate on the "register" field.
func RegisterContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldRegister, v))
}

// RegisterHasPrefix applies the HasPrefix predicate on the "register" field.
func RegisterHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldRegister, v))
}

// RegisterHasSuffix applies the HasSuffix predicate on the "register" field.
func RegisterHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldRegister, v))
}

// RegisterEqualFold applies the EqualFold predicate on the "register" field.
func RegisterEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldRegister, v))
}

// RegisterContainsFold applies the ContainsFold predicate on the "register" field.
func RegisterContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldRegister, v))
}

// CertClearPasswordEQ applies the EQ predicate on the "cert_clear_password" field.
func CertClearPasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCertClearPassword, v))
}

// CertClearPasswordNEQ applies the NEQ predicate on the "cert_clear_password" field.
func CertClearPasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCertClearPassword, v))
}

// CertClearPasswordIn applies the In predicate on the "cert_clear_password" field.
func CertClearPasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldCertClearPassword, vs...))
}

// CertClearPasswordNotIn applies the NotIn predicate on the "cert_clear_password" field.
func CertClearPasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCertClearPassword, vs...))
}

// CertClearPasswordGT applies the GT predicate on the "cert_clear_password" field.
func CertClearPasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldCertClearPassword, v))
}

// CertClearPasswordGTE applies the GTE predicate on the "cert_clear_password" field.
func CertClearPasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCertClearPassword, v))
}

// CertClearPasswordLT applies the LT predicate on the "cert_clear_password" field.
func CertClearPasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldCertClearPassword, v))
}

// CertClearPasswordLTE applies the LTE predicate on the "cert_clear_password" field.
func CertClearPasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCertClearPassword, v))
}

// CertClearPasswordContains applies the Contains predicate on the "cert_clear_password" field.
func CertClearPasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldCertClearPassword, v))
}

// CertClearPasswordHasPrefix applies the HasPrefix predicate on the "cert_clear_password" field.
func CertClearPasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldCertClearPassword, v))
}

// CertClearPasswordHasSuffix applies the HasSuffix predicate on the "cert_clear_password" field.
func CertClearPasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldCertClearPassword, v))
}

// CertClearPasswordIsNil applies the IsNil predicate on the "cert_clear_password" field.
func CertClearPasswordIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCertClearPassword))
}

// CertClearPasswordNotNil applies the NotNil predicate on the "cert_clear_password" field.
func CertClearPasswordNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCertClearPassword))
}

// CertClearPasswordEqualFold applies the EqualFold predicate on the "cert_clear_password" field.
func CertClearPasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldCertClearPassword, v))
}

// CertClearPasswordContainsFold applies the ContainsFold predicate on the "cert_clear_password" field.
func CertClearPasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldCertClearPassword, v))
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldExpiry, v))
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldExpiry, v))
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldExpiry, vs...))
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldExpiry, vs...))
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldExpiry, v))
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldExpiry, v))
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldExpiry, v))
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldExpiry, v))
}

// ExpiryIsNil applies the IsNil predicate on the "expiry" field.
func ExpiryIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldExpiry))
}

// ExpiryNotNil applies the NotNil predicate on the "expiry" field.
func ExpiryNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldExpiry))
}

// OpenidEQ applies the EQ predicate on the "openid" field.
func OpenidEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOpenid, v))
}

// OpenidNEQ applies the NEQ predicate on the "openid" field.
func OpenidNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOpenid, v))
}

// OpenidIsNil applies the IsNil predicate on the "openid" field.
func OpenidIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldOpenid))
}

// OpenidNotNil applies the NotNil predicate on the "openid" field.
func OpenidNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldOpenid))
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreated, v))
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreated, v))
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreated, vs...))
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreated, vs...))
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreated, v))
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreated, v))
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreated, v))
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreated, v))
}

// CreatedIsNil applies the IsNil predicate on the "created" field.
func CreatedIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCreated))
}

// CreatedNotNil applies the NotNil predicate on the "created" field.
func CreatedNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCreated))
}

// ModifiedEQ applies the EQ predicate on the "modified" field.
func ModifiedEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldModified, v))
}

// ModifiedNEQ applies the NEQ predicate on the "modified" field.
func ModifiedNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldModified, v))
}

// ModifiedIn applies the In predicate on the "modified" field.
func ModifiedIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldModified, vs...))
}

// ModifiedNotIn applies the NotIn predicate on the "modified" field.
func ModifiedNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldModified, vs...))
}

// ModifiedGT applies the GT predicate on the "modified" field.
func ModifiedGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldModified, v))
}

// ModifiedGTE applies the GTE predicate on the "modified" field.
func ModifiedGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldModified, v))
}

// ModifiedLT applies the LT predicate on the "modified" field.
func ModifiedLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldModified, v))
}

// ModifiedLTE applies the LTE predicate on the "modified" field.
func ModifiedLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldModified, v))
}

// ModifiedIsNil applies the IsNil predicate on the "modified" field.
func ModifiedIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldModified))
}

// ModifiedNotNil applies the NotNil predicate on the "modified" field.
func ModifiedNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldModified))
}

// HasSessions applies the HasEdge predicate on the "sessions" edge.
func HasSessions() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSessionsWith applies the HasEdge predicate on the "sessions" edge with a given conditions (other predicates).
func HasSessionsWith(preds ...predicate.Sessions) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newSessionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
