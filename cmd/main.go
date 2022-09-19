package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/AlekseySauron/tomato/app"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Ошибка viper", err.Error())
		log.Fatal("Ошибка viper", err)
		return
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	application := app.NewApplication()
	go application.Run()
	<-signalChan
	application.Stop()
}
