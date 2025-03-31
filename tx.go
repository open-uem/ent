// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// Agent is the client for interacting with the Agent builders.
	Agent *AgentClient
	// Antivirus is the client for interacting with the Antivirus builders.
	Antivirus *AntivirusClient
	// App is the client for interacting with the App builders.
	App *AppClient
	// Certificate is the client for interacting with the Certificate builders.
	Certificate *CertificateClient
	// Computer is the client for interacting with the Computer builders.
	Computer *ComputerClient
	// Deployment is the client for interacting with the Deployment builders.
	Deployment *DeploymentClient
	// LogicalDisk is the client for interacting with the LogicalDisk builders.
	LogicalDisk *LogicalDiskClient
	// Metadata is the client for interacting with the Metadata builders.
	Metadata *MetadataClient
	// Monitor is the client for interacting with the Monitor builders.
	Monitor *MonitorClient
	// NetworkAdapter is the client for interacting with the NetworkAdapter builders.
	NetworkAdapter *NetworkAdapterClient
	// OperatingSystem is the client for interacting with the OperatingSystem builders.
	OperatingSystem *OperatingSystemClient
	// OrgMetadata is the client for interacting with the OrgMetadata builders.
	OrgMetadata *OrgMetadataClient
	// Printer is the client for interacting with the Printer builders.
	Printer *PrinterClient
	// Profile is the client for interacting with the Profile builders.
	Profile *ProfileClient
	// ProfileIssue is the client for interacting with the ProfileIssue builders.
	ProfileIssue *ProfileIssueClient
	// Release is the client for interacting with the Release builders.
	Release *ReleaseClient
	// Revocation is the client for interacting with the Revocation builders.
	Revocation *RevocationClient
	// Server is the client for interacting with the Server builders.
	Server *ServerClient
	// Sessions is the client for interacting with the Sessions builders.
	Sessions *SessionsClient
	// Settings is the client for interacting with the Settings builders.
	Settings *SettingsClient
	// Share is the client for interacting with the Share builders.
	Share *ShareClient
	// SoftwarePackage is the client for interacting with the SoftwarePackage builders.
	SoftwarePackage *SoftwarePackageClient
	// SystemUpdate is the client for interacting with the SystemUpdate builders.
	SystemUpdate *SystemUpdateClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// Task is the client for interacting with the Task builders.
	Task *TaskClient
	// Update is the client for interacting with the Update builders.
	Update *UpdateClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// WingetConfigExclusion is the client for interacting with the WingetConfigExclusion builders.
	WingetConfigExclusion *WingetConfigExclusionClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once
	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	txDriver.mu.Lock()
	hooks := append([]CommitHook(nil), txDriver.onCommit...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onCommit = append(txDriver.onCommit, f)
	txDriver.mu.Unlock()
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	txDriver.mu.Lock()
	hooks := append([]RollbackHook(nil), txDriver.onRollback...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onRollback = append(txDriver.onRollback, f)
	txDriver.mu.Unlock()
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.Agent = NewAgentClient(tx.config)
	tx.Antivirus = NewAntivirusClient(tx.config)
	tx.App = NewAppClient(tx.config)
	tx.Certificate = NewCertificateClient(tx.config)
	tx.Computer = NewComputerClient(tx.config)
	tx.Deployment = NewDeploymentClient(tx.config)
	tx.LogicalDisk = NewLogicalDiskClient(tx.config)
	tx.Metadata = NewMetadataClient(tx.config)
	tx.Monitor = NewMonitorClient(tx.config)
	tx.NetworkAdapter = NewNetworkAdapterClient(tx.config)
	tx.OperatingSystem = NewOperatingSystemClient(tx.config)
	tx.OrgMetadata = NewOrgMetadataClient(tx.config)
	tx.Printer = NewPrinterClient(tx.config)
	tx.Profile = NewProfileClient(tx.config)
	tx.ProfileIssue = NewProfileIssueClient(tx.config)
	tx.Release = NewReleaseClient(tx.config)
	tx.Revocation = NewRevocationClient(tx.config)
	tx.Server = NewServerClient(tx.config)
	tx.Sessions = NewSessionsClient(tx.config)
	tx.Settings = NewSettingsClient(tx.config)
	tx.Share = NewShareClient(tx.config)
	tx.SoftwarePackage = NewSoftwarePackageClient(tx.config)
	tx.SystemUpdate = NewSystemUpdateClient(tx.config)
	tx.Tag = NewTagClient(tx.config)
	tx.Task = NewTaskClient(tx.config)
	tx.Update = NewUpdateClient(tx.config)
	tx.User = NewUserClient(tx.config)
	tx.WingetConfigExclusion = NewWingetConfigExclusionClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: Agent.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
	// completion hooks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v any) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v any) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
