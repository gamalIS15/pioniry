package main

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"pioniry/db"
	"pioniry/routes"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error while reading config", err)
	}
}
func main() {

	//DB Connection
	dbs := db.DatabaseConn.GetDB(db.DatabaseConn{})
	con, _ := dbs.DB()
	defer func(con *sql.DB) {
		err := con.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(con)

	//Running Server
	echo := routes.Routing.GetRoutes(routes.Routing{})
	echo.Logger.Fatal(echo.Start(viper.GetString("HOSTNAME") + ":" + viper.GetString("PORT")))

}
