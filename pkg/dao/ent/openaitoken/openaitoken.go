// Code generated by ent, DO NOT EDIT.

package openaitoken

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the openaitoken type in the database.
	Label = "open_ai_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// Table holds the table name of the openaitoken in the database.
	Table = "open_ai_tokens"
)

// Columns holds all SQL columns for openaitoken fields.
var Columns = []string{
	FieldID,
	FieldToken,
	FieldIsActive,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TokenValidator is a validator for the "token" field. It is called by the builders before save.
	TokenValidator func(string) error
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
)

// OrderOption defines the ordering options for the OpenAIToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}
