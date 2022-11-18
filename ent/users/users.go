// Code generated by ent, DO NOT EDIT.

package users

const (
	// Label holds the string label denoting the users type in the database.
	Label = "users"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldIntroduction holds the string denoting the introduction field in the database.
	FieldIntroduction = "introduction"
	// Table holds the table name of the users in the database.
	Table = "users"
)

// Columns holds all SQL columns for users fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldIntroduction,
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