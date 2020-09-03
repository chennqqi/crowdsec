// Code generated by entc, DO NOT EDIT.

package meta

import (
	"time"
)

const (
	// Label holds the string label denoting the meta type in the database.
	Label = "meta"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the meta in the database.
	Table = "meta"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "meta"
	// OwnerInverseTable is the table name for the Signal entity.
	// It exists in this package in order to avoid circular dependency with the "signal" package.
	OwnerInverseTable = "signals"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "signal_metas"
)

// Columns holds all SQL columns for meta fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldKey,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Meta type.
var ForeignKeys = []string{
	"signal_metas",
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
)