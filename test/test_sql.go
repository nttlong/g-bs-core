// test connect to mssql server
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	db, err := sql.Open("mssql", "sqlserver://sa:nttl0ng2024@127.0.0.1:1433")
	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}
	//try connect
	e := db.Ping()
	if e != nil {
		fmt.Println(e)
		return
	}
	// d create test table'
	_, err = db.Exec("CREATE TABLE test (id INT, name VARCHAR(50))")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	fmt.Println("Connected to database")
}
