// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AlarmType string

// Enum values for AlarmType
const (
	AlarmTypeMetric    AlarmType = "Metric"
	AlarmTypeComposite AlarmType = "Composite"
	AlarmTypeCanary    AlarmType = "Canary"
	AlarmTypeLogs      AlarmType = "Logs"
	AlarmTypeEvent     AlarmType = "Event"
)

// Values returns all known values for AlarmType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AlarmType) Values() []AlarmType {
	return []AlarmType{
		"Metric",
		"Composite",
		"Canary",
		"Logs",
		"Event",
	}
}

type AppAssessmentScheduleType string

// Enum values for AppAssessmentScheduleType
const (
	AppAssessmentScheduleTypeDisabled AppAssessmentScheduleType = "Disabled"
	AppAssessmentScheduleTypeDaily    AppAssessmentScheduleType = "Daily"
)

// Values returns all known values for AppAssessmentScheduleType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppAssessmentScheduleType) Values() []AppAssessmentScheduleType {
	return []AppAssessmentScheduleType{
		"Disabled",
		"Daily",
	}
}

type AppComplianceStatusType string

// Enum values for AppComplianceStatusType
const (
	AppComplianceStatusTypePolicyBreached  AppComplianceStatusType = "PolicyBreached"
	AppComplianceStatusTypePolicyMet       AppComplianceStatusType = "PolicyMet"
	AppComplianceStatusTypeNotAssessed     AppComplianceStatusType = "NotAssessed"
	AppComplianceStatusTypeChangesDetected AppComplianceStatusType = "ChangesDetected"
)

// Values returns all known values for AppComplianceStatusType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AppComplianceStatusType) Values() []AppComplianceStatusType {
	return []AppComplianceStatusType{
		"PolicyBreached",
		"PolicyMet",
		"NotAssessed",
		"ChangesDetected",
	}
}

type AppStatusType string

// Enum values for AppStatusType
const (
	AppStatusTypeActive   AppStatusType = "Active"
	AppStatusTypeDeleting AppStatusType = "Deleting"
)

// Values returns all known values for AppStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AppStatusType) Values() []AppStatusType {
	return []AppStatusType{
		"Active",
		"Deleting",
	}
}

type AssessmentInvoker string

// Enum values for AssessmentInvoker
const (
	AssessmentInvokerUser   AssessmentInvoker = "User"
	AssessmentInvokerSystem AssessmentInvoker = "System"
)

// Values returns all known values for AssessmentInvoker. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssessmentInvoker) Values() []AssessmentInvoker {
	return []AssessmentInvoker{
		"User",
		"System",
	}
}

type AssessmentStatus string

// Enum values for AssessmentStatus
const (
	AssessmentStatusPending    AssessmentStatus = "Pending"
	AssessmentStatusInprogress AssessmentStatus = "InProgress"
	AssessmentStatusFailed     AssessmentStatus = "Failed"
	AssessmentStatusSuccess    AssessmentStatus = "Success"
)

// Values returns all known values for AssessmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssessmentStatus) Values() []AssessmentStatus {
	return []AssessmentStatus{
		"Pending",
		"InProgress",
		"Failed",
		"Success",
	}
}

type ComplianceStatus string

// Enum values for ComplianceStatus
const (
	ComplianceStatusPolicyBreached ComplianceStatus = "PolicyBreached"
	ComplianceStatusPolicyMet      ComplianceStatus = "PolicyMet"
)

// Values returns all known values for ComplianceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComplianceStatus) Values() []ComplianceStatus {
	return []ComplianceStatus{
		"PolicyBreached",
		"PolicyMet",
	}
}

type ConfigRecommendationOptimizationType string

// Enum values for ConfigRecommendationOptimizationType
const (
	ConfigRecommendationOptimizationTypeLeastCost          ConfigRecommendationOptimizationType = "LeastCost"
	ConfigRecommendationOptimizationTypeLeastChange        ConfigRecommendationOptimizationType = "LeastChange"
	ConfigRecommendationOptimizationTypeBestAzRecovery     ConfigRecommendationOptimizationType = "BestAZRecovery"
	ConfigRecommendationOptimizationTypeLeastErrors        ConfigRecommendationOptimizationType = "LeastErrors"
	ConfigRecommendationOptimizationTypeBestAttainable     ConfigRecommendationOptimizationType = "BestAttainable"
	ConfigRecommendationOptimizationTypeBestRegionRecovery ConfigRecommendationOptimizationType = "BestRegionRecovery"
)

// Values returns all known values for ConfigRecommendationOptimizationType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConfigRecommendationOptimizationType) Values() []ConfigRecommendationOptimizationType {
	return []ConfigRecommendationOptimizationType{
		"LeastCost",
		"LeastChange",
		"BestAZRecovery",
		"LeastErrors",
		"BestAttainable",
		"BestRegionRecovery",
	}
}

type CostFrequency string

// Enum values for CostFrequency
const (
	CostFrequencyHourly  CostFrequency = "Hourly"
	CostFrequencyDaily   CostFrequency = "Daily"
	CostFrequencyMonthly CostFrequency = "Monthly"
	CostFrequencyYearly  CostFrequency = "Yearly"
)

// Values returns all known values for CostFrequency. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CostFrequency) Values() []CostFrequency {
	return []CostFrequency{
		"Hourly",
		"Daily",
		"Monthly",
		"Yearly",
	}
}

type DataLocationConstraint string

// Enum values for DataLocationConstraint
const (
	DataLocationConstraintAnyLocation   DataLocationConstraint = "AnyLocation"
	DataLocationConstraintSameContinent DataLocationConstraint = "SameContinent"
	DataLocationConstraintSameCountry   DataLocationConstraint = "SameCountry"
)

// Values returns all known values for DataLocationConstraint. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataLocationConstraint) Values() []DataLocationConstraint {
	return []DataLocationConstraint{
		"AnyLocation",
		"SameContinent",
		"SameCountry",
	}
}

type DisruptionType string

// Enum values for DisruptionType
const (
	DisruptionTypeSoftware DisruptionType = "Software"
	DisruptionTypeHardware DisruptionType = "Hardware"
	DisruptionTypeAz       DisruptionType = "AZ"
	DisruptionTypeRegion   DisruptionType = "Region"
)

// Values returns all known values for DisruptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DisruptionType) Values() []DisruptionType {
	return []DisruptionType{
		"Software",
		"Hardware",
		"AZ",
		"Region",
	}
}

type EstimatedCostTier string

// Enum values for EstimatedCostTier
const (
	EstimatedCostTierL1 EstimatedCostTier = "L1"
	EstimatedCostTierL2 EstimatedCostTier = "L2"
	EstimatedCostTierL3 EstimatedCostTier = "L3"
	EstimatedCostTierL4 EstimatedCostTier = "L4"
)

// Values returns all known values for EstimatedCostTier. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EstimatedCostTier) Values() []EstimatedCostTier {
	return []EstimatedCostTier{
		"L1",
		"L2",
		"L3",
		"L4",
	}
}

type HaArchitecture string

// Enum values for HaArchitecture
const (
	HaArchitectureMultiSite        HaArchitecture = "MultiSite"
	HaArchitectureWarmStandby      HaArchitecture = "WarmStandby"
	HaArchitecturePilotLight       HaArchitecture = "PilotLight"
	HaArchitectureBackupAndRestore HaArchitecture = "BackupAndRestore"
	HaArchitectureNoRecoveryPlan   HaArchitecture = "NoRecoveryPlan"
)

// Values returns all known values for HaArchitecture. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HaArchitecture) Values() []HaArchitecture {
	return []HaArchitecture{
		"MultiSite",
		"WarmStandby",
		"PilotLight",
		"BackupAndRestore",
		"NoRecoveryPlan",
	}
}

type PhysicalIdentifierType string

// Enum values for PhysicalIdentifierType
const (
	PhysicalIdentifierTypeArn    PhysicalIdentifierType = "Arn"
	PhysicalIdentifierTypeNative PhysicalIdentifierType = "Native"
)

// Values returns all known values for PhysicalIdentifierType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhysicalIdentifierType) Values() []PhysicalIdentifierType {
	return []PhysicalIdentifierType{
		"Arn",
		"Native",
	}
}

type RecommendationComplianceStatus string

// Enum values for RecommendationComplianceStatus
const (
	RecommendationComplianceStatusBreachedUnattainable RecommendationComplianceStatus = "BreachedUnattainable"
	RecommendationComplianceStatusBreachedCanMeet      RecommendationComplianceStatus = "BreachedCanMeet"
	RecommendationComplianceStatusMetCanImprove        RecommendationComplianceStatus = "MetCanImprove"
)

// Values returns all known values for RecommendationComplianceStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RecommendationComplianceStatus) Values() []RecommendationComplianceStatus {
	return []RecommendationComplianceStatus{
		"BreachedUnattainable",
		"BreachedCanMeet",
		"MetCanImprove",
	}
}

type RecommendationTemplateStatus string

// Enum values for RecommendationTemplateStatus
const (
	RecommendationTemplateStatusPending    RecommendationTemplateStatus = "Pending"
	RecommendationTemplateStatusInProgress RecommendationTemplateStatus = "InProgress"
	RecommendationTemplateStatusFailed     RecommendationTemplateStatus = "Failed"
	RecommendationTemplateStatusSuccess    RecommendationTemplateStatus = "Success"
)

// Values returns all known values for RecommendationTemplateStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationTemplateStatus) Values() []RecommendationTemplateStatus {
	return []RecommendationTemplateStatus{
		"Pending",
		"InProgress",
		"Failed",
		"Success",
	}
}

type RenderRecommendationType string

// Enum values for RenderRecommendationType
const (
	RenderRecommendationTypeAlarm RenderRecommendationType = "Alarm"
	RenderRecommendationTypeSop   RenderRecommendationType = "Sop"
	RenderRecommendationTypeTest  RenderRecommendationType = "Test"
)

// Values returns all known values for RenderRecommendationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RenderRecommendationType) Values() []RenderRecommendationType {
	return []RenderRecommendationType{
		"Alarm",
		"Sop",
		"Test",
	}
}

type ResiliencyPolicyTier string

// Enum values for ResiliencyPolicyTier
const (
	ResiliencyPolicyTierMissionCritical ResiliencyPolicyTier = "MissionCritical"
	ResiliencyPolicyTierCritical        ResiliencyPolicyTier = "Critical"
	ResiliencyPolicyTierImportant       ResiliencyPolicyTier = "Important"
	ResiliencyPolicyTierCoreServices    ResiliencyPolicyTier = "CoreServices"
	ResiliencyPolicyTierNonCritical     ResiliencyPolicyTier = "NonCritical"
)

// Values returns all known values for ResiliencyPolicyTier. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResiliencyPolicyTier) Values() []ResiliencyPolicyTier {
	return []ResiliencyPolicyTier{
		"MissionCritical",
		"Critical",
		"Important",
		"CoreServices",
		"NonCritical",
	}
}

type ResourceImportStatusType string

// Enum values for ResourceImportStatusType
const (
	ResourceImportStatusTypePending    ResourceImportStatusType = "Pending"
	ResourceImportStatusTypeInProgress ResourceImportStatusType = "InProgress"
	ResourceImportStatusTypeFailed     ResourceImportStatusType = "Failed"
	ResourceImportStatusTypeSuccess    ResourceImportStatusType = "Success"
)

// Values returns all known values for ResourceImportStatusType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceImportStatusType) Values() []ResourceImportStatusType {
	return []ResourceImportStatusType{
		"Pending",
		"InProgress",
		"Failed",
		"Success",
	}
}

type ResourceMappingType string

// Enum values for ResourceMappingType
const (
	ResourceMappingTypeCfnStack       ResourceMappingType = "CfnStack"
	ResourceMappingTypeResource       ResourceMappingType = "Resource"
	ResourceMappingTypeAppRegistryApp ResourceMappingType = "AppRegistryApp"
	ResourceMappingTypeResourceGroup  ResourceMappingType = "ResourceGroup"
	ResourceMappingTypeTerraform      ResourceMappingType = "Terraform"
)

// Values returns all known values for ResourceMappingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceMappingType) Values() []ResourceMappingType {
	return []ResourceMappingType{
		"CfnStack",
		"Resource",
		"AppRegistryApp",
		"ResourceGroup",
		"Terraform",
	}
}

type ResourceResolutionStatusType string

// Enum values for ResourceResolutionStatusType
const (
	ResourceResolutionStatusTypePending    ResourceResolutionStatusType = "Pending"
	ResourceResolutionStatusTypeInProgress ResourceResolutionStatusType = "InProgress"
	ResourceResolutionStatusTypeFailed     ResourceResolutionStatusType = "Failed"
	ResourceResolutionStatusTypeSuccess    ResourceResolutionStatusType = "Success"
)

// Values returns all known values for ResourceResolutionStatusType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceResolutionStatusType) Values() []ResourceResolutionStatusType {
	return []ResourceResolutionStatusType{
		"Pending",
		"InProgress",
		"Failed",
		"Success",
	}
}

type SopServiceType string

// Enum values for SopServiceType
const (
	SopServiceTypeSsm SopServiceType = "SSM"
)

// Values returns all known values for SopServiceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SopServiceType) Values() []SopServiceType {
	return []SopServiceType{
		"SSM",
	}
}

type TemplateFormat string

// Enum values for TemplateFormat
const (
	TemplateFormatCfnYaml TemplateFormat = "CfnYaml"
	TemplateFormatCfnJson TemplateFormat = "CfnJson"
)

// Values returns all known values for TemplateFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TemplateFormat) Values() []TemplateFormat {
	return []TemplateFormat{
		"CfnYaml",
		"CfnJson",
	}
}

type TestRisk string

// Enum values for TestRisk
const (
	TestRiskSmall  TestRisk = "Small"
	TestRiskMedium TestRisk = "Medium"
	TestRiskHigh   TestRisk = "High"
)

// Values returns all known values for TestRisk. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TestRisk) Values() []TestRisk {
	return []TestRisk{
		"Small",
		"Medium",
		"High",
	}
}

type TestType string

// Enum values for TestType
const (
	TestTypeSoftware TestType = "Software"
	TestTypeHardware TestType = "Hardware"
	TestTypeAz       TestType = "AZ"
	TestTypeRegion   TestType = "Region"
)

// Values returns all known values for TestType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TestType) Values() []TestType {
	return []TestType{
		"Software",
		"Hardware",
		"AZ",
		"Region",
	}
}
