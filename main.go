package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := establishDB()
	if err != nil {
		log.Fatalf("Unable to create DB; Error message: %v", err)
	}

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

		_, err = db.Exec("CREATE TABLE `customers` (`till_id` INTEGER PRIMARY KEY AUTOINCREMENT, `client_id` VARCHAR(64) NULL, `first_name` VARCHAR(255) NOT NULL, `last_name` VARCHAR(255) NOT NULL, `guid` VARCHAR(255) NULL, `dob` DATETIME NULL, `type` VARCHAR(1))")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	} else {
		return err
	}
	return nil
}
