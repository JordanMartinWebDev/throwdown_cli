package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := establishDB()
	if err != nil {
		log.Fatalf("Unable to create DB; Error message: %v", err)
	}

	db, err := sql.Open("sqlite3", "./data/throwdown.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

func establishDB() error {
	if _, err := os.Stat("./data/throwdown.db"); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll("./data", 0755)
		os.Create("./data/throwdown.db")

		db, err := sql.Open("sqlite3", "./data/throwdown.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		_, err = db.Exec("CREATE TABLE `tags` (`tag_id` INTEGER PRIMARY KEY AUTOINCREMENT, `name` VARCHAR(255) NOT NULL, `description` VARCHAR(255) NOT NULL)")
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("CREATE TABLE `ideas` (`idea_id` INTEGER PRIMARY KEY AUTOINCREMENT, `name` VARCHAR(255) NOT NULL, `description` VARCHAR(255) NOT NULL)")
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("CREATE TABLE `tag_to_idea` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `tag_id` INTEGER REFERENCES tags(tag_id), `idea_id` INTERGER REFERENCES ideas(`idea_id`))")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return err
	}
	return nil
}
