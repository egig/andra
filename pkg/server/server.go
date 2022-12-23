package server

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

var r *gin.Engine
var db *sql.DB


type Context struct {
	gin.Context
}

func init() {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	r = gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


	r.GET("/_types", func(c *gin.Context) {
		// type storage -> type driver -> list
		//c.Set()
		rows, err := db.Query("select table_name from information_schema.tables where table_schema='testapp0'")
		if err != nil {
			panic(err)
		}

		names := []string{}
		for rows.Next() {
			var tName string
			rows.Scan(&tName)
			names = append(names, tName)

		}

		fmt.Println(names)
	})
}

func Run(context context.Context) {
	port := context.Value("port").(string)

	var projDir = context.Value("dir").(string)
	//r.Use(static.Serve("/assets", static.LocalFile(path.Join(projDir, "assets"), false)))

	r.NoRoute(func(c *gin.Context) {

		// Create route resolver here
		md := goldmark.New(
			goldmark.WithExtensions(
				meta.Meta,
			),
		)

		var buf bytes.Buffer
		contentDir := path.Join(projDir, "content")
		var file = fmt.Sprintf("%s%s.md", contentDir, c.Request.RequestURI)
		src, err := ioutil.ReadFile(file)
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				c.Writer.WriteHeader(http.StatusNotFound)
				return
			}
			panic(err)
		}

		ctx := parser.NewContext()
		err = md.Convert(src, &buf, parser.WithContext(ctx))
		if err != nil {
			panic(err)
		}

		metaData := meta.Get(ctx)
		c.JSON(http.StatusOK, map[string]interface{}{
			"meta": metaData,
			"body": buf.String(),
		})
	})

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
