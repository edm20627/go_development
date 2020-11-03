package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io"

	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-gorp/gorp"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"gopkg.in/go-playground/validator.v9"
)

type Comment struct {
	Id      int64     `json:"id" db:"id,primarykey,autoincrement"`
	Name    string    `json:"name" db:"name,notnull,size:200" validate:"required,max=200"`
	Text    string    `json:"text" db:"text,notnull,size:399" validate:"required,max=399"`
	Created time.Time `json:"created" db:"created,notnull"`
	Updated time.Time `json:"updated" db:"updated,notnull"`
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

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

const (
	ExitCodeOK int = iota
	ExitCodeError
	ExitCodeFileError
)

var Version = "0.0.0"

type CLI struct {
	outStream, errStream io.Writer
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

	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.Static("/", "static/")

	e.GET("/api/comments", func(c echo.Context) error {
		var comments []Comment
		_, err := dbmap.Select(&comments,
			"select * from comments order by created desc limit 10")
		if err != nil {
			c.Logger().Error("Select: ", err)
			return c.String(http.StatusBadRequest, "Select: "+err.Error())
		}
		return c.JSON(http.StatusOK, comments)
	})

	e.POST("/api/comments", func(c echo.Context) error {
		var comment Comment
		if err := c.Bind(&comment); err != nil {
			c.Logger().Error("Bind: ", err)
			return c.String(http.StatusBadRequest, "Bind: "+err.Error())
		}
		if err = c.Validate(&comment); err != nil {
			c.Logger().Error("Validate: ", err)
			return c.String(http.StatusBadRequest, "Validate: "+err.Error())
		}
		if err = dbmap.Insert(&comment); err != nil {
			c.Logger().Error("Insert: ", err)
			return c.String(http.StatusBadRequest, "Insert: "+err.Error())
		}
		c.Logger().Info("ADDED: %v", comment.Id)
		return c.JSON(http.StatusCreated, "")
	})

	e.Logger.Fatal(e.Start(":8080"))

	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

func (c *CLI) Run(args []string) int {
	os.Args = args // 簡易テスト用
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	if showVersion {
		fmt.Fprintf(c.outStream, "version: %s \n", Version)
		return 0
	} else {
		fmt.Fprintln(c.errStream, "バージョンオプションがありません")
		return 1
	}
}
