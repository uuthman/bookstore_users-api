package users_db

import (
	"os"
	"log"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //accessing driver
)

const(
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_user_schema = "mysql_user_schema"
)


var (
	Client *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_user_schema)
)

func init(){
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",username,password,host,schema)

	var err error
	Client,err = sql.Open("mysql",dataSourceName)

	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil{
		panic(err)
	}

	log.Println("database successfully configured")
}