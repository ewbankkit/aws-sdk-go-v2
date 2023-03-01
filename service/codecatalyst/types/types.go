// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about a specified personal access token (PAT).
type AccessTokenSummary struct {

	// The system-generated ID of the personal access token.
	//
	// This member is required.
	Id *string

	// The friendly name of the personal access token.
	//
	// This member is required.
	Name *string

	// The date and time when the personal access token will expire, in coordinated
	// universal time (UTC) timestamp format as specified in RFC 3339
	// (https://www.rfc-editor.org/rfc/rfc3339#section-5.6).
	ExpiresTime *time.Time

	noSmithyDocumentSerde
}

// Information about connection details for a Dev Environment.
type DevEnvironmentAccessDetails struct {

	// The URL used to send commands to and from the Dev Environment.
	//
	// This member is required.
	StreamUrl *string

	// An encrypted token value that contains session and caller information used to
	// authenticate the connection.
	//
	// This member is required.
	TokenValue *string

	noSmithyDocumentSerde
}

// Information about the source repsitory for a Dev Environment.
type DevEnvironmentRepositorySummary struct {

	// The name of the source repository.
	//
	// This member is required.
	RepositoryName *string

	// The name of the branch in a source repository cloned into the Dev Environment.
	BranchName *string

	noSmithyDocumentSerde
}

// Information about the configuration of a Dev Environment session.
type DevEnvironmentSessionConfiguration struct {

	// The type of the session.
	//
	// This member is required.
	SessionType DevEnvironmentSessionType

	// Information about optional commands that will be run on the Dev Environment when
	// the SSH session begins.
	ExecuteCommandSessionConfiguration *ExecuteCommandSessionConfiguration

	noSmithyDocumentSerde
}

// Information about a Dev Environment.
type DevEnvironmentSummary struct {

	// The system-generated unique ID of the user who created the Dev Environment.
	//
	// This member is required.
	CreatorId *string

	// The system-generated unique ID for the Dev Environment.
	//
	// This member is required.
	Id *string

	// The amount of time the Dev Environment will run without any activity detected
	// before stopping, in minutes. Dev Environments consume compute minutes when
	// running.
	//
	// This member is required.
	InactivityTimeoutMinutes int32

	// The Amazon EC2 instace type used for the Dev Environment.
	//
	// This member is required.
	InstanceType InstanceType

	// The time when the Dev Environment was last updated, in coordinated universal
	// time (UTC) timestamp format as specified in RFC 3339
	// (https://www.rfc-editor.org/rfc/rfc3339#section-5.6).
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// Information about the configuration of persistent storage for the Dev
	// Environment.
	//
	// This member is required.
	PersistentStorage *PersistentStorage

	// Information about the repositories that will be cloned into the Dev Environment.
	// If no rvalue is specified, no repository is cloned.
	//
	// This member is required.
	Repositories []DevEnvironmentRepositorySummary

	// The status of the Dev Environment.
	//
	// This member is required.
	Status DevEnvironmentStatus

	// The user-specified alias for the Dev Environment.
	Alias *string

	// Information about the integrated development environment (IDE) configured for a
	// Dev Environment.
	Ides []Ide

	// The name of the project in the space.
	ProjectName *string

	// The name of the space.
	SpaceName *string

	// The reason for the status.
	StatusReason *string

	noSmithyDocumentSerde
}

// Information about an email address.
type EmailAddress struct {

	// The email address.
	Email *string

	// Whether the email address has been verified.
	Verified *bool

	noSmithyDocumentSerde
}

// Information about an entry in an event log of Amazon CodeCatalyst activity.
type EventLogEntry struct {

	// The category for the event.
	//
	// This member is required.
	EventCategory *string

	// The name of the event.
	//
	// This member is required.
	EventName *string

	// The source of the event.
	//
	// This member is required.
	EventSource *string

	// The time the event took place, in coordinated universal time (UTC) timestamp
	// format as specified in RFC 3339
	// (https://www.rfc-editor.org/rfc/rfc3339#section-5.6).
	//
	// This member is required.
	EventTime *time.Time

	// The type of the event.
	//
	// This member is required.
	EventType *string

	// The system-generated unique ID of the event.
	//
	// This member is required.
	Id *string

	// The type of the event.
	//
	// This member is required.
	OperationType OperationType

	// The system-generated unique ID of the user whose actions are recorded in the
	// event.
	//
	// This member is required.
	UserIdentity *UserIdentity

	// The code of the error, if any.
	ErrorCode *string

	// Information about the project where the event occurred.
	ProjectInformation *ProjectInformation

	// The system-generated unique ID of the request.
	RequestId *string

	// Information about the payload of the request.
	RequestPayload *EventPayload

	// Information about the payload of the response, if any.
	ResponsePayload *EventPayload

	// The IP address of the user whose actions are recorded in the event.
	SourceIpAddress *string

	//
	UserAgent *string

	noSmithyDocumentSerde
}

// Information about the payload of an event recording Amazon CodeCatalyst
// activity.
type EventPayload struct {

	// The type of content in the event payload.
	ContentType *string

	// The data included in the event payload.
	Data *string

	noSmithyDocumentSerde
}

// Information about the commands that will be run on a Dev Environment when an SSH
// session begins.
type ExecuteCommandSessionConfiguration struct {

	// The command used at the beginning of the SSH session to a Dev Environment.
	//
	// This member is required.
	Command *string

	// An array of arguments containing arguments and members.
	Arguments []string

	noSmithyDocumentSerde
}

type Filter struct {

	//
	//
	// This member is required.
	Key *string

	//
	//
	// This member is required.
	Values []string

	//
	ComparisonOperator *string

	noSmithyDocumentSerde
}

// Information about an integrated development environment (IDE) used in a Dev
// Environment.
type Ide struct {

	// The name of the IDE.
	Name *string

	// A link to the IDE runtime image.
	Runtime *string

	noSmithyDocumentSerde
}

// Information about the configuration of an integrated development environment
// (IDE) for a Dev Environment.
type IdeConfiguration struct {

	// The name of the IDE. Valid values include Cloud9, IntelliJ, PyCharm, GoLand, and
	// VSCode.
	Name *string

	// A link to the IDE runtime image. This parameter is not required for VSCode.
	Runtime *string

	noSmithyDocumentSerde
}

// Information about a source repository returned in a list of source repositories.
type ListSourceRepositoriesItem struct {

	// The time the source repository was created, in coordinated universal time (UTC)
	// timestamp format as specified in RFC 3339
	// (https://www.rfc-editor.org/rfc/rfc3339#section-5.6).
	//
	// This member is required.
	CreatedTime *time.Time

	// The system-generated unique ID of the source repository.
	//
	// This member is required.
	Id *string

	// The time the source repository was last updated, in coordinated universal time
	// (UTC) timestamp format as specified in RFC 3339
	// (https://www.rfc-editor.org/rfc/rfc3339#section-5.6).
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The name of the source repository.
	//
	// This member is required.
	Name *string

	// The description of the repository, if any.
	Description *string

	noSmithyDocumentSerde
}

// Information about a branch of a source repository returned in a list of
// branches.
type ListSourceRepositoryBranchesItem struct {

	// The commit ID of the tip of the branch at the time of the request, also known as
	// the head commit.
	HeadCommitId *string

	// The time the branch was last updated, in coordinated universal time (UTC)
	// timestamp format as specified in RFC 3339
	// (https://www.rfc-editor.org/rfc/rfc3339#section-5.6).
	LastUpdatedTime *time.Time

	// The name of the branch.
	Name *string

	// The Git reference name of the branch.
	Ref *string

	noSmithyDocumentSerde
}

// Information about the persistent storage for a Dev Environment.
type PersistentStorage struct {

	// The size of the persistent storage in gigabytes (specifically GiB). Valid values
	// for storage are based on memory sizes in 16GB increments. Valid values are 16,
	// 32, and 64.
	//
	// This member is required.
	SizeInGiB *int32

	noSmithyDocumentSerde
}

// Information about the configuration of persistent storage for a Dev Environment.
type PersistentStorageConfiguration struct {

	// The size of the persistent storage in gigabytes (specifically GiB). Valid values
	// for storage are based on memory sizes in 16GB increments. Valid values are 16,
	// 32, and 64.
	//
	// This member is required.
	SizeInGiB *int32

	noSmithyDocumentSerde
}

// Information about a project in a space.
type ProjectInformation struct {

	// The name of the project in the space.
	Name *string

	// The system-generated unique ID of the project.
	ProjectId *string

	noSmithyDocumentSerde
}

// nformation about the filter used to narrow the results returned in a list of
// projects.
type ProjectListFilter struct {

	// A key that can be used to sort results.
	//
	// This member is required.
	Key FilterKey

	// The value of the key.
	//
	// This member is required.
	Values []string

	// The operator used to compare the fields.
	ComparisonOperator ComparisonOperator

	noSmithyDocumentSerde
}

// Information about a project.
type ProjectSummary struct {

	// The name of the project in the space.
	//
	// This member is required.
	Name *string

	// The description of the project.
	Description *string

	// The friendly name displayed to users of the project in Amazon CodeCatalyst.
	DisplayName *string

	noSmithyDocumentSerde
}

// Information about a repository that will be cloned to a Dev Environment.
type RepositoryInput struct {

	// The name of the source repository.
	//
	// This member is required.
	RepositoryName *string

	// The name of the branch in a source repository.
	BranchName *string

	noSmithyDocumentSerde
}

// Information about an space.
type SpaceSummary struct {

	// The name of the space.
	//
	// This member is required.
	Name *string

	// The Amazon Web Services Region where the space exists.
	//
	// This member is required.
	RegionName *string

	// The description of the space.
	Description *string

	// The friendly name of the space displayed to users.
	DisplayName *string

	noSmithyDocumentSerde
}

// Information about a user whose activity is recorded in an event for a space.
type UserIdentity struct {

	//
	//
	// This member is required.
	PrincipalId *string

	// The role assigned to the user in a Amazon CodeCatalyst space or project when the
	// event occurred.
	//
	// This member is required.
	UserType UserType

	// The Amazon Web Services account number of the user in Amazon Web Services, if
	// any.
	AwsAccountId *string

	// The display name of the user in Amazon CodeCatalyst.
	UserName *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
