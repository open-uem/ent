// Code generated by ent, DO NOT EDIT.

package release

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/doncicuto/openuem_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldID, id))
}

// Channel applies equality check predicate on the "channel" field. It's identical to ChannelEQ.
func Channel(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldChannel, v))
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldSummary, v))
}

// ReleaseNotes applies equality check predicate on the "release_notes" field. It's identical to ReleaseNotesEQ.
func ReleaseNotes(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldReleaseNotes, v))
}

// FileURL applies equality check predicate on the "file_url" field. It's identical to FileURLEQ.
func FileURL(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldFileURL, v))
}

// Checksum applies equality check predicate on the "checksum" field. It's identical to ChecksumEQ.
func Checksum(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldChecksum, v))
}

// IsCritical applies equality check predicate on the "is_critical" field. It's identical to IsCriticalEQ.
func IsCritical(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldIsCritical, v))
}

// ChannelEQ applies the EQ predicate on the "channel" field.
func ChannelEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldChannel, v))
}

// ChannelNEQ applies the NEQ predicate on the "channel" field.
func ChannelNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldChannel, v))
}

// ChannelIn applies the In predicate on the "channel" field.
func ChannelIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldChannel, vs...))
}

// ChannelNotIn applies the NotIn predicate on the "channel" field.
func ChannelNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldChannel, vs...))
}

// ChannelGT applies the GT predicate on the "channel" field.
func ChannelGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldChannel, v))
}

// ChannelGTE applies the GTE predicate on the "channel" field.
func ChannelGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldChannel, v))
}

// ChannelLT applies the LT predicate on the "channel" field.
func ChannelLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldChannel, v))
}

// ChannelLTE applies the LTE predicate on the "channel" field.
func ChannelLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldChannel, v))
}

// ChannelContains applies the Contains predicate on the "channel" field.
func ChannelContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldChannel, v))
}

// ChannelHasPrefix applies the HasPrefix predicate on the "channel" field.
func ChannelHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldChannel, v))
}

// ChannelHasSuffix applies the HasSuffix predicate on the "channel" field.
func ChannelHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldChannel, v))
}

// ChannelIsNil applies the IsNil predicate on the "channel" field.
func ChannelIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldChannel))
}

// ChannelNotNil applies the NotNil predicate on the "channel" field.
func ChannelNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldChannel))
}

// ChannelEqualFold applies the EqualFold predicate on the "channel" field.
func ChannelEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldChannel, v))
}

// ChannelContainsFold applies the ContainsFold predicate on the "channel" field.
func ChannelContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldChannel, v))
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldSummary, v))
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldSummary, v))
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldSummary, vs...))
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldSummary, vs...))
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldSummary, v))
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldSummary, v))
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldSummary, v))
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldSummary, v))
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldSummary, v))
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldSummary, v))
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldSummary, v))
}

// SummaryIsNil applies the IsNil predicate on the "summary" field.
func SummaryIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldSummary))
}

// SummaryNotNil applies the NotNil predicate on the "summary" field.
func SummaryNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldSummary))
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldSummary, v))
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldSummary, v))
}

// ReleaseNotesEQ applies the EQ predicate on the "release_notes" field.
func ReleaseNotesEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldReleaseNotes, v))
}

// ReleaseNotesNEQ applies the NEQ predicate on the "release_notes" field.
func ReleaseNotesNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldReleaseNotes, v))
}

// ReleaseNotesIn applies the In predicate on the "release_notes" field.
func ReleaseNotesIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldReleaseNotes, vs...))
}

// ReleaseNotesNotIn applies the NotIn predicate on the "release_notes" field.
func ReleaseNotesNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldReleaseNotes, vs...))
}

// ReleaseNotesGT applies the GT predicate on the "release_notes" field.
func ReleaseNotesGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldReleaseNotes, v))
}

// ReleaseNotesGTE applies the GTE predicate on the "release_notes" field.
func ReleaseNotesGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldReleaseNotes, v))
}

// ReleaseNotesLT applies the LT predicate on the "release_notes" field.
func ReleaseNotesLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldReleaseNotes, v))
}

// ReleaseNotesLTE applies the LTE predicate on the "release_notes" field.
func ReleaseNotesLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldReleaseNotes, v))
}

// ReleaseNotesContains applies the Contains predicate on the "release_notes" field.
func ReleaseNotesContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldReleaseNotes, v))
}

// ReleaseNotesHasPrefix applies the HasPrefix predicate on the "release_notes" field.
func ReleaseNotesHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldReleaseNotes, v))
}

// ReleaseNotesHasSuffix applies the HasSuffix predicate on the "release_notes" field.
func ReleaseNotesHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldReleaseNotes, v))
}

// ReleaseNotesIsNil applies the IsNil predicate on the "release_notes" field.
func ReleaseNotesIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldReleaseNotes))
}

// ReleaseNotesNotNil applies the NotNil predicate on the "release_notes" field.
func ReleaseNotesNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldReleaseNotes))
}

// ReleaseNotesEqualFold applies the EqualFold predicate on the "release_notes" field.
func ReleaseNotesEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldReleaseNotes, v))
}

// ReleaseNotesContainsFold applies the ContainsFold predicate on the "release_notes" field.
func ReleaseNotesContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldReleaseNotes, v))
}

// FileURLEQ applies the EQ predicate on the "file_url" field.
func FileURLEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldFileURL, v))
}

// FileURLNEQ applies the NEQ predicate on the "file_url" field.
func FileURLNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldFileURL, v))
}

// FileURLIn applies the In predicate on the "file_url" field.
func FileURLIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldFileURL, vs...))
}

// FileURLNotIn applies the NotIn predicate on the "file_url" field.
func FileURLNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldFileURL, vs...))
}

// FileURLGT applies the GT predicate on the "file_url" field.
func FileURLGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldFileURL, v))
}

// FileURLGTE applies the GTE predicate on the "file_url" field.
func FileURLGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldFileURL, v))
}

// FileURLLT applies the LT predicate on the "file_url" field.
func FileURLLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldFileURL, v))
}

// FileURLLTE applies the LTE predicate on the "file_url" field.
func FileURLLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldFileURL, v))
}

// FileURLContains applies the Contains predicate on the "file_url" field.
func FileURLContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldFileURL, v))
}

// FileURLHasPrefix applies the HasPrefix predicate on the "file_url" field.
func FileURLHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldFileURL, v))
}

// FileURLHasSuffix applies the HasSuffix predicate on the "file_url" field.
func FileURLHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldFileURL, v))
}

// FileURLIsNil applies the IsNil predicate on the "file_url" field.
func FileURLIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldFileURL))
}

// FileURLNotNil applies the NotNil predicate on the "file_url" field.
func FileURLNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldFileURL))
}

// FileURLEqualFold applies the EqualFold predicate on the "file_url" field.
func FileURLEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldFileURL, v))
}

// FileURLContainsFold applies the ContainsFold predicate on the "file_url" field.
func FileURLContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldFileURL, v))
}

// ChecksumEQ applies the EQ predicate on the "checksum" field.
func ChecksumEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldChecksum, v))
}

// ChecksumNEQ applies the NEQ predicate on the "checksum" field.
func ChecksumNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldChecksum, v))
}

// ChecksumIn applies the In predicate on the "checksum" field.
func ChecksumIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldChecksum, vs...))
}

// ChecksumNotIn applies the NotIn predicate on the "checksum" field.
func ChecksumNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldChecksum, vs...))
}

// ChecksumGT applies the GT predicate on the "checksum" field.
func ChecksumGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldChecksum, v))
}

// ChecksumGTE applies the GTE predicate on the "checksum" field.
func ChecksumGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldChecksum, v))
}

// ChecksumLT applies the LT predicate on the "checksum" field.
func ChecksumLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldChecksum, v))
}

// ChecksumLTE applies the LTE predicate on the "checksum" field.
func ChecksumLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldChecksum, v))
}

// ChecksumContains applies the Contains predicate on the "checksum" field.
func ChecksumContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldChecksum, v))
}

// ChecksumHasPrefix applies the HasPrefix predicate on the "checksum" field.
func ChecksumHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldChecksum, v))
}

// ChecksumHasSuffix applies the HasSuffix predicate on the "checksum" field.
func ChecksumHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldChecksum, v))
}

// ChecksumIsNil applies the IsNil predicate on the "checksum" field.
func ChecksumIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldChecksum))
}

// ChecksumNotNil applies the NotNil predicate on the "checksum" field.
func ChecksumNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldChecksum))
}

// ChecksumEqualFold applies the EqualFold predicate on the "checksum" field.
func ChecksumEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldChecksum, v))
}

// ChecksumContainsFold applies the ContainsFold predicate on the "checksum" field.
func ChecksumContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldChecksum, v))
}

// IsCriticalEQ applies the EQ predicate on the "is_critical" field.
func IsCriticalEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldIsCritical, v))
}

// IsCriticalNEQ applies the NEQ predicate on the "is_critical" field.
func IsCriticalNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldIsCritical, v))
}

// IsCriticalIn applies the In predicate on the "is_critical" field.
func IsCriticalIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldIsCritical, vs...))
}

// IsCriticalNotIn applies the NotIn predicate on the "is_critical" field.
func IsCriticalNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldIsCritical, vs...))
}

// IsCriticalGT applies the GT predicate on the "is_critical" field.
func IsCriticalGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldIsCritical, v))
}

// IsCriticalGTE applies the GTE predicate on the "is_critical" field.
func IsCriticalGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldIsCritical, v))
}

// IsCriticalLT applies the LT predicate on the "is_critical" field.
func IsCriticalLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldIsCritical, v))
}

// IsCriticalLTE applies the LTE predicate on the "is_critical" field.
func IsCriticalLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldIsCritical, v))
}

// IsCriticalContains applies the Contains predicate on the "is_critical" field.
func IsCriticalContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldIsCritical, v))
}

// IsCriticalHasPrefix applies the HasPrefix predicate on the "is_critical" field.
func IsCriticalHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldIsCritical, v))
}

// IsCriticalHasSuffix applies the HasSuffix predicate on the "is_critical" field.
func IsCriticalHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldIsCritical, v))
}

// IsCriticalIsNil applies the IsNil predicate on the "is_critical" field.
func IsCriticalIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldIsCritical))
}

// IsCriticalNotNil applies the NotNil predicate on the "is_critical" field.
func IsCriticalNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldIsCritical))
}

// IsCriticalEqualFold applies the EqualFold predicate on the "is_critical" field.
func IsCriticalEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldIsCritical, v))
}

// IsCriticalContainsFold applies the ContainsFold predicate on the "is_critical" field.
func IsCriticalContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldIsCritical, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Release {
	return predicate.Release(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.Release {
	return predicate.Release(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Release) predicate.Release {
	return predicate.Release(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Release) predicate.Release {
	return predicate.Release(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Release) predicate.Release {
	return predicate.Release(sql.NotPredicates(p))
}