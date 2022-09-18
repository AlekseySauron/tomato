package dbpkg

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// //import "github.com/AlekseySauron/tomato/pkg/dbpkg"

// type service struct {
// 	dbrepo DBRepository
// }

// func NewService(repo DBRepository) service {
// 	return &service{
// 		dbrepo: repo,
// 	}
// }

func insertTask(db *sql.DB) {
	curTime := time.Now()
	result, err := db.Exec("insert into tomato.Tasks (user, status, time_start, time_control) values (?, ?, ?, ?)",
		"user2", "begin", curTime, curTime.Add(5*time.Minute))

	if err != nil {
		log.Panicln("Ошибка Insert", err)
	}
	fmt.Println(result.RowsAffected())
}

func printDbCount(db *sql.DB) {
	var count int
	rows, err := db.Query("select count(*) as count from tomato.Tasks")
	if err != nil {
		log.Panicln("Ошибка select", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(count)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
