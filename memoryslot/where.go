// Code generated by ent, DO NOT EDIT.

package memoryslot

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-uem/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLTE(FieldID, id))
}

// Slot applies equality check predicate on the "slot" field. It's identical to SlotEQ.
func Slot(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldSlot, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldSize, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldType, v))
}

// SerialNumber applies equality check predicate on the "serial_number" field. It's identical to SerialNumberEQ.
func SerialNumber(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldSerialNumber, v))
}

// PartNumber applies equality check predicate on the "part_number" field. It's identical to PartNumberEQ.
func PartNumber(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldPartNumber, v))
}

// Speed applies equality check predicate on the "speed" field. It's identical to SpeedEQ.
func Speed(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldSpeed, v))
}

// Manufacturer applies equality check predicate on the "manufacturer" field. It's identical to ManufacturerEQ.
func Manufacturer(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldManufacturer, v))
}

// SlotEQ applies the EQ predicate on the "slot" field.
func SlotEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldSlot, v))
}

// SlotNEQ applies the NEQ predicate on the "slot" field.
func SlotNEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNEQ(FieldSlot, v))
}

// SlotIn applies the In predicate on the "slot" field.
func SlotIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIn(FieldSlot, vs...))
}

// SlotNotIn applies the NotIn predicate on the "slot" field.
func SlotNotIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotIn(FieldSlot, vs...))
}

// SlotGT applies the GT predicate on the "slot" field.
func SlotGT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGT(FieldSlot, v))
}

// SlotGTE applies the GTE predicate on the "slot" field.
func SlotGTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGTE(FieldSlot, v))
}

// SlotLT applies the LT predicate on the "slot" field.
func SlotLT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLT(FieldSlot, v))
}

// SlotLTE applies the LTE predicate on the "slot" field.
func SlotLTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLTE(FieldSlot, v))
}

// SlotContains applies the Contains predicate on the "slot" field.
func SlotContains(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContains(FieldSlot, v))
}

// SlotHasPrefix applies the HasPrefix predicate on the "slot" field.
func SlotHasPrefix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasPrefix(FieldSlot, v))
}

// SlotHasSuffix applies the HasSuffix predicate on the "slot" field.
func SlotHasSuffix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasSuffix(FieldSlot, v))
}

// SlotIsNil applies the IsNil predicate on the "slot" field.
func SlotIsNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIsNull(FieldSlot))
}

// SlotNotNil applies the NotNil predicate on the "slot" field.
func SlotNotNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotNull(FieldSlot))
}

// SlotEqualFold applies the EqualFold predicate on the "slot" field.
func SlotEqualFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEqualFold(FieldSlot, v))
}

// SlotContainsFold applies the ContainsFold predicate on the "slot" field.
func SlotContainsFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContainsFold(FieldSlot, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLTE(FieldSize, v))
}

// SizeContains applies the Contains predicate on the "size" field.
func SizeContains(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContains(FieldSize, v))
}

// SizeHasPrefix applies the HasPrefix predicate on the "size" field.
func SizeHasPrefix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasPrefix(FieldSize, v))
}

// SizeHasSuffix applies the HasSuffix predicate on the "size" field.
func SizeHasSuffix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasSuffix(FieldSize, v))
}

// SizeIsNil applies the IsNil predicate on the "size" field.
func SizeIsNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIsNull(FieldSize))
}

// SizeNotNil applies the NotNil predicate on the "size" field.
func SizeNotNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotNull(FieldSize))
}

// SizeEqualFold applies the EqualFold predicate on the "size" field.
func SizeEqualFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEqualFold(FieldSize, v))
}

// SizeContainsFold applies the ContainsFold predicate on the "size" field.
func SizeContainsFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContainsFold(FieldSize, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContainsFold(FieldType, v))
}

// SerialNumberEQ applies the EQ predicate on the "serial_number" field.
func SerialNumberEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldSerialNumber, v))
}

// SerialNumberNEQ applies the NEQ predicate on the "serial_number" field.
func SerialNumberNEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNEQ(FieldSerialNumber, v))
}

// SerialNumberIn applies the In predicate on the "serial_number" field.
func SerialNumberIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIn(FieldSerialNumber, vs...))
}

// SerialNumberNotIn applies the NotIn predicate on the "serial_number" field.
func SerialNumberNotIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotIn(FieldSerialNumber, vs...))
}

// SerialNumberGT applies the GT predicate on the "serial_number" field.
func SerialNumberGT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGT(FieldSerialNumber, v))
}

// SerialNumberGTE applies the GTE predicate on the "serial_number" field.
func SerialNumberGTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGTE(FieldSerialNumber, v))
}

// SerialNumberLT applies the LT predicate on the "serial_number" field.
func SerialNumberLT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLT(FieldSerialNumber, v))
}

// SerialNumberLTE applies the LTE predicate on the "serial_number" field.
func SerialNumberLTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLTE(FieldSerialNumber, v))
}

// SerialNumberContains applies the Contains predicate on the "serial_number" field.
func SerialNumberContains(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContains(FieldSerialNumber, v))
}

// SerialNumberHasPrefix applies the HasPrefix predicate on the "serial_number" field.
func SerialNumberHasPrefix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasPrefix(FieldSerialNumber, v))
}

// SerialNumberHasSuffix applies the HasSuffix predicate on the "serial_number" field.
func SerialNumberHasSuffix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasSuffix(FieldSerialNumber, v))
}

// SerialNumberIsNil applies the IsNil predicate on the "serial_number" field.
func SerialNumberIsNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIsNull(FieldSerialNumber))
}

// SerialNumberNotNil applies the NotNil predicate on the "serial_number" field.
func SerialNumberNotNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotNull(FieldSerialNumber))
}

// SerialNumberEqualFold applies the EqualFold predicate on the "serial_number" field.
func SerialNumberEqualFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEqualFold(FieldSerialNumber, v))
}

// SerialNumberContainsFold applies the ContainsFold predicate on the "serial_number" field.
func SerialNumberContainsFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContainsFold(FieldSerialNumber, v))
}

// PartNumberEQ applies the EQ predicate on the "part_number" field.
func PartNumberEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldPartNumber, v))
}

// PartNumberNEQ applies the NEQ predicate on the "part_number" field.
func PartNumberNEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNEQ(FieldPartNumber, v))
}

// PartNumberIn applies the In predicate on the "part_number" field.
func PartNumberIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIn(FieldPartNumber, vs...))
}

// PartNumberNotIn applies the NotIn predicate on the "part_number" field.
func PartNumberNotIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotIn(FieldPartNumber, vs...))
}

// PartNumberGT applies the GT predicate on the "part_number" field.
func PartNumberGT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGT(FieldPartNumber, v))
}

// PartNumberGTE applies the GTE predicate on the "part_number" field.
func PartNumberGTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGTE(FieldPartNumber, v))
}

// PartNumberLT applies the LT predicate on the "part_number" field.
func PartNumberLT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLT(FieldPartNumber, v))
}

// PartNumberLTE applies the LTE predicate on the "part_number" field.
func PartNumberLTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLTE(FieldPartNumber, v))
}

// PartNumberContains applies the Contains predicate on the "part_number" field.
func PartNumberContains(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContains(FieldPartNumber, v))
}

// PartNumberHasPrefix applies the HasPrefix predicate on the "part_number" field.
func PartNumberHasPrefix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasPrefix(FieldPartNumber, v))
}

// PartNumberHasSuffix applies the HasSuffix predicate on the "part_number" field.
func PartNumberHasSuffix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasSuffix(FieldPartNumber, v))
}

// PartNumberIsNil applies the IsNil predicate on the "part_number" field.
func PartNumberIsNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIsNull(FieldPartNumber))
}

// PartNumberNotNil applies the NotNil predicate on the "part_number" field.
func PartNumberNotNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotNull(FieldPartNumber))
}

// PartNumberEqualFold applies the EqualFold predicate on the "part_number" field.
func PartNumberEqualFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEqualFold(FieldPartNumber, v))
}

// PartNumberContainsFold applies the ContainsFold predicate on the "part_number" field.
func PartNumberContainsFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContainsFold(FieldPartNumber, v))
}

// SpeedEQ applies the EQ predicate on the "speed" field.
func SpeedEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldSpeed, v))
}

// SpeedNEQ applies the NEQ predicate on the "speed" field.
func SpeedNEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNEQ(FieldSpeed, v))
}

// SpeedIn applies the In predicate on the "speed" field.
func SpeedIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIn(FieldSpeed, vs...))
}

// SpeedNotIn applies the NotIn predicate on the "speed" field.
func SpeedNotIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotIn(FieldSpeed, vs...))
}

// SpeedGT applies the GT predicate on the "speed" field.
func SpeedGT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGT(FieldSpeed, v))
}

// SpeedGTE applies the GTE predicate on the "speed" field.
func SpeedGTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGTE(FieldSpeed, v))
}

// SpeedLT applies the LT predicate on the "speed" field.
func SpeedLT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLT(FieldSpeed, v))
}

// SpeedLTE applies the LTE predicate on the "speed" field.
func SpeedLTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLTE(FieldSpeed, v))
}

// SpeedContains applies the Contains predicate on the "speed" field.
func SpeedContains(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContains(FieldSpeed, v))
}

// SpeedHasPrefix applies the HasPrefix predicate on the "speed" field.
func SpeedHasPrefix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasPrefix(FieldSpeed, v))
}

// SpeedHasSuffix applies the HasSuffix predicate on the "speed" field.
func SpeedHasSuffix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasSuffix(FieldSpeed, v))
}

// SpeedIsNil applies the IsNil predicate on the "speed" field.
func SpeedIsNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIsNull(FieldSpeed))
}

// SpeedNotNil applies the NotNil predicate on the "speed" field.
func SpeedNotNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotNull(FieldSpeed))
}

// SpeedEqualFold applies the EqualFold predicate on the "speed" field.
func SpeedEqualFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEqualFold(FieldSpeed, v))
}

// SpeedContainsFold applies the ContainsFold predicate on the "speed" field.
func SpeedContainsFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContainsFold(FieldSpeed, v))
}

// ManufacturerEQ applies the EQ predicate on the "manufacturer" field.
func ManufacturerEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEQ(FieldManufacturer, v))
}

// ManufacturerNEQ applies the NEQ predicate on the "manufacturer" field.
func ManufacturerNEQ(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNEQ(FieldManufacturer, v))
}

// ManufacturerIn applies the In predicate on the "manufacturer" field.
func ManufacturerIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIn(FieldManufacturer, vs...))
}

// ManufacturerNotIn applies the NotIn predicate on the "manufacturer" field.
func ManufacturerNotIn(vs ...string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotIn(FieldManufacturer, vs...))
}

// ManufacturerGT applies the GT predicate on the "manufacturer" field.
func ManufacturerGT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGT(FieldManufacturer, v))
}

// ManufacturerGTE applies the GTE predicate on the "manufacturer" field.
func ManufacturerGTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldGTE(FieldManufacturer, v))
}

// ManufacturerLT applies the LT predicate on the "manufacturer" field.
func ManufacturerLT(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLT(FieldManufacturer, v))
}

// ManufacturerLTE applies the LTE predicate on the "manufacturer" field.
func ManufacturerLTE(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldLTE(FieldManufacturer, v))
}

// ManufacturerContains applies the Contains predicate on the "manufacturer" field.
func ManufacturerContains(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContains(FieldManufacturer, v))
}

// ManufacturerHasPrefix applies the HasPrefix predicate on the "manufacturer" field.
func ManufacturerHasPrefix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasPrefix(FieldManufacturer, v))
}

// ManufacturerHasSuffix applies the HasSuffix predicate on the "manufacturer" field.
func ManufacturerHasSuffix(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldHasSuffix(FieldManufacturer, v))
}

// ManufacturerIsNil applies the IsNil predicate on the "manufacturer" field.
func ManufacturerIsNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldIsNull(FieldManufacturer))
}

// ManufacturerNotNil applies the NotNil predicate on the "manufacturer" field.
func ManufacturerNotNil() predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldNotNull(FieldManufacturer))
}

// ManufacturerEqualFold applies the EqualFold predicate on the "manufacturer" field.
func ManufacturerEqualFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldEqualFold(FieldManufacturer, v))
}

// ManufacturerContainsFold applies the ContainsFold predicate on the "manufacturer" field.
func ManufacturerContainsFold(v string) predicate.MemorySlot {
	return predicate.MemorySlot(sql.FieldContainsFold(FieldManufacturer, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.MemorySlot {
	return predicate.MemorySlot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.MemorySlot {
	return predicate.MemorySlot(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MemorySlot) predicate.MemorySlot {
	return predicate.MemorySlot(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MemorySlot) predicate.MemorySlot {
	return predicate.MemorySlot(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MemorySlot) predicate.MemorySlot {
	return predicate.MemorySlot(sql.NotPredicates(p))
}
