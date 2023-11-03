package dbutils

import "auth-api/globals"

func RowExists[T any](cond string, check_func func(*T) bool, args ...any) (*T, bool) {
	possible_match := new(T)
	globals.Db.Model(new(T)).Where(cond, args...).First(possible_match)
	return possible_match, check_func(possible_match)
}
