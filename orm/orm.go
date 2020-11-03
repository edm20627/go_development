package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-gorp/gorp"

	_ "github.com/lib/pq"
)

type Comment struct {
	Id      int64     `db:"id,primarykey,autoincrement"`
	Name    string    `db:"name,default:'名前なし',size:200"`
	Text    string    `db:"text,size:400"`
	Created time.Time `db:"created,notnull"`
	Updated time.Time `db:"updated,notnull"`
}

func (c *Comment) PreInsert(s gorp.SqlExecutor) error {
	c.Created = time.Now()
	c.Updated = c.Created
	return nil
}

func (c *Comment) PreUpdate(s gorp.SqlExecutor) error {
	c.Updated = time.Now()
	return nil
}

func main() {
	dsn := "host=postgres user=user dbname=app_db password=password sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Comment{}, "comments")
	err = dbmap.CreateTablesIfNotExists()

	err = dbmap.Insert(&Comment{Text: "こんにちわ"})
	if err != nil {
		log.Fatal(err)
	}

	var comment Comment
	dbmap.SelectOne(&comment, "select * from comments where id = 1")
	fmt.Println("SelectOne", comment)

	comment.Text = "こんばんは"
	dbmap.Update(&comment)
	fmt.Println("update", comment)

	var comments []Comment
	dbmap.Select(&comments, "select * from comments where name = $1", "名前なし")
	fmt.Println("Select", comments)
}
