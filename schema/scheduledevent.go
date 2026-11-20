package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ScheduledEvent holds the schema definition for the ScheduledEvent entity: a task
// that the console runs on selected computers automatically, once or recurring.
type ScheduledEvent struct {
	ent.Schema
}

// Fields of the ScheduledEvent.
func (ScheduledEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").NotEmpty(),
		// Type is one of the actions the console supports: update all software,
		// power off, reboot or wake on LAN
		field.String("type").NotEmpty(),
		// Cron is a five field cron expression for recurring events, empty for
		// one-time events
		field.String("cron").Optional(),
		// When is the time of a one-time event; recurring events leave it empty
		field.Time("when").Optional().Nillable(),
		field.Bool("enabled").Default(true),
		field.Time("last_run").Optional().Nillable(),
		field.Bool("last_run_failed").Default(false),
		field.String("last_result").Optional(),
		// PackageID, PackageName, PackageBranch, PackageBrewType and PackageVerified
		// define the software package an installpackage or updatepackage event
		// deploys to the computers of its organization (or site)
		field.String("package_id").Optional(),
		field.String("package_name").Optional(),
		field.String("package_branch").Optional(),
		field.String("package_brew_type").Optional(),
		field.Bool("package_verified").Default(false),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ScheduledEvent.
func (ScheduledEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Ref("scheduled_events").Unique().Required(),
		// An event without a site applies to every site of its organization
		edge.From("site", Site.Type).Ref("scheduled_events").Unique(),
		// The profile whose tasks a runprofile event asks the agents to run
		edge.From("profile", Profile.Type).Ref("scheduled_events").Unique(),
	}
}
