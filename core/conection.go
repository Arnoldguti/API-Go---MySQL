package core

import (
	_ "github.com/go-sql-driver/mysql"


)



var (
	SERVER_HOST        string
	SERVER_USER        string
	SERVER_PORT        string
	SERVER_PASS        string
	SERVER_DATABASE    string
)




func InitConfig() (string){

	SERVER_HOST = "IP"
	SERVER_USER = "USER"
	SERVER_PORT = "PORT"
	SERVER_PASS = "PASSWORD"
	SERVER_DATABASE = "DATABASE"


	return  SERVER_USER + ":" + SERVER_PASS + "@tcp(" + SERVER_HOST + ":" + SERVER_PORT + ")/" + SERVER_DATABASE


}






