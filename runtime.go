// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/deployment"
	"github.com/open-uem/ent/logicaldisk"
	"github.com/open-uem/ent/orgmetadata"
	"github.com/open-uem/ent/profile"
	"github.com/open-uem/ent/profileissue"
	"github.com/open-uem/ent/revocation"
	"github.com/open-uem/ent/schema"
	"github.com/open-uem/ent/sessions"
	"github.com/open-uem/ent/settings"
	"github.com/open-uem/ent/tag"
	"github.com/open-uem/ent/task"
	"github.com/open-uem/ent/user"
	"github.com/open-uem/ent/wingetconfigexclusion"
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
	// agentDescIP is the schema descriptor for ip field.
	agentDescIP := agentFields[3].Descriptor()
	// agent.DefaultIP holds the default value on creation for the ip field.
	agent.DefaultIP = agentDescIP.Default.(string)
	// agentDescMAC is the schema descriptor for mac field.
	agentDescMAC := agentFields[4].Descriptor()
	// agent.DefaultMAC holds the default value on creation for the mac field.
	agent.DefaultMAC = agentDescMAC.Default.(string)
	// agentDescVnc is the schema descriptor for vnc field.
	agentDescVnc := agentFields[7].Descriptor()
	// agent.DefaultVnc holds the default value on creation for the vnc field.
	agent.DefaultVnc = agentDescVnc.Default.(string)
	// agentDescUpdateTaskStatus is the schema descriptor for update_task_status field.
	agentDescUpdateTaskStatus := agentFields[9].Descriptor()
	// agent.DefaultUpdateTaskStatus holds the default value on creation for the update_task_status field.
	agent.DefaultUpdateTaskStatus = agentDescUpdateTaskStatus.Default.(string)
	// agentDescUpdateTaskDescription is the schema descriptor for update_task_description field.
	agentDescUpdateTaskDescription := agentFields[10].Descriptor()
	// agent.DefaultUpdateTaskDescription holds the default value on creation for the update_task_description field.
	agent.DefaultUpdateTaskDescription = agentDescUpdateTaskDescription.Default.(string)
	// agentDescUpdateTaskResult is the schema descriptor for update_task_result field.
	agentDescUpdateTaskResult := agentFields[11].Descriptor()
	// agent.DefaultUpdateTaskResult holds the default value on creation for the update_task_result field.
	agent.DefaultUpdateTaskResult = agentDescUpdateTaskResult.Default.(string)
	// agentDescUpdateTaskVersion is the schema descriptor for update_task_version field.
	agentDescUpdateTaskVersion := agentFields[13].Descriptor()
	// agent.DefaultUpdateTaskVersion holds the default value on creation for the update_task_version field.
	agent.DefaultUpdateTaskVersion = agentDescUpdateTaskVersion.Default.(string)
	// agentDescVncProxyPort is the schema descriptor for vnc_proxy_port field.
	agentDescVncProxyPort := agentFields[14].Descriptor()
	// agent.DefaultVncProxyPort holds the default value on creation for the vnc_proxy_port field.
	agent.DefaultVncProxyPort = agentDescVncProxyPort.Default.(string)
	// agentDescSftpPort is the schema descriptor for sftp_port field.
	agentDescSftpPort := agentFields[15].Descriptor()
	// agent.DefaultSftpPort holds the default value on creation for the sftp_port field.
	agent.DefaultSftpPort = agentDescSftpPort.Default.(string)
	// agentDescCertificateReady is the schema descriptor for certificate_ready field.
	agentDescCertificateReady := agentFields[17].Descriptor()
	// agent.DefaultCertificateReady holds the default value on creation for the certificate_ready field.
	agent.DefaultCertificateReady = agentDescCertificateReady.Default.(bool)
	// agentDescRestartRequired is the schema descriptor for restart_required field.
	agentDescRestartRequired := agentFields[18].Descriptor()
	// agent.DefaultRestartRequired holds the default value on creation for the restart_required field.
	agent.DefaultRestartRequired = agentDescRestartRequired.Default.(bool)
	// agentDescIsRemote is the schema descriptor for is_remote field.
	agentDescIsRemote := agentFields[19].Descriptor()
	// agent.DefaultIsRemote holds the default value on creation for the is_remote field.
	agent.DefaultIsRemote = agentDescIsRemote.Default.(bool)
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
	// deploymentDescByProfile is the schema descriptor for by_profile field.
	deploymentDescByProfile := deploymentFields[5].Descriptor()
	// deployment.DefaultByProfile holds the default value on creation for the by_profile field.
	deployment.DefaultByProfile = deploymentDescByProfile.Default.(bool)
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
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescName is the schema descriptor for name field.
	profileDescName := profileFields[0].Descriptor()
	// profile.NameValidator is a validator for the "name" field. It is called by the builders before save.
	profile.NameValidator = profileDescName.Validators[0].(func(string) error)
	// profileDescApplyToAll is the schema descriptor for apply_to_all field.
	profileDescApplyToAll := profileFields[1].Descriptor()
	// profile.DefaultApplyToAll holds the default value on creation for the apply_to_all field.
	profile.DefaultApplyToAll = profileDescApplyToAll.Default.(bool)
	profileissueFields := schema.ProfileIssue{}.Fields()
	_ = profileissueFields
	// profileissueDescWhen is the schema descriptor for when field.
	profileissueDescWhen := profileissueFields[1].Descriptor()
	// profileissue.DefaultWhen holds the default value on creation for the when field.
	profileissue.DefaultWhen = profileissueDescWhen.Default.(time.Time)
	// profileissue.UpdateDefaultWhen holds the default value on update for the when field.
	profileissue.UpdateDefaultWhen = profileissueDescWhen.UpdateDefault.(func() time.Time)
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
	// settingsDescRequestVncPin is the schema descriptor for request_vnc_pin field.
	settingsDescRequestVncPin := settingsFields[27].Descriptor()
	// settings.DefaultRequestVncPin holds the default value on creation for the request_vnc_pin field.
	settings.DefaultRequestVncPin = settingsDescRequestVncPin.Default.(bool)
	// settingsDescProfilesApplicationFrequenceInMinutes is the schema descriptor for profiles_application_frequence_in_minutes field.
	settingsDescProfilesApplicationFrequenceInMinutes := settingsFields[28].Descriptor()
	// settings.DefaultProfilesApplicationFrequenceInMinutes holds the default value on creation for the profiles_application_frequence_in_minutes field.
	settings.DefaultProfilesApplicationFrequenceInMinutes = settingsDescProfilesApplicationFrequenceInMinutes.Default.(int)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescTag is the schema descriptor for tag field.
	tagDescTag := tagFields[0].Descriptor()
	// tag.TagValidator is a validator for the "tag" field. It is called by the builders before save.
	tag.TagValidator = tagDescTag.Validators[0].(func(string) error)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescName is the schema descriptor for name field.
	taskDescName := taskFields[0].Descriptor()
	// task.NameValidator is a validator for the "name" field. It is called by the builders before save.
	task.NameValidator = taskDescName.Validators[0].(func(string) error)
	// taskDescPackageID is the schema descriptor for package_id field.
	taskDescPackageID := taskFields[2].Descriptor()
	// task.DefaultPackageID holds the default value on creation for the package_id field.
	task.DefaultPackageID = taskDescPackageID.Default.(string)
	// taskDescPackageName is the schema descriptor for package_name field.
	taskDescPackageName := taskFields[3].Descriptor()
	// task.DefaultPackageName holds the default value on creation for the package_name field.
	task.DefaultPackageName = taskDescPackageName.Default.(string)
	// taskDescRegistryKey is the schema descriptor for registry_key field.
	taskDescRegistryKey := taskFields[4].Descriptor()
	// task.DefaultRegistryKey holds the default value on creation for the registry_key field.
	task.DefaultRegistryKey = taskDescRegistryKey.Default.(string)
	// taskDescRegistryKeyValueName is the schema descriptor for registry_key_value_name field.
	taskDescRegistryKeyValueName := taskFields[5].Descriptor()
	// task.DefaultRegistryKeyValueName holds the default value on creation for the registry_key_value_name field.
	task.DefaultRegistryKeyValueName = taskDescRegistryKeyValueName.Default.(string)
	// taskDescRegistryKeyValueData is the schema descriptor for registry_key_value_data field.
	taskDescRegistryKeyValueData := taskFields[7].Descriptor()
	// task.DefaultRegistryKeyValueData holds the default value on creation for the registry_key_value_data field.
	task.DefaultRegistryKeyValueData = taskDescRegistryKeyValueData.Default.(string)
	// taskDescRegistryHex is the schema descriptor for registry_hex field.
	taskDescRegistryHex := taskFields[8].Descriptor()
	// task.DefaultRegistryHex holds the default value on creation for the registry_hex field.
	task.DefaultRegistryHex = taskDescRegistryHex.Default.(bool)
	// taskDescRegistryForce is the schema descriptor for registry_force field.
	taskDescRegistryForce := taskFields[9].Descriptor()
	// task.DefaultRegistryForce holds the default value on creation for the registry_force field.
	task.DefaultRegistryForce = taskDescRegistryForce.Default.(bool)
	// taskDescLocalUserUsername is the schema descriptor for local_user_username field.
	taskDescLocalUserUsername := taskFields[10].Descriptor()
	// task.DefaultLocalUserUsername holds the default value on creation for the local_user_username field.
	task.DefaultLocalUserUsername = taskDescLocalUserUsername.Default.(string)
	// taskDescLocalUserDescription is the schema descriptor for local_user_description field.
	taskDescLocalUserDescription := taskFields[11].Descriptor()
	// task.DefaultLocalUserDescription holds the default value on creation for the local_user_description field.
	task.DefaultLocalUserDescription = taskDescLocalUserDescription.Default.(string)
	// taskDescLocalUserDisable is the schema descriptor for local_user_disable field.
	taskDescLocalUserDisable := taskFields[12].Descriptor()
	// task.DefaultLocalUserDisable holds the default value on creation for the local_user_disable field.
	task.DefaultLocalUserDisable = taskDescLocalUserDisable.Default.(bool)
	// taskDescLocalUserFullname is the schema descriptor for local_user_fullname field.
	taskDescLocalUserFullname := taskFields[13].Descriptor()
	// task.DefaultLocalUserFullname holds the default value on creation for the local_user_fullname field.
	task.DefaultLocalUserFullname = taskDescLocalUserFullname.Default.(string)
	// taskDescLocalUserPassword is the schema descriptor for local_user_password field.
	taskDescLocalUserPassword := taskFields[14].Descriptor()
	// task.DefaultLocalUserPassword holds the default value on creation for the local_user_password field.
	task.DefaultLocalUserPassword = taskDescLocalUserPassword.Default.(string)
	// taskDescLocalUserPasswordChangeNotAllowed is the schema descriptor for local_user_password_change_not_allowed field.
	taskDescLocalUserPasswordChangeNotAllowed := taskFields[15].Descriptor()
	// task.DefaultLocalUserPasswordChangeNotAllowed holds the default value on creation for the local_user_password_change_not_allowed field.
	task.DefaultLocalUserPasswordChangeNotAllowed = taskDescLocalUserPasswordChangeNotAllowed.Default.(bool)
	// taskDescLocalUserPasswordChangeRequired is the schema descriptor for local_user_password_change_required field.
	taskDescLocalUserPasswordChangeRequired := taskFields[16].Descriptor()
	// task.DefaultLocalUserPasswordChangeRequired holds the default value on creation for the local_user_password_change_required field.
	task.DefaultLocalUserPasswordChangeRequired = taskDescLocalUserPasswordChangeRequired.Default.(bool)
	// taskDescLocalUserPasswordNeverExpires is the schema descriptor for local_user_password_never_expires field.
	taskDescLocalUserPasswordNeverExpires := taskFields[17].Descriptor()
	// task.DefaultLocalUserPasswordNeverExpires holds the default value on creation for the local_user_password_never_expires field.
	task.DefaultLocalUserPasswordNeverExpires = taskDescLocalUserPasswordNeverExpires.Default.(bool)
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
	wingetconfigexclusionFields := schema.WingetConfigExclusion{}.Fields()
	_ = wingetconfigexclusionFields
	// wingetconfigexclusionDescWhen is the schema descriptor for when field.
	wingetconfigexclusionDescWhen := wingetconfigexclusionFields[1].Descriptor()
	// wingetconfigexclusion.DefaultWhen holds the default value on creation for the when field.
	wingetconfigexclusion.DefaultWhen = wingetconfigexclusionDescWhen.Default.(func() time.Time)
}
