package util

import (
	"unicode"
)

func ConvertToMySqlNamePattern(tableName string) string {
	var newTableName string

	for i, char := range tableName {
		if unicode.IsUpper(char) && i != 0 {
			newTableName += "_"
		}
		newTableName += string(unicode.ToLower(char))
	}

	return newTableName
}
