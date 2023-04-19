// Code generated by ent, DO NOT EDIT.

package chatconfig

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the chatconfig type in the database.
	Label = "chat_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChatID holds the string denoting the chat_id field in the database.
	FieldChatID = "chat_id"
	// FieldJSON holds the string denoting the json field in the database.
	FieldJSON = "json"
	// Table holds the table name of the chatconfig in the database.
	Table = "chat_configs"
)

// Columns holds all SQL columns for chatconfig fields.
var Columns = []string{
	FieldID,
	FieldChatID,
	FieldJSON,
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

// OrderOption defines the ordering options for the ChatConfig queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChatID orders the results by the chat_id field.
func ByChatID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChatID, opts...).ToFunc()
}

// ByJSON orders the results by the json field.
func ByJSON(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJSON, opts...).ToFunc()
}