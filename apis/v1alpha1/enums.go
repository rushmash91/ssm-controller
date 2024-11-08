// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AssociationComplianceSeverity string

const (
	AssociationComplianceSeverity_CRITICAL    AssociationComplianceSeverity = "CRITICAL"
	AssociationComplianceSeverity_HIGH        AssociationComplianceSeverity = "HIGH"
	AssociationComplianceSeverity_MEDIUM      AssociationComplianceSeverity = "MEDIUM"
	AssociationComplianceSeverity_LOW         AssociationComplianceSeverity = "LOW"
	AssociationComplianceSeverity_UNSPECIFIED AssociationComplianceSeverity = "UNSPECIFIED"
)

type AssociationExecutionFilterKey string

const (
	AssociationExecutionFilterKey_ExecutionId AssociationExecutionFilterKey = "ExecutionId"
	AssociationExecutionFilterKey_Status      AssociationExecutionFilterKey = "Status"
	AssociationExecutionFilterKey_CreatedTime AssociationExecutionFilterKey = "CreatedTime"
)

type AssociationExecutionTargetsFilterKey string

const (
	AssociationExecutionTargetsFilterKey_Status       AssociationExecutionTargetsFilterKey = "Status"
	AssociationExecutionTargetsFilterKey_ResourceId   AssociationExecutionTargetsFilterKey = "ResourceId"
	AssociationExecutionTargetsFilterKey_ResourceType AssociationExecutionTargetsFilterKey = "ResourceType"
)

type AssociationFilterKey string

const (
	AssociationFilterKey_InstanceId            AssociationFilterKey = "InstanceId"
	AssociationFilterKey_Name                  AssociationFilterKey = "Name"
	AssociationFilterKey_AssociationId         AssociationFilterKey = "AssociationId"
	AssociationFilterKey_AssociationStatusName AssociationFilterKey = "AssociationStatusName"
	AssociationFilterKey_LastExecutedBefore    AssociationFilterKey = "LastExecutedBefore"
	AssociationFilterKey_LastExecutedAfter     AssociationFilterKey = "LastExecutedAfter"
	AssociationFilterKey_AssociationName       AssociationFilterKey = "AssociationName"
	AssociationFilterKey_ResourceGroupName     AssociationFilterKey = "ResourceGroupName"
)

type AssociationFilterOperatorType string

const (
	AssociationFilterOperatorType_EQUAL        AssociationFilterOperatorType = "EQUAL"
	AssociationFilterOperatorType_LESS_THAN    AssociationFilterOperatorType = "LESS_THAN"
	AssociationFilterOperatorType_GREATER_THAN AssociationFilterOperatorType = "GREATER_THAN"
)

type AssociationStatusName string

const (
	AssociationStatusName_Pending AssociationStatusName = "Pending"
	AssociationStatusName_Success AssociationStatusName = "Success"
	AssociationStatusName_Failed  AssociationStatusName = "Failed"
)

type AssociationSyncCompliance string

const (
	AssociationSyncCompliance_AUTO   AssociationSyncCompliance = "AUTO"
	AssociationSyncCompliance_MANUAL AssociationSyncCompliance = "MANUAL"
)

type AttachmentHashType string

const (
	AttachmentHashType_Sha256 AttachmentHashType = "Sha256"
)

type AttachmentsSourceKey string

const (
	AttachmentsSourceKey_SourceUrl           AttachmentsSourceKey = "SourceUrl"
	AttachmentsSourceKey_S3FileUrl           AttachmentsSourceKey = "S3FileUrl"
	AttachmentsSourceKey_AttachmentReference AttachmentsSourceKey = "AttachmentReference"
)

type AutomationExecutionFilterKey string

const (
	AutomationExecutionFilterKey_DocumentNamePrefix  AutomationExecutionFilterKey = "DocumentNamePrefix"
	AutomationExecutionFilterKey_ExecutionStatus     AutomationExecutionFilterKey = "ExecutionStatus"
	AutomationExecutionFilterKey_ExecutionId         AutomationExecutionFilterKey = "ExecutionId"
	AutomationExecutionFilterKey_ParentExecutionId   AutomationExecutionFilterKey = "ParentExecutionId"
	AutomationExecutionFilterKey_CurrentAction       AutomationExecutionFilterKey = "CurrentAction"
	AutomationExecutionFilterKey_StartTimeBefore     AutomationExecutionFilterKey = "StartTimeBefore"
	AutomationExecutionFilterKey_StartTimeAfter      AutomationExecutionFilterKey = "StartTimeAfter"
	AutomationExecutionFilterKey_AutomationType      AutomationExecutionFilterKey = "AutomationType"
	AutomationExecutionFilterKey_TagKey              AutomationExecutionFilterKey = "TagKey"
	AutomationExecutionFilterKey_TargetResourceGroup AutomationExecutionFilterKey = "TargetResourceGroup"
	AutomationExecutionFilterKey_AutomationSubtype   AutomationExecutionFilterKey = "AutomationSubtype"
	AutomationExecutionFilterKey_OpsItemId           AutomationExecutionFilterKey = "OpsItemId"
)

type AutomationExecutionStatus string

const (
	AutomationExecutionStatus_Pending                        AutomationExecutionStatus = "Pending"
	AutomationExecutionStatus_InProgress                     AutomationExecutionStatus = "InProgress"
	AutomationExecutionStatus_Waiting                        AutomationExecutionStatus = "Waiting"
	AutomationExecutionStatus_Success                        AutomationExecutionStatus = "Success"
	AutomationExecutionStatus_TimedOut                       AutomationExecutionStatus = "TimedOut"
	AutomationExecutionStatus_Cancelling                     AutomationExecutionStatus = "Cancelling"
	AutomationExecutionStatus_Cancelled                      AutomationExecutionStatus = "Cancelled"
	AutomationExecutionStatus_Failed                         AutomationExecutionStatus = "Failed"
	AutomationExecutionStatus_PendingApproval                AutomationExecutionStatus = "PendingApproval"
	AutomationExecutionStatus_Approved                       AutomationExecutionStatus = "Approved"
	AutomationExecutionStatus_Rejected                       AutomationExecutionStatus = "Rejected"
	AutomationExecutionStatus_Scheduled                      AutomationExecutionStatus = "Scheduled"
	AutomationExecutionStatus_RunbookInProgress              AutomationExecutionStatus = "RunbookInProgress"
	AutomationExecutionStatus_PendingChangeCalendarOverride  AutomationExecutionStatus = "PendingChangeCalendarOverride"
	AutomationExecutionStatus_ChangeCalendarOverrideApproved AutomationExecutionStatus = "ChangeCalendarOverrideApproved"
	AutomationExecutionStatus_ChangeCalendarOverrideRejected AutomationExecutionStatus = "ChangeCalendarOverrideRejected"
	AutomationExecutionStatus_CompletedWithSuccess           AutomationExecutionStatus = "CompletedWithSuccess"
	AutomationExecutionStatus_CompletedWithFailure           AutomationExecutionStatus = "CompletedWithFailure"
	AutomationExecutionStatus_Exited                         AutomationExecutionStatus = "Exited"
)

type AutomationSubtype string

const (
	AutomationSubtype_ChangeRequest AutomationSubtype = "ChangeRequest"
)

type AutomationType string

const (
	AutomationType_CrossAccount AutomationType = "CrossAccount"
	AutomationType_Local        AutomationType = "Local"
)

type CalendarState string

const (
	CalendarState_OPEN   CalendarState = "OPEN"
	CalendarState_CLOSED CalendarState = "CLOSED"
)

type CommandFilterKey string

const (
	CommandFilterKey_InvokedAfter   CommandFilterKey = "InvokedAfter"
	CommandFilterKey_InvokedBefore  CommandFilterKey = "InvokedBefore"
	CommandFilterKey_Status         CommandFilterKey = "Status"
	CommandFilterKey_ExecutionStage CommandFilterKey = "ExecutionStage"
	CommandFilterKey_DocumentName   CommandFilterKey = "DocumentName"
)

type CommandInvocationStatus string

const (
	CommandInvocationStatus_Pending    CommandInvocationStatus = "Pending"
	CommandInvocationStatus_InProgress CommandInvocationStatus = "InProgress"
	CommandInvocationStatus_Delayed    CommandInvocationStatus = "Delayed"
	CommandInvocationStatus_Success    CommandInvocationStatus = "Success"
	CommandInvocationStatus_Cancelled  CommandInvocationStatus = "Cancelled"
	CommandInvocationStatus_TimedOut   CommandInvocationStatus = "TimedOut"
	CommandInvocationStatus_Failed     CommandInvocationStatus = "Failed"
	CommandInvocationStatus_Cancelling CommandInvocationStatus = "Cancelling"
)

type CommandPluginStatus string

const (
	CommandPluginStatus_Pending    CommandPluginStatus = "Pending"
	CommandPluginStatus_InProgress CommandPluginStatus = "InProgress"
	CommandPluginStatus_Success    CommandPluginStatus = "Success"
	CommandPluginStatus_TimedOut   CommandPluginStatus = "TimedOut"
	CommandPluginStatus_Cancelled  CommandPluginStatus = "Cancelled"
	CommandPluginStatus_Failed     CommandPluginStatus = "Failed"
)

type CommandStatus string

const (
	CommandStatus_Pending    CommandStatus = "Pending"
	CommandStatus_InProgress CommandStatus = "InProgress"
	CommandStatus_Success    CommandStatus = "Success"
	CommandStatus_Cancelled  CommandStatus = "Cancelled"
	CommandStatus_Failed     CommandStatus = "Failed"
	CommandStatus_TimedOut   CommandStatus = "TimedOut"
	CommandStatus_Cancelling CommandStatus = "Cancelling"
)

type ComplianceQueryOperatorType string

const (
	ComplianceQueryOperatorType_EQUAL        ComplianceQueryOperatorType = "EQUAL"
	ComplianceQueryOperatorType_NOT_EQUAL    ComplianceQueryOperatorType = "NOT_EQUAL"
	ComplianceQueryOperatorType_BEGIN_WITH   ComplianceQueryOperatorType = "BEGIN_WITH"
	ComplianceQueryOperatorType_LESS_THAN    ComplianceQueryOperatorType = "LESS_THAN"
	ComplianceQueryOperatorType_GREATER_THAN ComplianceQueryOperatorType = "GREATER_THAN"
)

type ComplianceSeverity string

const (
	ComplianceSeverity_CRITICAL      ComplianceSeverity = "CRITICAL"
	ComplianceSeverity_HIGH          ComplianceSeverity = "HIGH"
	ComplianceSeverity_MEDIUM        ComplianceSeverity = "MEDIUM"
	ComplianceSeverity_LOW           ComplianceSeverity = "LOW"
	ComplianceSeverity_INFORMATIONAL ComplianceSeverity = "INFORMATIONAL"
	ComplianceSeverity_UNSPECIFIED   ComplianceSeverity = "UNSPECIFIED"
)

type ComplianceStatus string

const (
	ComplianceStatus_COMPLIANT     ComplianceStatus = "COMPLIANT"
	ComplianceStatus_NON_COMPLIANT ComplianceStatus = "NON_COMPLIANT"
)

type ComplianceUploadType string

const (
	ComplianceUploadType_COMPLETE ComplianceUploadType = "COMPLETE"
	ComplianceUploadType_PARTIAL  ComplianceUploadType = "PARTIAL"
)

type ConnectionStatus string

const (
	ConnectionStatus_connected    ConnectionStatus = "connected"
	ConnectionStatus_notconnected ConnectionStatus = "notconnected"
)

type DescribeActivationsFilterKeys string

const (
	DescribeActivationsFilterKeys_ActivationIds       DescribeActivationsFilterKeys = "ActivationIds"
	DescribeActivationsFilterKeys_DefaultInstanceName DescribeActivationsFilterKeys = "DefaultInstanceName"
	DescribeActivationsFilterKeys_IamRole             DescribeActivationsFilterKeys = "IamRole"
)

type DocumentFilterKey string

const (
	DocumentFilterKey_Name          DocumentFilterKey = "Name"
	DocumentFilterKey_Owner         DocumentFilterKey = "Owner"
	DocumentFilterKey_PlatformTypes DocumentFilterKey = "PlatformTypes"
	DocumentFilterKey_DocumentType  DocumentFilterKey = "DocumentType"
)

type DocumentFormat string

const (
	DocumentFormat_YAML DocumentFormat = "YAML"
	DocumentFormat_JSON DocumentFormat = "JSON"
	DocumentFormat_TEXT DocumentFormat = "TEXT"
)

type DocumentHashType string

const (
	DocumentHashType_Sha256 DocumentHashType = "Sha256"
	DocumentHashType_Sha1   DocumentHashType = "Sha1"
)

type DocumentMetadataEnum string

const (
	DocumentMetadataEnum_DocumentReviews DocumentMetadataEnum = "DocumentReviews"
)

type DocumentParameterType string

const (
	DocumentParameterType_String     DocumentParameterType = "String"
	DocumentParameterType_StringList DocumentParameterType = "StringList"
)

type DocumentPermissionType string

const (
	DocumentPermissionType_Share DocumentPermissionType = "Share"
)

type DocumentReviewAction string

const (
	DocumentReviewAction_SendForReview DocumentReviewAction = "SendForReview"
	DocumentReviewAction_UpdateReview  DocumentReviewAction = "UpdateReview"
	DocumentReviewAction_Approve       DocumentReviewAction = "Approve"
	DocumentReviewAction_Reject        DocumentReviewAction = "Reject"
)

type DocumentReviewCommentType string

const (
	DocumentReviewCommentType_Comment DocumentReviewCommentType = "Comment"
)

type DocumentStatus_SDK string

const (
	DocumentStatus_SDK_Creating DocumentStatus_SDK = "Creating"
	DocumentStatus_SDK_Active   DocumentStatus_SDK = "Active"
	DocumentStatus_SDK_Updating DocumentStatus_SDK = "Updating"
	DocumentStatus_SDK_Deleting DocumentStatus_SDK = "Deleting"
	DocumentStatus_SDK_Failed   DocumentStatus_SDK = "Failed"
)

type DocumentType string

const (
	DocumentType_Command                        DocumentType = "Command"
	DocumentType_Policy                         DocumentType = "Policy"
	DocumentType_Automation                     DocumentType = "Automation"
	DocumentType_Session                        DocumentType = "Session"
	DocumentType_Package                        DocumentType = "Package"
	DocumentType_ApplicationConfiguration       DocumentType = "ApplicationConfiguration"
	DocumentType_ApplicationConfigurationSchema DocumentType = "ApplicationConfigurationSchema"
	DocumentType_DeploymentStrategy             DocumentType = "DeploymentStrategy"
	DocumentType_ChangeCalendar                 DocumentType = "ChangeCalendar"
	DocumentType_Automation_ChangeTemplate      DocumentType = "Automation.ChangeTemplate"
	DocumentType_ProblemAnalysis                DocumentType = "ProblemAnalysis"
	DocumentType_ProblemAnalysisTemplate        DocumentType = "ProblemAnalysisTemplate"
	DocumentType_CloudFormation                 DocumentType = "CloudFormation"
	DocumentType_ConformancePackTemplate        DocumentType = "ConformancePackTemplate"
	DocumentType_QuickSetup                     DocumentType = "QuickSetup"
)

type ExecutionMode string

const (
	ExecutionMode_Auto        ExecutionMode = "Auto"
	ExecutionMode_Interactive ExecutionMode = "Interactive"
)

type ExternalAlarmState string

const (
	ExternalAlarmState_UNKNOWN ExternalAlarmState = "UNKNOWN"
	ExternalAlarmState_ALARM   ExternalAlarmState = "ALARM"
)

type Fault string

const (
	Fault_Client  Fault = "Client"
	Fault_Server  Fault = "Server"
	Fault_Unknown Fault = "Unknown"
)

type InstanceInformationFilterKey string

const (
	InstanceInformationFilterKey_InstanceIds       InstanceInformationFilterKey = "InstanceIds"
	InstanceInformationFilterKey_AgentVersion      InstanceInformationFilterKey = "AgentVersion"
	InstanceInformationFilterKey_PingStatus        InstanceInformationFilterKey = "PingStatus"
	InstanceInformationFilterKey_PlatformTypes     InstanceInformationFilterKey = "PlatformTypes"
	InstanceInformationFilterKey_ActivationIds     InstanceInformationFilterKey = "ActivationIds"
	InstanceInformationFilterKey_IamRole           InstanceInformationFilterKey = "IamRole"
	InstanceInformationFilterKey_ResourceType      InstanceInformationFilterKey = "ResourceType"
	InstanceInformationFilterKey_AssociationStatus InstanceInformationFilterKey = "AssociationStatus"
)

type InstancePatchStateOperatorType string

const (
	InstancePatchStateOperatorType_Equal       InstancePatchStateOperatorType = "Equal"
	InstancePatchStateOperatorType_NotEqual    InstancePatchStateOperatorType = "NotEqual"
	InstancePatchStateOperatorType_LessThan    InstancePatchStateOperatorType = "LessThan"
	InstancePatchStateOperatorType_GreaterThan InstancePatchStateOperatorType = "GreaterThan"
)

type InventoryAttributeDataType string

const (
	InventoryAttributeDataType_string InventoryAttributeDataType = "string"
	InventoryAttributeDataType_number InventoryAttributeDataType = "number"
)

type InventoryDeletionStatus string

const (
	InventoryDeletionStatus_InProgress InventoryDeletionStatus = "InProgress"
	InventoryDeletionStatus_Complete   InventoryDeletionStatus = "Complete"
)

type InventoryQueryOperatorType string

const (
	InventoryQueryOperatorType_Equal       InventoryQueryOperatorType = "Equal"
	InventoryQueryOperatorType_NotEqual    InventoryQueryOperatorType = "NotEqual"
	InventoryQueryOperatorType_BeginWith   InventoryQueryOperatorType = "BeginWith"
	InventoryQueryOperatorType_LessThan    InventoryQueryOperatorType = "LessThan"
	InventoryQueryOperatorType_GreaterThan InventoryQueryOperatorType = "GreaterThan"
	InventoryQueryOperatorType_Exists      InventoryQueryOperatorType = "Exists"
)

type InventorySchemaDeleteOption string

const (
	InventorySchemaDeleteOption_DisableSchema InventorySchemaDeleteOption = "DisableSchema"
	InventorySchemaDeleteOption_DeleteSchema  InventorySchemaDeleteOption = "DeleteSchema"
)

type LastResourceDataSyncStatus string

const (
	LastResourceDataSyncStatus_Successful LastResourceDataSyncStatus = "Successful"
	LastResourceDataSyncStatus_Failed     LastResourceDataSyncStatus = "Failed"
	LastResourceDataSyncStatus_InProgress LastResourceDataSyncStatus = "InProgress"
)

type MaintenanceWindowExecutionStatus string

const (
	MaintenanceWindowExecutionStatus_PENDING             MaintenanceWindowExecutionStatus = "PENDING"
	MaintenanceWindowExecutionStatus_IN_PROGRESS         MaintenanceWindowExecutionStatus = "IN_PROGRESS"
	MaintenanceWindowExecutionStatus_SUCCESS             MaintenanceWindowExecutionStatus = "SUCCESS"
	MaintenanceWindowExecutionStatus_FAILED              MaintenanceWindowExecutionStatus = "FAILED"
	MaintenanceWindowExecutionStatus_TIMED_OUT           MaintenanceWindowExecutionStatus = "TIMED_OUT"
	MaintenanceWindowExecutionStatus_CANCELLING          MaintenanceWindowExecutionStatus = "CANCELLING"
	MaintenanceWindowExecutionStatus_CANCELLED           MaintenanceWindowExecutionStatus = "CANCELLED"
	MaintenanceWindowExecutionStatus_SKIPPED_OVERLAPPING MaintenanceWindowExecutionStatus = "SKIPPED_OVERLAPPING"
)

type MaintenanceWindowResourceType string

const (
	MaintenanceWindowResourceType_INSTANCE       MaintenanceWindowResourceType = "INSTANCE"
	MaintenanceWindowResourceType_RESOURCE_GROUP MaintenanceWindowResourceType = "RESOURCE_GROUP"
)

type MaintenanceWindowTaskCutoffBehavior string

const (
	MaintenanceWindowTaskCutoffBehavior_CONTINUE_TASK MaintenanceWindowTaskCutoffBehavior = "CONTINUE_TASK"
	MaintenanceWindowTaskCutoffBehavior_CANCEL_TASK   MaintenanceWindowTaskCutoffBehavior = "CANCEL_TASK"
)

type MaintenanceWindowTaskType string

const (
	MaintenanceWindowTaskType_RUN_COMMAND    MaintenanceWindowTaskType = "RUN_COMMAND"
	MaintenanceWindowTaskType_AUTOMATION     MaintenanceWindowTaskType = "AUTOMATION"
	MaintenanceWindowTaskType_STEP_FUNCTIONS MaintenanceWindowTaskType = "STEP_FUNCTIONS"
	MaintenanceWindowTaskType_LAMBDA         MaintenanceWindowTaskType = "LAMBDA"
)

type NotificationEvent string

const (
	NotificationEvent_All        NotificationEvent = "All"
	NotificationEvent_InProgress NotificationEvent = "InProgress"
	NotificationEvent_Success    NotificationEvent = "Success"
	NotificationEvent_TimedOut   NotificationEvent = "TimedOut"
	NotificationEvent_Cancelled  NotificationEvent = "Cancelled"
	NotificationEvent_Failed     NotificationEvent = "Failed"
)

type NotificationType string

const (
	NotificationType_Command    NotificationType = "Command"
	NotificationType_Invocation NotificationType = "Invocation"
)

type OperatingSystem string

const (
	OperatingSystem_WINDOWS                 OperatingSystem = "WINDOWS"
	OperatingSystem_AMAZON_LINUX            OperatingSystem = "AMAZON_LINUX"
	OperatingSystem_AMAZON_LINUX_2          OperatingSystem = "AMAZON_LINUX_2"
	OperatingSystem_AMAZON_LINUX_2022       OperatingSystem = "AMAZON_LINUX_2022"
	OperatingSystem_UBUNTU                  OperatingSystem = "UBUNTU"
	OperatingSystem_REDHAT_ENTERPRISE_LINUX OperatingSystem = "REDHAT_ENTERPRISE_LINUX"
	OperatingSystem_SUSE                    OperatingSystem = "SUSE"
	OperatingSystem_CENTOS                  OperatingSystem = "CENTOS"
	OperatingSystem_ORACLE_LINUX            OperatingSystem = "ORACLE_LINUX"
	OperatingSystem_DEBIAN                  OperatingSystem = "DEBIAN"
	OperatingSystem_MACOS                   OperatingSystem = "MACOS"
	OperatingSystem_RASPBIAN                OperatingSystem = "RASPBIAN"
	OperatingSystem_ROCKY_LINUX             OperatingSystem = "ROCKY_LINUX"
	OperatingSystem_ALMA_LINUX              OperatingSystem = "ALMA_LINUX"
	OperatingSystem_AMAZON_LINUX_2023       OperatingSystem = "AMAZON_LINUX_2023"
)

type OpsFilterOperatorType string

const (
	OpsFilterOperatorType_Equal       OpsFilterOperatorType = "Equal"
	OpsFilterOperatorType_NotEqual    OpsFilterOperatorType = "NotEqual"
	OpsFilterOperatorType_BeginWith   OpsFilterOperatorType = "BeginWith"
	OpsFilterOperatorType_LessThan    OpsFilterOperatorType = "LessThan"
	OpsFilterOperatorType_GreaterThan OpsFilterOperatorType = "GreaterThan"
	OpsFilterOperatorType_Exists      OpsFilterOperatorType = "Exists"
)

type OpsItemDataType string

const (
	OpsItemDataType_SearchableString OpsItemDataType = "SearchableString"
	OpsItemDataType_String           OpsItemDataType = "String"
)

type OpsItemEventFilterKey string

const (
	OpsItemEventFilterKey_OpsItemId OpsItemEventFilterKey = "OpsItemId"
)

type OpsItemEventFilterOperator string

const (
	OpsItemEventFilterOperator_Equal OpsItemEventFilterOperator = "Equal"
)

type OpsItemFilterKey string

const (
	OpsItemFilterKey_Status                              OpsItemFilterKey = "Status"
	OpsItemFilterKey_CreatedBy                           OpsItemFilterKey = "CreatedBy"
	OpsItemFilterKey_Source                              OpsItemFilterKey = "Source"
	OpsItemFilterKey_Priority                            OpsItemFilterKey = "Priority"
	OpsItemFilterKey_Title                               OpsItemFilterKey = "Title"
	OpsItemFilterKey_OpsItemId                           OpsItemFilterKey = "OpsItemId"
	OpsItemFilterKey_CreatedTime                         OpsItemFilterKey = "CreatedTime"
	OpsItemFilterKey_LastModifiedTime                    OpsItemFilterKey = "LastModifiedTime"
	OpsItemFilterKey_ActualStartTime                     OpsItemFilterKey = "ActualStartTime"
	OpsItemFilterKey_ActualEndTime                       OpsItemFilterKey = "ActualEndTime"
	OpsItemFilterKey_PlannedStartTime                    OpsItemFilterKey = "PlannedStartTime"
	OpsItemFilterKey_PlannedEndTime                      OpsItemFilterKey = "PlannedEndTime"
	OpsItemFilterKey_OperationalData                     OpsItemFilterKey = "OperationalData"
	OpsItemFilterKey_OperationalDataKey                  OpsItemFilterKey = "OperationalDataKey"
	OpsItemFilterKey_OperationalDataValue                OpsItemFilterKey = "OperationalDataValue"
	OpsItemFilterKey_ResourceId                          OpsItemFilterKey = "ResourceId"
	OpsItemFilterKey_AutomationId                        OpsItemFilterKey = "AutomationId"
	OpsItemFilterKey_Category                            OpsItemFilterKey = "Category"
	OpsItemFilterKey_Severity                            OpsItemFilterKey = "Severity"
	OpsItemFilterKey_OpsItemType                         OpsItemFilterKey = "OpsItemType"
	OpsItemFilterKey_ChangeRequestByRequesterArn         OpsItemFilterKey = "ChangeRequestByRequesterArn"
	OpsItemFilterKey_ChangeRequestByRequesterName        OpsItemFilterKey = "ChangeRequestByRequesterName"
	OpsItemFilterKey_ChangeRequestByApproverArn          OpsItemFilterKey = "ChangeRequestByApproverArn"
	OpsItemFilterKey_ChangeRequestByApproverName         OpsItemFilterKey = "ChangeRequestByApproverName"
	OpsItemFilterKey_ChangeRequestByTemplate             OpsItemFilterKey = "ChangeRequestByTemplate"
	OpsItemFilterKey_ChangeRequestByTargetsResourceGroup OpsItemFilterKey = "ChangeRequestByTargetsResourceGroup"
	OpsItemFilterKey_InsightByType                       OpsItemFilterKey = "InsightByType"
	OpsItemFilterKey_AccountId                           OpsItemFilterKey = "AccountId"
)

type OpsItemFilterOperator string

const (
	OpsItemFilterOperator_Equal       OpsItemFilterOperator = "Equal"
	OpsItemFilterOperator_Contains    OpsItemFilterOperator = "Contains"
	OpsItemFilterOperator_GreaterThan OpsItemFilterOperator = "GreaterThan"
	OpsItemFilterOperator_LessThan    OpsItemFilterOperator = "LessThan"
)

type OpsItemRelatedItemsFilterKey string

const (
	OpsItemRelatedItemsFilterKey_ResourceType  OpsItemRelatedItemsFilterKey = "ResourceType"
	OpsItemRelatedItemsFilterKey_AssociationId OpsItemRelatedItemsFilterKey = "AssociationId"
	OpsItemRelatedItemsFilterKey_ResourceUri   OpsItemRelatedItemsFilterKey = "ResourceUri"
)

type OpsItemRelatedItemsFilterOperator string

const (
	OpsItemRelatedItemsFilterOperator_Equal OpsItemRelatedItemsFilterOperator = "Equal"
)

type OpsItemStatus string

const (
	OpsItemStatus_Open                           OpsItemStatus = "Open"
	OpsItemStatus_InProgress                     OpsItemStatus = "InProgress"
	OpsItemStatus_Resolved                       OpsItemStatus = "Resolved"
	OpsItemStatus_Pending                        OpsItemStatus = "Pending"
	OpsItemStatus_TimedOut                       OpsItemStatus = "TimedOut"
	OpsItemStatus_Cancelling                     OpsItemStatus = "Cancelling"
	OpsItemStatus_Cancelled                      OpsItemStatus = "Cancelled"
	OpsItemStatus_Failed                         OpsItemStatus = "Failed"
	OpsItemStatus_CompletedWithSuccess           OpsItemStatus = "CompletedWithSuccess"
	OpsItemStatus_CompletedWithFailure           OpsItemStatus = "CompletedWithFailure"
	OpsItemStatus_Scheduled                      OpsItemStatus = "Scheduled"
	OpsItemStatus_RunbookInProgress              OpsItemStatus = "RunbookInProgress"
	OpsItemStatus_PendingChangeCalendarOverride  OpsItemStatus = "PendingChangeCalendarOverride"
	OpsItemStatus_ChangeCalendarOverrideApproved OpsItemStatus = "ChangeCalendarOverrideApproved"
	OpsItemStatus_ChangeCalendarOverrideRejected OpsItemStatus = "ChangeCalendarOverrideRejected"
	OpsItemStatus_PendingApproval                OpsItemStatus = "PendingApproval"
	OpsItemStatus_Approved                       OpsItemStatus = "Approved"
	OpsItemStatus_Rejected                       OpsItemStatus = "Rejected"
	OpsItemStatus_Closed                         OpsItemStatus = "Closed"
)

type ParameterTier string

const (
	ParameterTier_Standard            ParameterTier = "Standard"
	ParameterTier_Advanced            ParameterTier = "Advanced"
	ParameterTier_Intelligent_Tiering ParameterTier = "Intelligent-Tiering"
)

type ParameterType string

const (
	ParameterType_String       ParameterType = "String"
	ParameterType_StringList   ParameterType = "StringList"
	ParameterType_SecureString ParameterType = "SecureString"
)

type ParametersFilterKey string

const (
	ParametersFilterKey_Name  ParametersFilterKey = "Name"
	ParametersFilterKey_Type  ParametersFilterKey = "Type"
	ParametersFilterKey_KeyId ParametersFilterKey = "KeyId"
)

type PatchAction string

const (
	PatchAction_ALLOW_AS_DEPENDENCY PatchAction = "ALLOW_AS_DEPENDENCY"
	PatchAction_BLOCK               PatchAction = "BLOCK"
)

type PatchComplianceDataState string

const (
	PatchComplianceDataState_INSTALLED                PatchComplianceDataState = "INSTALLED"
	PatchComplianceDataState_INSTALLED_OTHER          PatchComplianceDataState = "INSTALLED_OTHER"
	PatchComplianceDataState_INSTALLED_PENDING_REBOOT PatchComplianceDataState = "INSTALLED_PENDING_REBOOT"
	PatchComplianceDataState_INSTALLED_REJECTED       PatchComplianceDataState = "INSTALLED_REJECTED"
	PatchComplianceDataState_MISSING                  PatchComplianceDataState = "MISSING"
	PatchComplianceDataState_NOT_APPLICABLE           PatchComplianceDataState = "NOT_APPLICABLE"
	PatchComplianceDataState_FAILED                   PatchComplianceDataState = "FAILED"
)

type PatchComplianceLevel string

const (
	PatchComplianceLevel_CRITICAL      PatchComplianceLevel = "CRITICAL"
	PatchComplianceLevel_HIGH          PatchComplianceLevel = "HIGH"
	PatchComplianceLevel_MEDIUM        PatchComplianceLevel = "MEDIUM"
	PatchComplianceLevel_LOW           PatchComplianceLevel = "LOW"
	PatchComplianceLevel_INFORMATIONAL PatchComplianceLevel = "INFORMATIONAL"
	PatchComplianceLevel_UNSPECIFIED   PatchComplianceLevel = "UNSPECIFIED"
)

type PatchDeploymentStatus string

const (
	PatchDeploymentStatus_APPROVED          PatchDeploymentStatus = "APPROVED"
	PatchDeploymentStatus_PENDING_APPROVAL  PatchDeploymentStatus = "PENDING_APPROVAL"
	PatchDeploymentStatus_EXPLICIT_APPROVED PatchDeploymentStatus = "EXPLICIT_APPROVED"
	PatchDeploymentStatus_EXPLICIT_REJECTED PatchDeploymentStatus = "EXPLICIT_REJECTED"
)

type PatchFilterKey string

const (
	PatchFilterKey_ARCH           PatchFilterKey = "ARCH"
	PatchFilterKey_ADVISORY_ID    PatchFilterKey = "ADVISORY_ID"
	PatchFilterKey_BUGZILLA_ID    PatchFilterKey = "BUGZILLA_ID"
	PatchFilterKey_PATCH_SET      PatchFilterKey = "PATCH_SET"
	PatchFilterKey_PRODUCT        PatchFilterKey = "PRODUCT"
	PatchFilterKey_PRODUCT_FAMILY PatchFilterKey = "PRODUCT_FAMILY"
	PatchFilterKey_CLASSIFICATION PatchFilterKey = "CLASSIFICATION"
	PatchFilterKey_CVE_ID         PatchFilterKey = "CVE_ID"
	PatchFilterKey_EPOCH          PatchFilterKey = "EPOCH"
	PatchFilterKey_MSRC_SEVERITY  PatchFilterKey = "MSRC_SEVERITY"
	PatchFilterKey_NAME           PatchFilterKey = "NAME"
	PatchFilterKey_PATCH_ID       PatchFilterKey = "PATCH_ID"
	PatchFilterKey_SECTION        PatchFilterKey = "SECTION"
	PatchFilterKey_PRIORITY       PatchFilterKey = "PRIORITY"
	PatchFilterKey_REPOSITORY     PatchFilterKey = "REPOSITORY"
	PatchFilterKey_RELEASE        PatchFilterKey = "RELEASE"
	PatchFilterKey_SEVERITY       PatchFilterKey = "SEVERITY"
	PatchFilterKey_SECURITY       PatchFilterKey = "SECURITY"
	PatchFilterKey_VERSION        PatchFilterKey = "VERSION"
)

type PatchOperationType string

const (
	PatchOperationType_Scan    PatchOperationType = "Scan"
	PatchOperationType_Install PatchOperationType = "Install"
)

type PatchProperty string

const (
	PatchProperty_PRODUCT        PatchProperty = "PRODUCT"
	PatchProperty_PRODUCT_FAMILY PatchProperty = "PRODUCT_FAMILY"
	PatchProperty_CLASSIFICATION PatchProperty = "CLASSIFICATION"
	PatchProperty_MSRC_SEVERITY  PatchProperty = "MSRC_SEVERITY"
	PatchProperty_PRIORITY       PatchProperty = "PRIORITY"
	PatchProperty_SEVERITY       PatchProperty = "SEVERITY"
)

type PatchSet string

const (
	PatchSet_OS          PatchSet = "OS"
	PatchSet_APPLICATION PatchSet = "APPLICATION"
)

type PingStatus string

const (
	PingStatus_Online         PingStatus = "Online"
	PingStatus_ConnectionLost PingStatus = "ConnectionLost"
	PingStatus_Inactive       PingStatus = "Inactive"
)

type PlatformType string

const (
	PlatformType_Windows PlatformType = "Windows"
	PlatformType_Linux   PlatformType = "Linux"
	PlatformType_MacOS   PlatformType = "MacOS"
)

type RebootOption string

const (
	RebootOption_RebootIfNeeded RebootOption = "RebootIfNeeded"
	RebootOption_NoReboot       RebootOption = "NoReboot"
)

type ResourceDataSyncS3Format string

const (
	ResourceDataSyncS3Format_JsonSerDe ResourceDataSyncS3Format = "JsonSerDe"
)

type ResourceType string

const (
	ResourceType_ManagedInstance ResourceType = "ManagedInstance"
	ResourceType_EC2Instance     ResourceType = "EC2Instance"
)

type ResourceTypeForTagging string

const (
	ResourceTypeForTagging_Document          ResourceTypeForTagging = "Document"
	ResourceTypeForTagging_ManagedInstance   ResourceTypeForTagging = "ManagedInstance"
	ResourceTypeForTagging_MaintenanceWindow ResourceTypeForTagging = "MaintenanceWindow"
	ResourceTypeForTagging_Parameter         ResourceTypeForTagging = "Parameter"
	ResourceTypeForTagging_PatchBaseline     ResourceTypeForTagging = "PatchBaseline"
	ResourceTypeForTagging_OpsItem           ResourceTypeForTagging = "OpsItem"
	ResourceTypeForTagging_OpsMetadata       ResourceTypeForTagging = "OpsMetadata"
	ResourceTypeForTagging_Automation        ResourceTypeForTagging = "Automation"
	ResourceTypeForTagging_Association       ResourceTypeForTagging = "Association"
)

type ReviewStatus string

const (
	ReviewStatus_APPROVED     ReviewStatus = "APPROVED"
	ReviewStatus_NOT_REVIEWED ReviewStatus = "NOT_REVIEWED"
	ReviewStatus_PENDING      ReviewStatus = "PENDING"
	ReviewStatus_REJECTED     ReviewStatus = "REJECTED"
)

type SessionFilterKey string

const (
	SessionFilterKey_InvokedAfter  SessionFilterKey = "InvokedAfter"
	SessionFilterKey_InvokedBefore SessionFilterKey = "InvokedBefore"
	SessionFilterKey_Target        SessionFilterKey = "Target"
	SessionFilterKey_Owner         SessionFilterKey = "Owner"
	SessionFilterKey_Status        SessionFilterKey = "Status"
	SessionFilterKey_SessionId     SessionFilterKey = "SessionId"
)

type SessionState string

const (
	SessionState_Active  SessionState = "Active"
	SessionState_History SessionState = "History"
)

type SessionStatus string

const (
	SessionStatus_Connected    SessionStatus = "Connected"
	SessionStatus_Connecting   SessionStatus = "Connecting"
	SessionStatus_Disconnected SessionStatus = "Disconnected"
	SessionStatus_Terminated   SessionStatus = "Terminated"
	SessionStatus_Terminating  SessionStatus = "Terminating"
	SessionStatus_Failed       SessionStatus = "Failed"
)

type SignalType string

const (
	SignalType_Approve   SignalType = "Approve"
	SignalType_Reject    SignalType = "Reject"
	SignalType_StartStep SignalType = "StartStep"
	SignalType_StopStep  SignalType = "StopStep"
	SignalType_Resume    SignalType = "Resume"
)

type SourceType string

const (
	SourceType_AWS__EC2__Instance        SourceType = "AWS::EC2::Instance"
	SourceType_AWS__IoT__Thing           SourceType = "AWS::IoT::Thing"
	SourceType_AWS__SSM__ManagedInstance SourceType = "AWS::SSM::ManagedInstance"
)

type StepExecutionFilterKey string

const (
	StepExecutionFilterKey_StartTimeBefore         StepExecutionFilterKey = "StartTimeBefore"
	StepExecutionFilterKey_StartTimeAfter          StepExecutionFilterKey = "StartTimeAfter"
	StepExecutionFilterKey_StepExecutionStatus     StepExecutionFilterKey = "StepExecutionStatus"
	StepExecutionFilterKey_StepExecutionId         StepExecutionFilterKey = "StepExecutionId"
	StepExecutionFilterKey_StepName                StepExecutionFilterKey = "StepName"
	StepExecutionFilterKey_Action                  StepExecutionFilterKey = "Action"
	StepExecutionFilterKey_ParentStepExecutionId   StepExecutionFilterKey = "ParentStepExecutionId"
	StepExecutionFilterKey_ParentStepIteration     StepExecutionFilterKey = "ParentStepIteration"
	StepExecutionFilterKey_ParentStepIteratorValue StepExecutionFilterKey = "ParentStepIteratorValue"
)

type StopType string

const (
	StopType_Complete StopType = "Complete"
	StopType_Cancel   StopType = "Cancel"
)
