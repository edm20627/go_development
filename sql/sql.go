package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// dsn := os.Open("DNS")
	dsn := "host=postgres user=user dbname=app_db password=password sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// insert
	result, err := db.Exec(`INSERT INTO users(name, age) VALUES($1, $2)`, "Bob", 18)
	if err != nil {
		log.Fatal(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("aaa")
		log.Fatal(err)
	}

	// rows, err := db.Query(`SELECT id, name, age FROM users ORDER BY name`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for rows.Next() {
	// 	var id int64
	// 	var name string
	// 	var age int64
	// 	err = rows.Scan(&id, &name, &age)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(id, name, age)
	// }

	// row := db.QueryRow(`SELECT id, name, age FROM users WHERE id=$1`, 1)
	// var id interface{}
	// var name interface{}
	// var age interface{}
	// err = row.Scan(&id, &name, &age)
	// fmt.Println(id, name, age)

	rows, err := db.Query(`SELECT id, name, age FROM users ORDER BY name`)
	if err != nil {
		log.Fatal(err)
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	refs := make([]interface{}, len(columns))
	for i := 0; i < len(columns); i++ {
		refs[i] = &values[1]
	}

	for rows.Next() {

		err = rows.Scan(refs...)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(values...)
	}

}
