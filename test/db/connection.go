package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

func NewConnection() *sql.DB {
	buf, err := ioutil.ReadFile("./test/db/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	config := make(map[string]map[string]string, 20)
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		log.Fatal(err)
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config["db"]["user"], config["db"]["password"], config["db"]["host"], config["db"]["port"], config["db"]["database"])
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
