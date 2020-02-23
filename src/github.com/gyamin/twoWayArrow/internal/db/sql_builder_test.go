package db

import (
	"testing"
)

type User struct {
	name string
	age  int
	tel  string
}

type Users []User

func TestBuildBulkInsertSql(t *testing.T) {

	var users []map[string]interface{}

	users = append(users)

	users = append(users, User{"太郎", 31, "012340000001"})
	users = append(users, User{"次郎", 21, "012340000002"})

	sql := BuildBulkInsertSql("users", users)

}
