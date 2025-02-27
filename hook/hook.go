// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/open-uem/ent"
)

// The AgentFunc type is an adapter to allow the use of ordinary
// function as Agent mutator.
type AgentFunc func(context.Context, *ent.AgentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AgentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AgentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AgentMutation", m)
}

// The AntivirusFunc type is an adapter to allow the use of ordinary
// function as Antivirus mutator.
type AntivirusFunc func(context.Context, *ent.AntivirusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AntivirusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AntivirusMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AntivirusMutation", m)
}

// The AppFunc type is an adapter to allow the use of ordinary
// function as App mutator.
type AppFunc func(context.Context, *ent.AppMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppMutation", m)
}

// The CertificateFunc type is an adapter to allow the use of ordinary
// function as Certificate mutator.
type CertificateFunc func(context.Context, *ent.CertificateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CertificateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CertificateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CertificateMutation", m)
}

// The ComputerFunc type is an adapter to allow the use of ordinary
// function as Computer mutator.
type ComputerFunc func(context.Context, *ent.ComputerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ComputerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ComputerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ComputerMutation", m)
}

// The DeploymentFunc type is an adapter to allow the use of ordinary
// function as Deployment mutator.
type DeploymentFunc func(context.Context, *ent.DeploymentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeploymentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DeploymentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeploymentMutation", m)
}

// The LogicalDiskFunc type is an adapter to allow the use of ordinary
// function as LogicalDisk mutator.
type LogicalDiskFunc func(context.Context, *ent.LogicalDiskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LogicalDiskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LogicalDiskMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LogicalDiskMutation", m)
}

// The MetadataFunc type is an adapter to allow the use of ordinary
// function as Metadata mutator.
type MetadataFunc func(context.Context, *ent.MetadataMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MetadataFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MetadataMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MetadataMutation", m)
}

// The MonitorFunc type is an adapter to allow the use of ordinary
// function as Monitor mutator.
type MonitorFunc func(context.Context, *ent.MonitorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MonitorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MonitorMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MonitorMutation", m)
}

// The NetworkAdapterFunc type is an adapter to allow the use of ordinary
// function as NetworkAdapter mutator.
type NetworkAdapterFunc func(context.Context, *ent.NetworkAdapterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NetworkAdapterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NetworkAdapterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NetworkAdapterMutation", m)
}

// The OperatingSystemFunc type is an adapter to allow the use of ordinary
// function as OperatingSystem mutator.
type OperatingSystemFunc func(context.Context, *ent.OperatingSystemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OperatingSystemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OperatingSystemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OperatingSystemMutation", m)
}

// The OrgMetadataFunc type is an adapter to allow the use of ordinary
// function as OrgMetadata mutator.
type OrgMetadataFunc func(context.Context, *ent.OrgMetadataMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrgMetadataFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrgMetadataMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrgMetadataMutation", m)
}

// The PrinterFunc type is an adapter to allow the use of ordinary
// function as Printer mutator.
type PrinterFunc func(context.Context, *ent.PrinterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PrinterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PrinterMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PrinterMutation", m)
}

// The ProfileFunc type is an adapter to allow the use of ordinary
// function as Profile mutator.
type ProfileFunc func(context.Context, *ent.ProfileMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProfileFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProfileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProfileMutation", m)
}

// The ReleaseFunc type is an adapter to allow the use of ordinary
// function as Release mutator.
type ReleaseFunc func(context.Context, *ent.ReleaseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReleaseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReleaseMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReleaseMutation", m)
}

// The RevocationFunc type is an adapter to allow the use of ordinary
// function as Revocation mutator.
type RevocationFunc func(context.Context, *ent.RevocationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RevocationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RevocationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RevocationMutation", m)
}

// The ServerFunc type is an adapter to allow the use of ordinary
// function as Server mutator.
type ServerFunc func(context.Context, *ent.ServerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ServerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ServerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ServerMutation", m)
}

// The SessionsFunc type is an adapter to allow the use of ordinary
// function as Sessions mutator.
type SessionsFunc func(context.Context, *ent.SessionsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SessionsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SessionsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SessionsMutation", m)
}

// The SettingsFunc type is an adapter to allow the use of ordinary
// function as Settings mutator.
type SettingsFunc func(context.Context, *ent.SettingsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SettingsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SettingsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SettingsMutation", m)
}

// The ShareFunc type is an adapter to allow the use of ordinary
// function as Share mutator.
type ShareFunc func(context.Context, *ent.ShareMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ShareFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ShareMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ShareMutation", m)
}

// The SystemUpdateFunc type is an adapter to allow the use of ordinary
// function as SystemUpdate mutator.
type SystemUpdateFunc func(context.Context, *ent.SystemUpdateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SystemUpdateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SystemUpdateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SystemUpdateMutation", m)
}

// The TagFunc type is an adapter to allow the use of ordinary
// function as Tag mutator.
type TagFunc func(context.Context, *ent.TagMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TagFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TagMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TagMutation", m)
}

// The TaskFunc type is an adapter to allow the use of ordinary
// function as Task mutator.
type TaskFunc func(context.Context, *ent.TaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TaskMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskMutation", m)
}

// The UpdateFunc type is an adapter to allow the use of ordinary
// function as Update mutator.
type UpdateFunc func(context.Context, *ent.UpdateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UpdateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UpdateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UpdateMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
