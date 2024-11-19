// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"time"

	"github.com/doncicuto/openuem_ent/agent"
	"github.com/doncicuto/openuem_ent/deployment"
	"github.com/doncicuto/openuem_ent/logicaldisk"
	"github.com/doncicuto/openuem_ent/orgmetadata"
	"github.com/doncicuto/openuem_ent/revocation"
	"github.com/doncicuto/openuem_ent/schema"
	"github.com/doncicuto/openuem_ent/sessions"
	"github.com/doncicuto/openuem_ent/settings"
	"github.com/doncicuto/openuem_ent/tag"
	"github.com/doncicuto/openuem_ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	agentFields := schema.Agent{}.Fields()
	_ = agentFields
	// agentDescOs is the schema descriptor for os field.
	agentDescOs := agentFields[1].Descriptor()
	// agent.OsValidator is a validator for the "os" field. It is called by the builders before save.
	agent.OsValidator = agentDescOs.Validators[0].(func(string) error)
	// agentDescHostname is the schema descriptor for hostname field.
	agentDescHostname := agentFields[2].Descriptor()
	// agent.HostnameValidator is a validator for the "hostname" field. It is called by the builders before save.
	agent.HostnameValidator = agentDescHostname.Validators[0].(func(string) error)
	// agentDescVersion is the schema descriptor for version field.
	agentDescVersion := agentFields[3].Descriptor()
	// agent.VersionValidator is a validator for the "version" field. It is called by the builders before save.
	agent.VersionValidator = agentDescVersion.Validators[0].(func(string) error)
	// agentDescIP is the schema descriptor for ip field.
	agentDescIP := agentFields[4].Descriptor()
	// agent.DefaultIP holds the default value on creation for the ip field.
	agent.DefaultIP = agentDescIP.Default.(string)
	// agentDescMAC is the schema descriptor for mac field.
	agentDescMAC := agentFields[5].Descriptor()
	// agent.DefaultMAC holds the default value on creation for the mac field.
	agent.DefaultMAC = agentDescMAC.Default.(string)
	// agentDescEnabled is the schema descriptor for enabled field.
	agentDescEnabled := agentFields[8].Descriptor()
	// agent.DefaultEnabled holds the default value on creation for the enabled field.
	agent.DefaultEnabled = agentDescEnabled.Default.(bool)
	// agentDescVnc is the schema descriptor for vnc field.
	agentDescVnc := agentFields[9].Descriptor()
	// agent.DefaultVnc holds the default value on creation for the vnc field.
	agent.DefaultVnc = agentDescVnc.Default.(string)
	// agentDescUpdateTaskStatus is the schema descriptor for update_task_status field.
	agentDescUpdateTaskStatus := agentFields[11].Descriptor()
	// agent.DefaultUpdateTaskStatus holds the default value on creation for the update_task_status field.
	agent.DefaultUpdateTaskStatus = agentDescUpdateTaskStatus.Default.(string)
	// agentDescUpdateTaskDescription is the schema descriptor for update_task_description field.
	agentDescUpdateTaskDescription := agentFields[12].Descriptor()
	// agent.DefaultUpdateTaskDescription holds the default value on creation for the update_task_description field.
	agent.DefaultUpdateTaskDescription = agentDescUpdateTaskDescription.Default.(string)
	// agentDescUpdateTaskResult is the schema descriptor for update_task_result field.
	agentDescUpdateTaskResult := agentFields[13].Descriptor()
	// agent.DefaultUpdateTaskResult holds the default value on creation for the update_task_result field.
	agent.DefaultUpdateTaskResult = agentDescUpdateTaskResult.Default.(string)
	// agentDescUpdateTaskVersion is the schema descriptor for update_task_version field.
	agentDescUpdateTaskVersion := agentFields[15].Descriptor()
	// agent.DefaultUpdateTaskVersion holds the default value on creation for the update_task_version field.
	agent.DefaultUpdateTaskVersion = agentDescUpdateTaskVersion.Default.(string)
	// agentDescID is the schema descriptor for id field.
	agentDescID := agentFields[0].Descriptor()
	// agent.IDValidator is a validator for the "id" field. It is called by the builders before save.
	agent.IDValidator = agentDescID.Validators[0].(func(string) error)
	deploymentFields := schema.Deployment{}.Fields()
	_ = deploymentFields
	// deploymentDescInstalled is the schema descriptor for installed field.
	deploymentDescInstalled := deploymentFields[3].Descriptor()
	// deployment.DefaultInstalled holds the default value on creation for the installed field.
	deployment.DefaultInstalled = deploymentDescInstalled.Default.(func() time.Time)
	// deploymentDescUpdated is the schema descriptor for updated field.
	deploymentDescUpdated := deploymentFields[4].Descriptor()
	// deployment.DefaultUpdated holds the default value on creation for the updated field.
	deployment.DefaultUpdated = deploymentDescUpdated.Default.(func() time.Time)
	// deployment.UpdateDefaultUpdated holds the default value on update for the updated field.
	deployment.UpdateDefaultUpdated = deploymentDescUpdated.UpdateDefault.(func() time.Time)
	logicaldiskFields := schema.LogicalDisk{}.Fields()
	_ = logicaldiskFields
	// logicaldiskDescUsage is the schema descriptor for usage field.
	logicaldiskDescUsage := logicaldiskFields[2].Descriptor()
	// logicaldisk.DefaultUsage holds the default value on creation for the usage field.
	logicaldisk.DefaultUsage = logicaldiskDescUsage.Default.(int8)
	orgmetadataFields := schema.OrgMetadata{}.Fields()
	_ = orgmetadataFields
	// orgmetadataDescName is the schema descriptor for name field.
	orgmetadataDescName := orgmetadataFields[0].Descriptor()
	// orgmetadata.NameValidator is a validator for the "name" field. It is called by the builders before save.
	orgmetadata.NameValidator = orgmetadataDescName.Validators[0].(func(string) error)
	revocationFields := schema.Revocation{}.Fields()
	_ = revocationFields
	// revocationDescReason is the schema descriptor for reason field.
	revocationDescReason := revocationFields[1].Descriptor()
	// revocation.DefaultReason holds the default value on creation for the reason field.
	revocation.DefaultReason = revocationDescReason.Default.(int)
	// revocationDescRevoked is the schema descriptor for revoked field.
	revocationDescRevoked := revocationFields[4].Descriptor()
	// revocation.DefaultRevoked holds the default value on creation for the revoked field.
	revocation.DefaultRevoked = revocationDescRevoked.Default.(func() time.Time)
	sessionsFields := schema.Sessions{}.Fields()
	_ = sessionsFields
	// sessionsDescData is the schema descriptor for data field.
	sessionsDescData := sessionsFields[1].Descriptor()
	// sessions.DataValidator is a validator for the "data" field. It is called by the builders before save.
	sessions.DataValidator = sessionsDescData.Validators[0].(func([]byte) error)
	// sessionsDescID is the schema descriptor for id field.
	sessionsDescID := sessionsFields[0].Descriptor()
	// sessions.IDValidator is a validator for the "id" field. It is called by the builders before save.
	sessions.IDValidator = sessionsDescID.Validators[0].(func(string) error)
	settingsFields := schema.Settings{}.Fields()
	_ = settingsFields
	// settingsDescCountry is the schema descriptor for country field.
	settingsDescCountry := settingsFields[7].Descriptor()
	// settings.DefaultCountry holds the default value on creation for the country field.
	settings.DefaultCountry = settingsDescCountry.Default.(string)
	// settingsDescSMTPPort is the schema descriptor for smtp_port field.
	settingsDescSMTPPort := settingsFields[9].Descriptor()
	// settings.DefaultSMTPPort holds the default value on creation for the smtp_port field.
	settings.DefaultSMTPPort = settingsDescSMTPPort.Default.(int)
	// settingsDescSMTPAuth is the schema descriptor for smtp_auth field.
	settingsDescSMTPAuth := settingsFields[12].Descriptor()
	// settings.DefaultSMTPAuth holds the default value on creation for the smtp_auth field.
	settings.DefaultSMTPAuth = settingsDescSMTPAuth.Default.(string)
	// settingsDescSMTPTLS is the schema descriptor for smtp_tls field.
	settingsDescSMTPTLS := settingsFields[13].Descriptor()
	// settings.DefaultSMTPTLS holds the default value on creation for the smtp_tls field.
	settings.DefaultSMTPTLS = settingsDescSMTPTLS.Default.(bool)
	// settingsDescSMTPStarttls is the schema descriptor for smtp_starttls field.
	settingsDescSMTPStarttls := settingsFields[14].Descriptor()
	// settings.DefaultSMTPStarttls holds the default value on creation for the smtp_starttls field.
	settings.DefaultSMTPStarttls = settingsDescSMTPStarttls.Default.(bool)
	// settingsDescMaxUploadSize is the schema descriptor for max_upload_size field.
	settingsDescMaxUploadSize := settingsFields[18].Descriptor()
	// settings.DefaultMaxUploadSize holds the default value on creation for the max_upload_size field.
	settings.DefaultMaxUploadSize = settingsDescMaxUploadSize.Default.(string)
	// settingsDescUserCertYearsValid is the schema descriptor for user_cert_years_valid field.
	settingsDescUserCertYearsValid := settingsFields[19].Descriptor()
	// settings.DefaultUserCertYearsValid holds the default value on creation for the user_cert_years_valid field.
	settings.DefaultUserCertYearsValid = settingsDescUserCertYearsValid.Default.(int)
	// settingsDescNatsRequestTimeoutSeconds is the schema descriptor for nats_request_timeout_seconds field.
	settingsDescNatsRequestTimeoutSeconds := settingsFields[20].Descriptor()
	// settings.DefaultNatsRequestTimeoutSeconds holds the default value on creation for the nats_request_timeout_seconds field.
	settings.DefaultNatsRequestTimeoutSeconds = settingsDescNatsRequestTimeoutSeconds.Default.(int)
	// settingsDescRefreshTimeInMinutes is the schema descriptor for refresh_time_in_minutes field.
	settingsDescRefreshTimeInMinutes := settingsFields[21].Descriptor()
	// settings.DefaultRefreshTimeInMinutes holds the default value on creation for the refresh_time_in_minutes field.
	settings.DefaultRefreshTimeInMinutes = settingsDescRefreshTimeInMinutes.Default.(int)
	// settingsDescSessionLifetimeInMinutes is the schema descriptor for session_lifetime_in_minutes field.
	settingsDescSessionLifetimeInMinutes := settingsFields[22].Descriptor()
	// settings.DefaultSessionLifetimeInMinutes holds the default value on creation for the session_lifetime_in_minutes field.
	settings.DefaultSessionLifetimeInMinutes = settingsDescSessionLifetimeInMinutes.Default.(int)
	// settingsDescUpdateChannel is the schema descriptor for update_channel field.
	settingsDescUpdateChannel := settingsFields[23].Descriptor()
	// settings.DefaultUpdateChannel holds the default value on creation for the update_channel field.
	settings.DefaultUpdateChannel = settingsDescUpdateChannel.Default.(string)
	// settingsDescCreated is the schema descriptor for created field.
	settingsDescCreated := settingsFields[24].Descriptor()
	// settings.DefaultCreated holds the default value on creation for the created field.
	settings.DefaultCreated = settingsDescCreated.Default.(func() time.Time)
	// settingsDescModified is the schema descriptor for modified field.
	settingsDescModified := settingsFields[25].Descriptor()
	// settings.DefaultModified holds the default value on creation for the modified field.
	settings.DefaultModified = settingsDescModified.Default.(func() time.Time)
	// settings.UpdateDefaultModified holds the default value on update for the modified field.
	settings.UpdateDefaultModified = settingsDescModified.UpdateDefault.(func() time.Time)
	// settingsDescAgentReportFrequenceInMinutes is the schema descriptor for agent_report_frequence_in_minutes field.
	settingsDescAgentReportFrequenceInMinutes := settingsFields[26].Descriptor()
	// settings.DefaultAgentReportFrequenceInMinutes holds the default value on creation for the agent_report_frequence_in_minutes field.
	settings.DefaultAgentReportFrequenceInMinutes = settingsDescAgentReportFrequenceInMinutes.Default.(int)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescTag is the schema descriptor for tag field.
	tagDescTag := tagFields[0].Descriptor()
	// tag.TagValidator is a validator for the "tag" field. It is called by the builders before save.
	tag.TagValidator = tagDescTag.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmailVerified is the schema descriptor for email_verified field.
	userDescEmailVerified := userFields[5].Descriptor()
	// user.DefaultEmailVerified holds the default value on creation for the email_verified field.
	user.DefaultEmailVerified = userDescEmailVerified.Default.(bool)
	// userDescRegister is the schema descriptor for register field.
	userDescRegister := userFields[6].Descriptor()
	// user.DefaultRegister holds the default value on creation for the register field.
	user.DefaultRegister = userDescRegister.Default.(string)
	// userDescCreated is the schema descriptor for created field.
	userDescCreated := userFields[9].Descriptor()
	// user.DefaultCreated holds the default value on creation for the created field.
	user.DefaultCreated = userDescCreated.Default.(func() time.Time)
	// userDescModified is the schema descriptor for modified field.
	userDescModified := userFields[10].Descriptor()
	// user.DefaultModified holds the default value on creation for the modified field.
	user.DefaultModified = userDescModified.Default.(func() time.Time)
	// user.UpdateDefaultModified holds the default value on update for the modified field.
	user.UpdateDefaultModified = userDescModified.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
}
