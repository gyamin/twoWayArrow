package db

import (
	"testing"
)

func TestBuildInsertSql(t *testing.T) {

	var users []map[string]interface{}

	user := make(map[string]interface{})
	user["name"] = "太郎"
	user["age"] = 20
	user["tel"] = "09010002000"
	users = append(users, user)

	user = make(map[string]interface{})
	user["name"] = "次郎"
	user["age"] = 21
	user["tel"] = "09010002001"
	users = append(users, user)

	sql := BuildInsertSql("users", users)

	println(sql)

}
