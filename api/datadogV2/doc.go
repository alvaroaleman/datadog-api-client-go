// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

// List of APIs:
//   - [APIManagementApi.CreateOpenAPI]
//   - [APIManagementApi.DeleteOpenAPI]
//   - [APIManagementApi.GetOpenAPI]
//   - [APIManagementApi.UpdateOpenAPI]
//   - [APMRetentionFiltersApi.CreateApmRetentionFilter]
//   - [APMRetentionFiltersApi.DeleteApmRetentionFilter]
//   - [APMRetentionFiltersApi.GetApmRetentionFilter]
//   - [APMRetentionFiltersApi.ListApmRetentionFilters]
//   - [APMRetentionFiltersApi.ReorderApmRetentionFilters]
//   - [APMRetentionFiltersApi.UpdateApmRetentionFilter]
//   - [AuditApi.ListAuditLogs]
//   - [AuditApi.SearchAuditLogs]
//   - [AuthNMappingsApi.CreateAuthNMapping]
//   - [AuthNMappingsApi.DeleteAuthNMapping]
//   - [AuthNMappingsApi.GetAuthNMapping]
//   - [AuthNMappingsApi.ListAuthNMappings]
//   - [AuthNMappingsApi.UpdateAuthNMapping]
//   - [CIVisibilityPipelinesApi.AggregateCIAppPipelineEvents]
//   - [CIVisibilityPipelinesApi.CreateCIAppPipelineEvent]
//   - [CIVisibilityPipelinesApi.ListCIAppPipelineEvents]
//   - [CIVisibilityPipelinesApi.SearchCIAppPipelineEvents]
//   - [CIVisibilityTestsApi.AggregateCIAppTestEvents]
//   - [CIVisibilityTestsApi.ListCIAppTestEvents]
//   - [CIVisibilityTestsApi.SearchCIAppTestEvents]
//   - [CaseManagementApi.ArchiveCase]
//   - [CaseManagementApi.AssignCase]
//   - [CaseManagementApi.CreateCase]
//   - [CaseManagementApi.CreateProject]
//   - [CaseManagementApi.DeleteProject]
//   - [CaseManagementApi.GetCase]
//   - [CaseManagementApi.GetProject]
//   - [CaseManagementApi.GetProjects]
//   - [CaseManagementApi.SearchCases]
//   - [CaseManagementApi.UnarchiveCase]
//   - [CaseManagementApi.UnassignCase]
//   - [CaseManagementApi.UpdatePriority]
//   - [CaseManagementApi.UpdateStatus]
//   - [CloudCostManagementApi.CreateCostAWSCURConfig]
//   - [CloudCostManagementApi.CreateCostAzureUCConfigs]
//   - [CloudCostManagementApi.DeleteCostAWSCURConfig]
//   - [CloudCostManagementApi.DeleteCostAzureUCConfig]
//   - [CloudCostManagementApi.GetCloudCostActivity]
//   - [CloudCostManagementApi.ListAWSRelatedAccounts]
//   - [CloudCostManagementApi.ListCostAWSCURConfigs]
//   - [CloudCostManagementApi.ListCostAzureUCConfigs]
//   - [CloudCostManagementApi.UpdateCostAWSCURConfig]
//   - [CloudCostManagementApi.UpdateCostAzureUCConfigs]
//   - [CloudWorkloadSecurityApi.CreateCSMThreatsAgentRule]
//   - [CloudWorkloadSecurityApi.CreateCloudWorkloadSecurityAgentRule]
//   - [CloudWorkloadSecurityApi.DeleteCSMThreatsAgentRule]
//   - [CloudWorkloadSecurityApi.DeleteCloudWorkloadSecurityAgentRule]
//   - [CloudWorkloadSecurityApi.DownloadCSMThreatsPolicy]
//   - [CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile]
//   - [CloudWorkloadSecurityApi.GetCSMThreatsAgentRule]
//   - [CloudWorkloadSecurityApi.GetCloudWorkloadSecurityAgentRule]
//   - [CloudWorkloadSecurityApi.ListCSMThreatsAgentRules]
//   - [CloudWorkloadSecurityApi.ListCloudWorkloadSecurityAgentRules]
//   - [CloudWorkloadSecurityApi.UpdateCSMThreatsAgentRule]
//   - [CloudWorkloadSecurityApi.UpdateCloudWorkloadSecurityAgentRule]
//   - [CloudflareIntegrationApi.CreateCloudflareAccount]
//   - [CloudflareIntegrationApi.DeleteCloudflareAccount]
//   - [CloudflareIntegrationApi.GetCloudflareAccount]
//   - [CloudflareIntegrationApi.ListCloudflareAccounts]
//   - [CloudflareIntegrationApi.UpdateCloudflareAccount]
//   - [ConfluentCloudApi.CreateConfluentAccount]
//   - [ConfluentCloudApi.CreateConfluentResource]
//   - [ConfluentCloudApi.DeleteConfluentAccount]
//   - [ConfluentCloudApi.DeleteConfluentResource]
//   - [ConfluentCloudApi.GetConfluentAccount]
//   - [ConfluentCloudApi.GetConfluentResource]
//   - [ConfluentCloudApi.ListConfluentAccount]
//   - [ConfluentCloudApi.ListConfluentResource]
//   - [ConfluentCloudApi.UpdateConfluentAccount]
//   - [ConfluentCloudApi.UpdateConfluentResource]
//   - [ContainerImagesApi.ListContainerImages]
//   - [ContainersApi.ListContainers]
//   - [DORAMetricsApi.CreateDORADeployment]
//   - [DORAMetricsApi.CreateDORAIncident]
//   - [DashboardListsApi.CreateDashboardListItems]
//   - [DashboardListsApi.DeleteDashboardListItems]
//   - [DashboardListsApi.GetDashboardListItems]
//   - [DashboardListsApi.UpdateDashboardListItems]
//   - [DowntimesApi.CancelDowntime]
//   - [DowntimesApi.CreateDowntime]
//   - [DowntimesApi.GetDowntime]
//   - [DowntimesApi.ListDowntimes]
//   - [DowntimesApi.ListMonitorDowntimes]
//   - [DowntimesApi.UpdateDowntime]
//   - [EventsApi.ListEvents]
//   - [EventsApi.SearchEvents]
//   - [FastlyIntegrationApi.CreateFastlyAccount]
//   - [FastlyIntegrationApi.CreateFastlyService]
//   - [FastlyIntegrationApi.DeleteFastlyAccount]
//   - [FastlyIntegrationApi.DeleteFastlyService]
//   - [FastlyIntegrationApi.GetFastlyAccount]
//   - [FastlyIntegrationApi.GetFastlyService]
//   - [FastlyIntegrationApi.ListFastlyAccounts]
//   - [FastlyIntegrationApi.ListFastlyServices]
//   - [FastlyIntegrationApi.UpdateFastlyAccount]
//   - [FastlyIntegrationApi.UpdateFastlyService]
//   - [GCPIntegrationApi.CreateGCPSTSAccount]
//   - [GCPIntegrationApi.DeleteGCPSTSAccount]
//   - [GCPIntegrationApi.GetGCPSTSDelegate]
//   - [GCPIntegrationApi.ListGCPSTSAccounts]
//   - [GCPIntegrationApi.MakeGCPSTSDelegate]
//   - [GCPIntegrationApi.UpdateGCPSTSAccount]
//   - [IPAllowlistApi.GetIPAllowlist]
//   - [IPAllowlistApi.UpdateIPAllowlist]
//   - [IncidentServicesApi.CreateIncidentService]
//   - [IncidentServicesApi.DeleteIncidentService]
//   - [IncidentServicesApi.GetIncidentService]
//   - [IncidentServicesApi.ListIncidentServices]
//   - [IncidentServicesApi.UpdateIncidentService]
//   - [IncidentTeamsApi.CreateIncidentTeam]
//   - [IncidentTeamsApi.DeleteIncidentTeam]
//   - [IncidentTeamsApi.GetIncidentTeam]
//   - [IncidentTeamsApi.ListIncidentTeams]
//   - [IncidentTeamsApi.UpdateIncidentTeam]
//   - [IncidentsApi.CreateIncident]
//   - [IncidentsApi.CreateIncidentIntegration]
//   - [IncidentsApi.CreateIncidentTodo]
//   - [IncidentsApi.DeleteIncident]
//   - [IncidentsApi.DeleteIncidentIntegration]
//   - [IncidentsApi.DeleteIncidentTodo]
//   - [IncidentsApi.GetIncident]
//   - [IncidentsApi.GetIncidentIntegration]
//   - [IncidentsApi.GetIncidentTodo]
//   - [IncidentsApi.ListIncidentAttachments]
//   - [IncidentsApi.ListIncidentIntegrations]
//   - [IncidentsApi.ListIncidentTodos]
//   - [IncidentsApi.ListIncidents]
//   - [IncidentsApi.SearchIncidents]
//   - [IncidentsApi.UpdateIncident]
//   - [IncidentsApi.UpdateIncidentAttachments]
//   - [IncidentsApi.UpdateIncidentIntegration]
//   - [IncidentsApi.UpdateIncidentTodo]
//   - [KeyManagementApi.CreateAPIKey]
//   - [KeyManagementApi.CreateCurrentUserApplicationKey]
//   - [KeyManagementApi.DeleteAPIKey]
//   - [KeyManagementApi.DeleteApplicationKey]
//   - [KeyManagementApi.DeleteCurrentUserApplicationKey]
//   - [KeyManagementApi.GetAPIKey]
//   - [KeyManagementApi.GetApplicationKey]
//   - [KeyManagementApi.GetCurrentUserApplicationKey]
//   - [KeyManagementApi.ListAPIKeys]
//   - [KeyManagementApi.ListApplicationKeys]
//   - [KeyManagementApi.ListCurrentUserApplicationKeys]
//   - [KeyManagementApi.UpdateAPIKey]
//   - [KeyManagementApi.UpdateApplicationKey]
//   - [KeyManagementApi.UpdateCurrentUserApplicationKey]
//   - [LogsApi.AggregateLogs]
//   - [LogsApi.ListLogs]
//   - [LogsApi.ListLogsGet]
//   - [LogsApi.SubmitLog]
//   - [LogsArchivesApi.AddReadRoleToArchive]
//   - [LogsArchivesApi.CreateLogsArchive]
//   - [LogsArchivesApi.DeleteLogsArchive]
//   - [LogsArchivesApi.GetLogsArchive]
//   - [LogsArchivesApi.GetLogsArchiveOrder]
//   - [LogsArchivesApi.ListArchiveReadRoles]
//   - [LogsArchivesApi.ListLogsArchives]
//   - [LogsArchivesApi.RemoveRoleFromArchive]
//   - [LogsArchivesApi.UpdateLogsArchive]
//   - [LogsArchivesApi.UpdateLogsArchiveOrder]
//   - [LogsMetricsApi.CreateLogsMetric]
//   - [LogsMetricsApi.DeleteLogsMetric]
//   - [LogsMetricsApi.GetLogsMetric]
//   - [LogsMetricsApi.ListLogsMetrics]
//   - [LogsMetricsApi.UpdateLogsMetric]
//   - [MetricsApi.CreateBulkTagsMetricsConfiguration]
//   - [MetricsApi.CreateTagConfiguration]
//   - [MetricsApi.DeleteBulkTagsMetricsConfiguration]
//   - [MetricsApi.DeleteTagConfiguration]
//   - [MetricsApi.EstimateMetricsOutputSeries]
//   - [MetricsApi.ListActiveMetricConfigurations]
//   - [MetricsApi.ListMetricAssets]
//   - [MetricsApi.ListTagConfigurationByName]
//   - [MetricsApi.ListTagConfigurations]
//   - [MetricsApi.ListTagsByMetricName]
//   - [MetricsApi.ListVolumesByMetricName]
//   - [MetricsApi.QueryScalarData]
//   - [MetricsApi.QueryTimeseriesData]
//   - [MetricsApi.SubmitMetrics]
//   - [MetricsApi.UpdateTagConfiguration]
//   - [MonitorsApi.CreateMonitorConfigPolicy]
//   - [MonitorsApi.DeleteMonitorConfigPolicy]
//   - [MonitorsApi.GetMonitorConfigPolicy]
//   - [MonitorsApi.ListMonitorConfigPolicies]
//   - [MonitorsApi.UpdateMonitorConfigPolicy]
//   - [OktaIntegrationApi.CreateOktaAccount]
//   - [OktaIntegrationApi.DeleteOktaAccount]
//   - [OktaIntegrationApi.GetOktaAccount]
//   - [OktaIntegrationApi.ListOktaAccounts]
//   - [OktaIntegrationApi.UpdateOktaAccount]
//   - [OpsgenieIntegrationApi.CreateOpsgenieService]
//   - [OpsgenieIntegrationApi.DeleteOpsgenieService]
//   - [OpsgenieIntegrationApi.GetOpsgenieService]
//   - [OpsgenieIntegrationApi.ListOpsgenieServices]
//   - [OpsgenieIntegrationApi.UpdateOpsgenieService]
//   - [OrganizationsApi.UploadIdPMetadata]
//   - [PowerpackApi.CreatePowerpack]
//   - [PowerpackApi.DeletePowerpack]
//   - [PowerpackApi.GetPowerpack]
//   - [PowerpackApi.ListPowerpacks]
//   - [PowerpackApi.UpdatePowerpack]
//   - [ProcessesApi.ListProcesses]
//   - [RUMApi.AggregateRUMEvents]
//   - [RUMApi.CreateRUMApplication]
//   - [RUMApi.DeleteRUMApplication]
//   - [RUMApi.GetRUMApplication]
//   - [RUMApi.GetRUMApplications]
//   - [RUMApi.ListRUMEvents]
//   - [RUMApi.SearchRUMEvents]
//   - [RUMApi.UpdateRUMApplication]
//   - [RestrictionPoliciesApi.DeleteRestrictionPolicy]
//   - [RestrictionPoliciesApi.GetRestrictionPolicy]
//   - [RestrictionPoliciesApi.UpdateRestrictionPolicy]
//   - [RolesApi.AddPermissionToRole]
//   - [RolesApi.AddUserToRole]
//   - [RolesApi.CloneRole]
//   - [RolesApi.CreateRole]
//   - [RolesApi.DeleteRole]
//   - [RolesApi.GetRole]
//   - [RolesApi.ListPermissions]
//   - [RolesApi.ListRolePermissions]
//   - [RolesApi.ListRoleUsers]
//   - [RolesApi.ListRoles]
//   - [RolesApi.RemovePermissionFromRole]
//   - [RolesApi.RemoveUserFromRole]
//   - [RolesApi.UpdateRole]
//   - [SecurityMonitoringApi.CreateSecurityFilter]
//   - [SecurityMonitoringApi.CreateSecurityMonitoringRule]
//   - [SecurityMonitoringApi.CreateSecurityMonitoringSuppression]
//   - [SecurityMonitoringApi.DeleteSecurityFilter]
//   - [SecurityMonitoringApi.DeleteSecurityMonitoringRule]
//   - [SecurityMonitoringApi.DeleteSecurityMonitoringSuppression]
//   - [SecurityMonitoringApi.EditSecurityMonitoringSignalAssignee]
//   - [SecurityMonitoringApi.EditSecurityMonitoringSignalIncidents]
//   - [SecurityMonitoringApi.EditSecurityMonitoringSignalState]
//   - [SecurityMonitoringApi.GetFinding]
//   - [SecurityMonitoringApi.GetSecurityFilter]
//   - [SecurityMonitoringApi.GetSecurityMonitoringRule]
//   - [SecurityMonitoringApi.GetSecurityMonitoringSignal]
//   - [SecurityMonitoringApi.GetSecurityMonitoringSuppression]
//   - [SecurityMonitoringApi.ListFindings]
//   - [SecurityMonitoringApi.ListSecurityFilters]
//   - [SecurityMonitoringApi.ListSecurityMonitoringRules]
//   - [SecurityMonitoringApi.ListSecurityMonitoringSignals]
//   - [SecurityMonitoringApi.ListSecurityMonitoringSuppressions]
//   - [SecurityMonitoringApi.MuteFindings]
//   - [SecurityMonitoringApi.SearchSecurityMonitoringSignals]
//   - [SecurityMonitoringApi.UpdateSecurityFilter]
//   - [SecurityMonitoringApi.UpdateSecurityMonitoringRule]
//   - [SecurityMonitoringApi.UpdateSecurityMonitoringSuppression]
//   - [SensitiveDataScannerApi.CreateScanningGroup]
//   - [SensitiveDataScannerApi.CreateScanningRule]
//   - [SensitiveDataScannerApi.DeleteScanningGroup]
//   - [SensitiveDataScannerApi.DeleteScanningRule]
//   - [SensitiveDataScannerApi.ListScanningGroups]
//   - [SensitiveDataScannerApi.ListStandardPatterns]
//   - [SensitiveDataScannerApi.ReorderScanningGroups]
//   - [SensitiveDataScannerApi.UpdateScanningGroup]
//   - [SensitiveDataScannerApi.UpdateScanningRule]
//   - [ServiceAccountsApi.CreateServiceAccount]
//   - [ServiceAccountsApi.CreateServiceAccountApplicationKey]
//   - [ServiceAccountsApi.DeleteServiceAccountApplicationKey]
//   - [ServiceAccountsApi.GetServiceAccountApplicationKey]
//   - [ServiceAccountsApi.ListServiceAccountApplicationKeys]
//   - [ServiceAccountsApi.UpdateServiceAccountApplicationKey]
//   - [ServiceDefinitionApi.CreateOrUpdateServiceDefinitions]
//   - [ServiceDefinitionApi.DeleteServiceDefinition]
//   - [ServiceDefinitionApi.GetServiceDefinition]
//   - [ServiceDefinitionApi.ListServiceDefinitions]
//   - [ServiceScorecardsApi.CreateScorecardOutcomesBatch]
//   - [ServiceScorecardsApi.CreateScorecardRule]
//   - [ServiceScorecardsApi.DeleteScorecardRule]
//   - [ServiceScorecardsApi.ListScorecardOutcomes]
//   - [ServiceScorecardsApi.ListScorecardRules]
//   - [SpansApi.AggregateSpans]
//   - [SpansApi.ListSpans]
//   - [SpansApi.ListSpansGet]
//   - [SpansMetricsApi.CreateSpansMetric]
//   - [SpansMetricsApi.DeleteSpansMetric]
//   - [SpansMetricsApi.GetSpansMetric]
//   - [SpansMetricsApi.ListSpansMetrics]
//   - [SpansMetricsApi.UpdateSpansMetric]
//   - [SyntheticsApi.GetOnDemandConcurrencyCap]
//   - [SyntheticsApi.SetOnDemandConcurrencyCap]
//   - [TeamsApi.CreateTeam]
//   - [TeamsApi.CreateTeamLink]
//   - [TeamsApi.CreateTeamMembership]
//   - [TeamsApi.DeleteTeam]
//   - [TeamsApi.DeleteTeamLink]
//   - [TeamsApi.DeleteTeamMembership]
//   - [TeamsApi.GetTeam]
//   - [TeamsApi.GetTeamLink]
//   - [TeamsApi.GetTeamLinks]
//   - [TeamsApi.GetTeamMemberships]
//   - [TeamsApi.GetTeamPermissionSettings]
//   - [TeamsApi.GetUserMemberships]
//   - [TeamsApi.ListTeams]
//   - [TeamsApi.UpdateTeam]
//   - [TeamsApi.UpdateTeamLink]
//   - [TeamsApi.UpdateTeamMembership]
//   - [TeamsApi.UpdateTeamPermissionSetting]
//   - [UsageMeteringApi.GetActiveBillingDimensions]
//   - [UsageMeteringApi.GetCostByOrg]
//   - [UsageMeteringApi.GetEstimatedCostByOrg]
//   - [UsageMeteringApi.GetHistoricalCostByOrg]
//   - [UsageMeteringApi.GetHourlyUsage]
//   - [UsageMeteringApi.GetMonthlyCostAttribution]
//   - [UsageMeteringApi.GetProjectedCost]
//   - [UsageMeteringApi.GetUsageApplicationSecurityMonitoring]
//   - [UsageMeteringApi.GetUsageLambdaTracedInvocations]
//   - [UsageMeteringApi.GetUsageObservabilityPipelines]
//   - [UsersApi.CreateUser]
//   - [UsersApi.DisableUser]
//   - [UsersApi.GetInvitation]
//   - [UsersApi.GetUser]
//   - [UsersApi.ListUserOrganizations]
//   - [UsersApi.ListUserPermissions]
//   - [UsersApi.ListUsers]
//   - [UsersApi.SendInvitations]
//   - [UsersApi.UpdateUser]
package datadogV2
