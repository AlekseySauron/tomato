package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AlekseySauron/tomato/pkg/dbpkg"
	"github.com/AlekseySauron/tomato/pkg/delivery/telegrampkg"
	"github.com/spf13/viper"
)

type Application struct {
	repo dbpkg.DBRepository
}

func NewApplication() *Application {
	return &Application{}

}

func (a *Application) Run() {
	a.repo = a.getDB()

	//service := dbpkg.NewService(a.repo)

	telegrampkg.Start(a.repo)

}

func (a *Application) getDB() dbpkg.DBRepository {
	db, err := sql.Open("mysql", viper.GetString("db.connect"))
	if err != nil {
		log.Fatalln(err)
	}
	return dbpkg.NewSQLRepository(db)
}

func (a *Application) Stop() {
	fmt.Println("Stopping...")
	a.repo.Close()
}
