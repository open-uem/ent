// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/doncicuto/openuem_ent/agent"
	"github.com/doncicuto/openuem_ent/antivirus"
	"github.com/doncicuto/openuem_ent/app"
	"github.com/doncicuto/openuem_ent/computer"
	"github.com/doncicuto/openuem_ent/deployment"
	"github.com/doncicuto/openuem_ent/logicaldisk"
	"github.com/doncicuto/openuem_ent/monitor"
	"github.com/doncicuto/openuem_ent/networkadapter"
	"github.com/doncicuto/openuem_ent/operatingsystem"
	"github.com/doncicuto/openuem_ent/printer"
	"github.com/doncicuto/openuem_ent/share"
	"github.com/doncicuto/openuem_ent/systemupdate"
)

// AgentCreate is the builder for creating a Agent entity.
type AgentCreate struct {
	config
	mutation *AgentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOs sets the "os" field.
func (ac *AgentCreate) SetOs(s string) *AgentCreate {
	ac.mutation.SetOs(s)
	return ac
}

// SetHostname sets the "hostname" field.
func (ac *AgentCreate) SetHostname(s string) *AgentCreate {
	ac.mutation.SetHostname(s)
	return ac
}

// SetVersion sets the "version" field.
func (ac *AgentCreate) SetVersion(s string) *AgentCreate {
	ac.mutation.SetVersion(s)
	return ac
}

// SetIP sets the "ip" field.
func (ac *AgentCreate) SetIP(s string) *AgentCreate {
	ac.mutation.SetIP(s)
	return ac
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (ac *AgentCreate) SetNillableIP(s *string) *AgentCreate {
	if s != nil {
		ac.SetIP(*s)
	}
	return ac
}

// SetFirstContact sets the "first_contact" field.
func (ac *AgentCreate) SetFirstContact(t time.Time) *AgentCreate {
	ac.mutation.SetFirstContact(t)
	return ac
}

// SetNillableFirstContact sets the "first_contact" field if the given value is not nil.
func (ac *AgentCreate) SetNillableFirstContact(t *time.Time) *AgentCreate {
	if t != nil {
		ac.SetFirstContact(*t)
	}
	return ac
}

// SetLastContact sets the "last_contact" field.
func (ac *AgentCreate) SetLastContact(t time.Time) *AgentCreate {
	ac.mutation.SetLastContact(t)
	return ac
}

// SetNillableLastContact sets the "last_contact" field if the given value is not nil.
func (ac *AgentCreate) SetNillableLastContact(t *time.Time) *AgentCreate {
	if t != nil {
		ac.SetLastContact(*t)
	}
	return ac
}

// SetEnabled sets the "enabled" field.
func (ac *AgentCreate) SetEnabled(b bool) *AgentCreate {
	ac.mutation.SetEnabled(b)
	return ac
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (ac *AgentCreate) SetNillableEnabled(b *bool) *AgentCreate {
	if b != nil {
		ac.SetEnabled(*b)
	}
	return ac
}

// SetVnc sets the "vnc" field.
func (ac *AgentCreate) SetVnc(s string) *AgentCreate {
	ac.mutation.SetVnc(s)
	return ac
}

// SetNillableVnc sets the "vnc" field if the given value is not nil.
func (ac *AgentCreate) SetNillableVnc(s *string) *AgentCreate {
	if s != nil {
		ac.SetVnc(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AgentCreate) SetID(s string) *AgentCreate {
	ac.mutation.SetID(s)
	return ac
}

// SetComputerID sets the "computer" edge to the Computer entity by ID.
func (ac *AgentCreate) SetComputerID(id int) *AgentCreate {
	ac.mutation.SetComputerID(id)
	return ac
}

// SetNillableComputerID sets the "computer" edge to the Computer entity by ID if the given value is not nil.
func (ac *AgentCreate) SetNillableComputerID(id *int) *AgentCreate {
	if id != nil {
		ac = ac.SetComputerID(*id)
	}
	return ac
}

// SetComputer sets the "computer" edge to the Computer entity.
func (ac *AgentCreate) SetComputer(c *Computer) *AgentCreate {
	return ac.SetComputerID(c.ID)
}

// SetOperatingsystemID sets the "operatingsystem" edge to the OperatingSystem entity by ID.
func (ac *AgentCreate) SetOperatingsystemID(id int) *AgentCreate {
	ac.mutation.SetOperatingsystemID(id)
	return ac
}

// SetNillableOperatingsystemID sets the "operatingsystem" edge to the OperatingSystem entity by ID if the given value is not nil.
func (ac *AgentCreate) SetNillableOperatingsystemID(id *int) *AgentCreate {
	if id != nil {
		ac = ac.SetOperatingsystemID(*id)
	}
	return ac
}

// SetOperatingsystem sets the "operatingsystem" edge to the OperatingSystem entity.
func (ac *AgentCreate) SetOperatingsystem(o *OperatingSystem) *AgentCreate {
	return ac.SetOperatingsystemID(o.ID)
}

// SetSystemupdateID sets the "systemupdate" edge to the SystemUpdate entity by ID.
func (ac *AgentCreate) SetSystemupdateID(id int) *AgentCreate {
	ac.mutation.SetSystemupdateID(id)
	return ac
}

// SetNillableSystemupdateID sets the "systemupdate" edge to the SystemUpdate entity by ID if the given value is not nil.
func (ac *AgentCreate) SetNillableSystemupdateID(id *int) *AgentCreate {
	if id != nil {
		ac = ac.SetSystemupdateID(*id)
	}
	return ac
}

// SetSystemupdate sets the "systemupdate" edge to the SystemUpdate entity.
func (ac *AgentCreate) SetSystemupdate(s *SystemUpdate) *AgentCreate {
	return ac.SetSystemupdateID(s.ID)
}

// SetAntivirusID sets the "antivirus" edge to the Antivirus entity by ID.
func (ac *AgentCreate) SetAntivirusID(id int) *AgentCreate {
	ac.mutation.SetAntivirusID(id)
	return ac
}

// SetNillableAntivirusID sets the "antivirus" edge to the Antivirus entity by ID if the given value is not nil.
func (ac *AgentCreate) SetNillableAntivirusID(id *int) *AgentCreate {
	if id != nil {
		ac = ac.SetAntivirusID(*id)
	}
	return ac
}

// SetAntivirus sets the "antivirus" edge to the Antivirus entity.
func (ac *AgentCreate) SetAntivirus(a *Antivirus) *AgentCreate {
	return ac.SetAntivirusID(a.ID)
}

// AddLogicaldiskIDs adds the "logicaldisks" edge to the LogicalDisk entity by IDs.
func (ac *AgentCreate) AddLogicaldiskIDs(ids ...int) *AgentCreate {
	ac.mutation.AddLogicaldiskIDs(ids...)
	return ac
}

// AddLogicaldisks adds the "logicaldisks" edges to the LogicalDisk entity.
func (ac *AgentCreate) AddLogicaldisks(l ...*LogicalDisk) *AgentCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ac.AddLogicaldiskIDs(ids...)
}

// AddAppIDs adds the "apps" edge to the App entity by IDs.
func (ac *AgentCreate) AddAppIDs(ids ...int) *AgentCreate {
	ac.mutation.AddAppIDs(ids...)
	return ac
}

// AddApps adds the "apps" edges to the App entity.
func (ac *AgentCreate) AddApps(a ...*App) *AgentCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAppIDs(ids...)
}

// AddMonitorIDs adds the "monitors" edge to the Monitor entity by IDs.
func (ac *AgentCreate) AddMonitorIDs(ids ...int) *AgentCreate {
	ac.mutation.AddMonitorIDs(ids...)
	return ac
}

// AddMonitors adds the "monitors" edges to the Monitor entity.
func (ac *AgentCreate) AddMonitors(m ...*Monitor) *AgentCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ac.AddMonitorIDs(ids...)
}

// AddShareIDs adds the "shares" edge to the Share entity by IDs.
func (ac *AgentCreate) AddShareIDs(ids ...int) *AgentCreate {
	ac.mutation.AddShareIDs(ids...)
	return ac
}

// AddShares adds the "shares" edges to the Share entity.
func (ac *AgentCreate) AddShares(s ...*Share) *AgentCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddShareIDs(ids...)
}

// AddPrinterIDs adds the "printers" edge to the Printer entity by IDs.
func (ac *AgentCreate) AddPrinterIDs(ids ...int) *AgentCreate {
	ac.mutation.AddPrinterIDs(ids...)
	return ac
}

// AddPrinters adds the "printers" edges to the Printer entity.
func (ac *AgentCreate) AddPrinters(p ...*Printer) *AgentCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPrinterIDs(ids...)
}

// AddNetworkadapterIDs adds the "networkadapters" edge to the NetworkAdapter entity by IDs.
func (ac *AgentCreate) AddNetworkadapterIDs(ids ...int) *AgentCreate {
	ac.mutation.AddNetworkadapterIDs(ids...)
	return ac
}

// AddNetworkadapters adds the "networkadapters" edges to the NetworkAdapter entity.
func (ac *AgentCreate) AddNetworkadapters(n ...*NetworkAdapter) *AgentCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ac.AddNetworkadapterIDs(ids...)
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (ac *AgentCreate) AddDeploymentIDs(ids ...int) *AgentCreate {
	ac.mutation.AddDeploymentIDs(ids...)
	return ac
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (ac *AgentCreate) AddDeployments(d ...*Deployment) *AgentCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ac.AddDeploymentIDs(ids...)
}

// Mutation returns the AgentMutation object of the builder.
func (ac *AgentCreate) Mutation() *AgentMutation {
	return ac.mutation
}

// Save creates the Agent in the database.
func (ac *AgentCreate) Save(ctx context.Context) (*Agent, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AgentCreate) SaveX(ctx context.Context) *Agent {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AgentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AgentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AgentCreate) defaults() {
	if _, ok := ac.mutation.IP(); !ok {
		v := agent.DefaultIP
		ac.mutation.SetIP(v)
	}
	if _, ok := ac.mutation.Enabled(); !ok {
		v := agent.DefaultEnabled
		ac.mutation.SetEnabled(v)
	}
	if _, ok := ac.mutation.Vnc(); !ok {
		v := agent.DefaultVnc
		ac.mutation.SetVnc(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AgentCreate) check() error {
	if _, ok := ac.mutation.Os(); !ok {
		return &ValidationError{Name: "os", err: errors.New(`openuem_ent: missing required field "Agent.os"`)}
	}
	if v, ok := ac.mutation.Os(); ok {
		if err := agent.OsValidator(v); err != nil {
			return &ValidationError{Name: "os", err: fmt.Errorf(`openuem_ent: validator failed for field "Agent.os": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Hostname(); !ok {
		return &ValidationError{Name: "hostname", err: errors.New(`openuem_ent: missing required field "Agent.hostname"`)}
	}
	if v, ok := ac.mutation.Hostname(); ok {
		if err := agent.HostnameValidator(v); err != nil {
			return &ValidationError{Name: "hostname", err: fmt.Errorf(`openuem_ent: validator failed for field "Agent.hostname": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`openuem_ent: missing required field "Agent.version"`)}
	}
	if v, ok := ac.mutation.Version(); ok {
		if err := agent.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`openuem_ent: validator failed for field "Agent.version": %w`, err)}
		}
	}
	if _, ok := ac.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`openuem_ent: missing required field "Agent.ip"`)}
	}
	if _, ok := ac.mutation.Enabled(); !ok {
		return &ValidationError{Name: "enabled", err: errors.New(`openuem_ent: missing required field "Agent.enabled"`)}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := agent.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`openuem_ent: validator failed for field "Agent.id": %w`, err)}
		}
	}
	return nil
}

func (ac *AgentCreate) sqlSave(ctx context.Context) (*Agent, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Agent.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AgentCreate) createSpec() (*Agent, *sqlgraph.CreateSpec) {
	var (
		_node = &Agent{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(agent.Table, sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Os(); ok {
		_spec.SetField(agent.FieldOs, field.TypeString, value)
		_node.Os = value
	}
	if value, ok := ac.mutation.Hostname(); ok {
		_spec.SetField(agent.FieldHostname, field.TypeString, value)
		_node.Hostname = value
	}
	if value, ok := ac.mutation.Version(); ok {
		_spec.SetField(agent.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := ac.mutation.IP(); ok {
		_spec.SetField(agent.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := ac.mutation.FirstContact(); ok {
		_spec.SetField(agent.FieldFirstContact, field.TypeTime, value)
		_node.FirstContact = value
	}
	if value, ok := ac.mutation.LastContact(); ok {
		_spec.SetField(agent.FieldLastContact, field.TypeTime, value)
		_node.LastContact = value
	}
	if value, ok := ac.mutation.Enabled(); ok {
		_spec.SetField(agent.FieldEnabled, field.TypeBool, value)
		_node.Enabled = value
	}
	if value, ok := ac.mutation.Vnc(); ok {
		_spec.SetField(agent.FieldVnc, field.TypeString, value)
		_node.Vnc = value
	}
	if nodes := ac.mutation.ComputerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   agent.ComputerTable,
			Columns: []string{agent.ComputerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(computer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.OperatingsystemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   agent.OperatingsystemTable,
			Columns: []string{agent.OperatingsystemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operatingsystem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.SystemupdateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   agent.SystemupdateTable,
			Columns: []string{agent.SystemupdateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(systemupdate.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AntivirusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   agent.AntivirusTable,
			Columns: []string{agent.AntivirusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(antivirus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.LogicaldisksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.LogicaldisksTable,
			Columns: []string{agent.LogicaldisksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(logicaldisk.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AppsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.AppsTable,
			Columns: []string{agent.AppsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.MonitorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.MonitorsTable,
			Columns: []string{agent.MonitorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(monitor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.SharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.SharesTable,
			Columns: []string{agent.SharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.PrintersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.PrintersTable,
			Columns: []string{agent.PrintersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(printer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.NetworkadaptersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.NetworkadaptersTable,
			Columns: []string{agent.NetworkadaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(networkadapter.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.DeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agent.DeploymentsTable,
			Columns: []string{agent.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Agent.Create().
//		SetOs(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AgentUpsert) {
//			SetOs(v+v).
//		}).
//		Exec(ctx)
func (ac *AgentCreate) OnConflict(opts ...sql.ConflictOption) *AgentUpsertOne {
	ac.conflict = opts
	return &AgentUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Agent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AgentCreate) OnConflictColumns(columns ...string) *AgentUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AgentUpsertOne{
		create: ac,
	}
}

type (
	// AgentUpsertOne is the builder for "upsert"-ing
	//  one Agent node.
	AgentUpsertOne struct {
		create *AgentCreate
	}

	// AgentUpsert is the "OnConflict" setter.
	AgentUpsert struct {
		*sql.UpdateSet
	}
)

// SetOs sets the "os" field.
func (u *AgentUpsert) SetOs(v string) *AgentUpsert {
	u.Set(agent.FieldOs, v)
	return u
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *AgentUpsert) UpdateOs() *AgentUpsert {
	u.SetExcluded(agent.FieldOs)
	return u
}

// SetHostname sets the "hostname" field.
func (u *AgentUpsert) SetHostname(v string) *AgentUpsert {
	u.Set(agent.FieldHostname, v)
	return u
}

// UpdateHostname sets the "hostname" field to the value that was provided on create.
func (u *AgentUpsert) UpdateHostname() *AgentUpsert {
	u.SetExcluded(agent.FieldHostname)
	return u
}

// SetVersion sets the "version" field.
func (u *AgentUpsert) SetVersion(v string) *AgentUpsert {
	u.Set(agent.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *AgentUpsert) UpdateVersion() *AgentUpsert {
	u.SetExcluded(agent.FieldVersion)
	return u
}

// SetIP sets the "ip" field.
func (u *AgentUpsert) SetIP(v string) *AgentUpsert {
	u.Set(agent.FieldIP, v)
	return u
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *AgentUpsert) UpdateIP() *AgentUpsert {
	u.SetExcluded(agent.FieldIP)
	return u
}

// SetFirstContact sets the "first_contact" field.
func (u *AgentUpsert) SetFirstContact(v time.Time) *AgentUpsert {
	u.Set(agent.FieldFirstContact, v)
	return u
}

// UpdateFirstContact sets the "first_contact" field to the value that was provided on create.
func (u *AgentUpsert) UpdateFirstContact() *AgentUpsert {
	u.SetExcluded(agent.FieldFirstContact)
	return u
}

// ClearFirstContact clears the value of the "first_contact" field.
func (u *AgentUpsert) ClearFirstContact() *AgentUpsert {
	u.SetNull(agent.FieldFirstContact)
	return u
}

// SetLastContact sets the "last_contact" field.
func (u *AgentUpsert) SetLastContact(v time.Time) *AgentUpsert {
	u.Set(agent.FieldLastContact, v)
	return u
}

// UpdateLastContact sets the "last_contact" field to the value that was provided on create.
func (u *AgentUpsert) UpdateLastContact() *AgentUpsert {
	u.SetExcluded(agent.FieldLastContact)
	return u
}

// ClearLastContact clears the value of the "last_contact" field.
func (u *AgentUpsert) ClearLastContact() *AgentUpsert {
	u.SetNull(agent.FieldLastContact)
	return u
}

// SetEnabled sets the "enabled" field.
func (u *AgentUpsert) SetEnabled(v bool) *AgentUpsert {
	u.Set(agent.FieldEnabled, v)
	return u
}

// UpdateEnabled sets the "enabled" field to the value that was provided on create.
func (u *AgentUpsert) UpdateEnabled() *AgentUpsert {
	u.SetExcluded(agent.FieldEnabled)
	return u
}

// SetVnc sets the "vnc" field.
func (u *AgentUpsert) SetVnc(v string) *AgentUpsert {
	u.Set(agent.FieldVnc, v)
	return u
}

// UpdateVnc sets the "vnc" field to the value that was provided on create.
func (u *AgentUpsert) UpdateVnc() *AgentUpsert {
	u.SetExcluded(agent.FieldVnc)
	return u
}

// ClearVnc clears the value of the "vnc" field.
func (u *AgentUpsert) ClearVnc() *AgentUpsert {
	u.SetNull(agent.FieldVnc)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Agent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(agent.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AgentUpsertOne) UpdateNewValues() *AgentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(agent.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Agent.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AgentUpsertOne) Ignore() *AgentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AgentUpsertOne) DoNothing() *AgentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AgentCreate.OnConflict
// documentation for more info.
func (u *AgentUpsertOne) Update(set func(*AgentUpsert)) *AgentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AgentUpsert{UpdateSet: update})
	}))
	return u
}

// SetOs sets the "os" field.
func (u *AgentUpsertOne) SetOs(v string) *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.SetOs(v)
	})
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *AgentUpsertOne) UpdateOs() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateOs()
	})
}

// SetHostname sets the "hostname" field.
func (u *AgentUpsertOne) SetHostname(v string) *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.SetHostname(v)
	})
}

// UpdateHostname sets the "hostname" field to the value that was provided on create.
func (u *AgentUpsertOne) UpdateHostname() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateHostname()
	})
}

// SetVersion sets the "version" field.
func (u *AgentUpsertOne) SetVersion(v string) *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *AgentUpsertOne) UpdateVersion() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateVersion()
	})
}

// SetIP sets the "ip" field.
func (u *AgentUpsertOne) SetIP(v string) *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *AgentUpsertOne) UpdateIP() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateIP()
	})
}

// SetFirstContact sets the "first_contact" field.
func (u *AgentUpsertOne) SetFirstContact(v time.Time) *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.SetFirstContact(v)
	})
}

// UpdateFirstContact sets the "first_contact" field to the value that was provided on create.
func (u *AgentUpsertOne) UpdateFirstContact() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateFirstContact()
	})
}

// ClearFirstContact clears the value of the "first_contact" field.
func (u *AgentUpsertOne) ClearFirstContact() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.ClearFirstContact()
	})
}

// SetLastContact sets the "last_contact" field.
func (u *AgentUpsertOne) SetLastContact(v time.Time) *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.SetLastContact(v)
	})
}

// UpdateLastContact sets the "last_contact" field to the value that was provided on create.
func (u *AgentUpsertOne) UpdateLastContact() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateLastContact()
	})
}

// ClearLastContact clears the value of the "last_contact" field.
func (u *AgentUpsertOne) ClearLastContact() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.ClearLastContact()
	})
}

// SetEnabled sets the "enabled" field.
func (u *AgentUpsertOne) SetEnabled(v bool) *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.SetEnabled(v)
	})
}

// UpdateEnabled sets the "enabled" field to the value that was provided on create.
func (u *AgentUpsertOne) UpdateEnabled() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateEnabled()
	})
}

// SetVnc sets the "vnc" field.
func (u *AgentUpsertOne) SetVnc(v string) *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.SetVnc(v)
	})
}

// UpdateVnc sets the "vnc" field to the value that was provided on create.
func (u *AgentUpsertOne) UpdateVnc() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateVnc()
	})
}

// ClearVnc clears the value of the "vnc" field.
func (u *AgentUpsertOne) ClearVnc() *AgentUpsertOne {
	return u.Update(func(s *AgentUpsert) {
		s.ClearVnc()
	})
}

// Exec executes the query.
func (u *AgentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for AgentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AgentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AgentUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("openuem_ent: AgentUpsertOne.ID is not supported by MySQL driver. Use AgentUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AgentUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AgentCreateBulk is the builder for creating many Agent entities in bulk.
type AgentCreateBulk struct {
	config
	err      error
	builders []*AgentCreate
	conflict []sql.ConflictOption
}

// Save creates the Agent entities in the database.
func (acb *AgentCreateBulk) Save(ctx context.Context) ([]*Agent, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Agent, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AgentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AgentCreateBulk) SaveX(ctx context.Context) []*Agent {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AgentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AgentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Agent.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AgentUpsert) {
//			SetOs(v+v).
//		}).
//		Exec(ctx)
func (acb *AgentCreateBulk) OnConflict(opts ...sql.ConflictOption) *AgentUpsertBulk {
	acb.conflict = opts
	return &AgentUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Agent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AgentCreateBulk) OnConflictColumns(columns ...string) *AgentUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AgentUpsertBulk{
		create: acb,
	}
}

// AgentUpsertBulk is the builder for "upsert"-ing
// a bulk of Agent nodes.
type AgentUpsertBulk struct {
	create *AgentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Agent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(agent.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AgentUpsertBulk) UpdateNewValues() *AgentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(agent.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Agent.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AgentUpsertBulk) Ignore() *AgentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AgentUpsertBulk) DoNothing() *AgentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AgentCreateBulk.OnConflict
// documentation for more info.
func (u *AgentUpsertBulk) Update(set func(*AgentUpsert)) *AgentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AgentUpsert{UpdateSet: update})
	}))
	return u
}

// SetOs sets the "os" field.
func (u *AgentUpsertBulk) SetOs(v string) *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.SetOs(v)
	})
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *AgentUpsertBulk) UpdateOs() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateOs()
	})
}

// SetHostname sets the "hostname" field.
func (u *AgentUpsertBulk) SetHostname(v string) *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.SetHostname(v)
	})
}

// UpdateHostname sets the "hostname" field to the value that was provided on create.
func (u *AgentUpsertBulk) UpdateHostname() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateHostname()
	})
}

// SetVersion sets the "version" field.
func (u *AgentUpsertBulk) SetVersion(v string) *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *AgentUpsertBulk) UpdateVersion() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateVersion()
	})
}

// SetIP sets the "ip" field.
func (u *AgentUpsertBulk) SetIP(v string) *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *AgentUpsertBulk) UpdateIP() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateIP()
	})
}

// SetFirstContact sets the "first_contact" field.
func (u *AgentUpsertBulk) SetFirstContact(v time.Time) *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.SetFirstContact(v)
	})
}

// UpdateFirstContact sets the "first_contact" field to the value that was provided on create.
func (u *AgentUpsertBulk) UpdateFirstContact() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateFirstContact()
	})
}

// ClearFirstContact clears the value of the "first_contact" field.
func (u *AgentUpsertBulk) ClearFirstContact() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.ClearFirstContact()
	})
}

// SetLastContact sets the "last_contact" field.
func (u *AgentUpsertBulk) SetLastContact(v time.Time) *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.SetLastContact(v)
	})
}

// UpdateLastContact sets the "last_contact" field to the value that was provided on create.
func (u *AgentUpsertBulk) UpdateLastContact() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateLastContact()
	})
}

// ClearLastContact clears the value of the "last_contact" field.
func (u *AgentUpsertBulk) ClearLastContact() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.ClearLastContact()
	})
}

// SetEnabled sets the "enabled" field.
func (u *AgentUpsertBulk) SetEnabled(v bool) *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.SetEnabled(v)
	})
}

// UpdateEnabled sets the "enabled" field to the value that was provided on create.
func (u *AgentUpsertBulk) UpdateEnabled() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateEnabled()
	})
}

// SetVnc sets the "vnc" field.
func (u *AgentUpsertBulk) SetVnc(v string) *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.SetVnc(v)
	})
}

// UpdateVnc sets the "vnc" field to the value that was provided on create.
func (u *AgentUpsertBulk) UpdateVnc() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.UpdateVnc()
	})
}

// ClearVnc clears the value of the "vnc" field.
func (u *AgentUpsertBulk) ClearVnc() *AgentUpsertBulk {
	return u.Update(func(s *AgentUpsert) {
		s.ClearVnc()
	})
}

// Exec executes the query.
func (u *AgentUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("openuem_ent: OnConflict was set for builder %d. Set it on the AgentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for AgentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AgentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
