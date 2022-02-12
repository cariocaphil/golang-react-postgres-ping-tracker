package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func registerPing(db *sql.DB) {

	_, err := db.Exec("INSERT INTO ping_timestamp (occurred) VALUES ($1)", time.Now())
	if err != nil {
		log.Println("Couldn't insert the ping")
		log.Println(err)
	}
}

func pingFunc(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		defer registerPing(db)
		r := db.QueryRow("SELECT occurred FROM ping_timestamp ORDER BY id DESC LIMIT 1")
		var lastDate pq.NullTime
		r.Scan(&lastDate)

		timeMessage := "first time!"
		if lastDate.Valid {
			timeMessage = fmt.Sprintf("%v ago", time.Now().Sub(lastDate.Time).String())
		}

		var occurred pq.NullTime

		// Query the DB
		rows, err := db.Query(`SELECT occurred FROM ping_timestamp ORDER BY id DESC LIMIT 3`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var occurences []string
		for rows.Next() {
			err := rows.Scan(&occurred)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(occurred)
			occurences = append(occurences, occurred.Time.String())
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		c.JSON(200, gin.H{
			"message":     "successful!" + " " + timeMessage,
			"occurrences": occurences,
		})
	}
}

func main() {

	r := gin.Default()
	// Serving static content from web - we will populate this from within the docker container
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	dbUrl := os.Getenv("DATABASE_URL")
	log.Printf("DB [%s]", dbUrl)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	log.Println("booyah")
	api.GET("/ping", pingFunc(db))

	r.Run()
}
