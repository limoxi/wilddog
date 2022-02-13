package db

import (
	"github.com/limoxi/ghost"
)

func init() {
	dbConf := ghost.Config.GetMap("database.default")
	ghost.ConnectDB(
		ghost.NewDbConfig(
			dbConf.GetString("engine", "postgres"),
			dbConf.GetString("host"),
			dbConf.GetString("port", "3306"),
			dbConf.GetString("user"),
			dbConf.GetString("password"),
			dbConf.GetString("dbname"),
			dbConf.GetBool("debug", false),
		),
	)
}
