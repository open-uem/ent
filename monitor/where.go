// Code generated by ent, DO NOT EDIT.

package monitor

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/doncicuto/openuem_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Monitor {
	return predicate.Monitor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Monitor {
	return predicate.Monitor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Monitor {
	return predicate.Monitor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Monitor {
	return predicate.Monitor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Monitor {
	return predicate.Monitor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Monitor {
	return predicate.Monitor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Monitor {
	return predicate.Monitor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Monitor {
	return predicate.Monitor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Monitor {
	return predicate.Monitor(sql.FieldLTE(FieldID, id))
}

// Manufacturer applies equality check predicate on the "manufacturer" field. It's identical to ManufacturerEQ.
func Manufacturer(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldEQ(FieldManufacturer, v))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldEQ(FieldModel, v))
}

// Serial applies equality check predicate on the "serial" field. It's identical to SerialEQ.
func Serial(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldEQ(FieldSerial, v))
}

// ManufacturerEQ applies the EQ predicate on the "manufacturer" field.
func ManufacturerEQ(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldEQ(FieldManufacturer, v))
}

// ManufacturerNEQ applies the NEQ predicate on the "manufacturer" field.
func ManufacturerNEQ(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldNEQ(FieldManufacturer, v))
}

// ManufacturerIn applies the In predicate on the "manufacturer" field.
func ManufacturerIn(vs ...string) predicate.Monitor {
	return predicate.Monitor(sql.FieldIn(FieldManufacturer, vs...))
}

// ManufacturerNotIn applies the NotIn predicate on the "manufacturer" field.
func ManufacturerNotIn(vs ...string) predicate.Monitor {
	return predicate.Monitor(sql.FieldNotIn(FieldManufacturer, vs...))
}

// ManufacturerGT applies the GT predicate on the "manufacturer" field.
func ManufacturerGT(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldGT(FieldManufacturer, v))
}

// ManufacturerGTE applies the GTE predicate on the "manufacturer" field.
func ManufacturerGTE(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldGTE(FieldManufacturer, v))
}

// ManufacturerLT applies the LT predicate on the "manufacturer" field.
func ManufacturerLT(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldLT(FieldManufacturer, v))
}

// ManufacturerLTE applies the LTE predicate on the "manufacturer" field.
func ManufacturerLTE(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldLTE(FieldManufacturer, v))
}

// ManufacturerContains applies the Contains predicate on the "manufacturer" field.
func ManufacturerContains(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldContains(FieldManufacturer, v))
}

// ManufacturerHasPrefix applies the HasPrefix predicate on the "manufacturer" field.
func ManufacturerHasPrefix(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldHasPrefix(FieldManufacturer, v))
}

// ManufacturerHasSuffix applies the HasSuffix predicate on the "manufacturer" field.
func ManufacturerHasSuffix(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldHasSuffix(FieldManufacturer, v))
}

// ManufacturerIsNil applies the IsNil predicate on the "manufacturer" field.
func ManufacturerIsNil() predicate.Monitor {
	return predicate.Monitor(sql.FieldIsNull(FieldManufacturer))
}

// ManufacturerNotNil applies the NotNil predicate on the "manufacturer" field.
func ManufacturerNotNil() predicate.Monitor {
	return predicate.Monitor(sql.FieldNotNull(FieldManufacturer))
}

// ManufacturerEqualFold applies the EqualFold predicate on the "manufacturer" field.
func ManufacturerEqualFold(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldEqualFold(FieldManufacturer, v))
}

// ManufacturerContainsFold applies the ContainsFold predicate on the "manufacturer" field.
func ManufacturerContainsFold(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldContainsFold(FieldManufacturer, v))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Monitor {
	return predicate.Monitor(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Monitor {
	return predicate.Monitor(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldLTE(FieldModel, v))
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldContains(FieldModel, v))
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldHasPrefix(FieldModel, v))
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldHasSuffix(FieldModel, v))
}

// ModelIsNil applies the IsNil predicate on the "model" field.
func ModelIsNil() predicate.Monitor {
	return predicate.Monitor(sql.FieldIsNull(FieldModel))
}

// ModelNotNil applies the NotNil predicate on the "model" field.
func ModelNotNil() predicate.Monitor {
	return predicate.Monitor(sql.FieldNotNull(FieldModel))
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldEqualFold(FieldModel, v))
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldContainsFold(FieldModel, v))
}

// SerialEQ applies the EQ predicate on the "serial" field.
func SerialEQ(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldEQ(FieldSerial, v))
}

// SerialNEQ applies the NEQ predicate on the "serial" field.
func SerialNEQ(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldNEQ(FieldSerial, v))
}

// SerialIn applies the In predicate on the "serial" field.
func SerialIn(vs ...string) predicate.Monitor {
	return predicate.Monitor(sql.FieldIn(FieldSerial, vs...))
}

// SerialNotIn applies the NotIn predicate on the "serial" field.
func SerialNotIn(vs ...string) predicate.Monitor {
	return predicate.Monitor(sql.FieldNotIn(FieldSerial, vs...))
}

// SerialGT applies the GT predicate on the "serial" field.
func SerialGT(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldGT(FieldSerial, v))
}

// SerialGTE applies the GTE predicate on the "serial" field.
func SerialGTE(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldGTE(FieldSerial, v))
}

// SerialLT applies the LT predicate on the "serial" field.
func SerialLT(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldLT(FieldSerial, v))
}

// SerialLTE applies the LTE predicate on the "serial" field.
func SerialLTE(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldLTE(FieldSerial, v))
}

// SerialContains applies the Contains predicate on the "serial" field.
func SerialContains(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldContains(FieldSerial, v))
}

// SerialHasPrefix applies the HasPrefix predicate on the "serial" field.
func SerialHasPrefix(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldHasPrefix(FieldSerial, v))
}

// SerialHasSuffix applies the HasSuffix predicate on the "serial" field.
func SerialHasSuffix(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldHasSuffix(FieldSerial, v))
}

// SerialIsNil applies the IsNil predicate on the "serial" field.
func SerialIsNil() predicate.Monitor {
	return predicate.Monitor(sql.FieldIsNull(FieldSerial))
}

// SerialNotNil applies the NotNil predicate on the "serial" field.
func SerialNotNil() predicate.Monitor {
	return predicate.Monitor(sql.FieldNotNull(FieldSerial))
}

// SerialEqualFold applies the EqualFold predicate on the "serial" field.
func SerialEqualFold(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldEqualFold(FieldSerial, v))
}

// SerialContainsFold applies the ContainsFold predicate on the "serial" field.
func SerialContainsFold(v string) predicate.Monitor {
	return predicate.Monitor(sql.FieldContainsFold(FieldSerial, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.Monitor {
	return predicate.Monitor(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Monitor) predicate.Monitor {
	return predicate.Monitor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Monitor) predicate.Monitor {
	return predicate.Monitor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Monitor) predicate.Monitor {
	return predicate.Monitor(sql.NotPredicates(p))
}
