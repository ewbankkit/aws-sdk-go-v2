// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

type DiscovererSummary struct {

	// The Status if the discoverer will discover schemas from events sent from
	// another account.
	CrossAccount *bool

	// The ARN of the discoverer.
	DiscovererArn *string

	// The ID of the discoverer.
	DiscovererId *string

	// The ARN of the event bus.
	SourceArn *string

	// The state of the discoverer.
	State DiscovererState

	// Tags associated with the resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type RegistrySummary struct {

	// The ARN of the registry.
	RegistryArn *string

	// The name of the registry.
	RegistryName *string

	// Tags associated with the registry.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A summary of schema details.
type SchemaSummary struct {

	// The date and time that schema was modified.
	LastModified *time.Time

	// The ARN of the schema.
	SchemaArn *string

	// The name of the schema.
	SchemaName *string

	// Tags associated with the schema.
	Tags map[string]string

	// The number of versions available for the schema.
	VersionCount *int64

	noSmithyDocumentSerde
}

type SchemaVersionSummary struct {

	// The ARN of the schema version.
	SchemaArn *string

	// The name of the schema.
	SchemaName *string

	// The version number of the schema.
	SchemaVersion *string

	// The type of schema.
	Type Type

	noSmithyDocumentSerde
}

type SearchSchemaSummary struct {

	// The name of the registry.
	RegistryName *string

	// The ARN of the schema.
	SchemaArn *string

	// The name of the schema.
	SchemaName *string

	// An array of schema version summaries.
	SchemaVersions []SearchSchemaVersionSummary

	noSmithyDocumentSerde
}

type SearchSchemaVersionSummary struct {

	// The date the schema version was created.
	CreatedDate *time.Time

	// The version number of the schema
	SchemaVersion *string

	// The type of schema.
	Type Type

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
