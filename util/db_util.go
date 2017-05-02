package util

import (
	"time"
	"gopkg.in/pg.v3"
	"fmt"
)

var Db *pg.DB

func Init() {
	config := InitConfig()
	Db = singleConnect(config)
}

func singleConnect(config *ServerConfig) *pg.DB {
	return pg.Connect(pgOptions(config.Host[0], config.Port, config.User, config.Password, config.Database))
}

var Index = 32
//var Dbs [Index] *pg.DB
var Dbs []*pg.DB

func InitDbs() {
	config := InitConfig()
	Dbs = make([]*pg.DB, 32)
	connect(config, Dbs)
}

func connect(config *ServerConfig, dbs []*pg.DB) {
	for i, host := range config.Host{
		dbs[i] = pg.Connect(pgOptions(host, config.Port, config.User, config.Password, config.Database))
	}
}

func pgOptions(host string, port string, user string, password string, database string) *pg.Options {
	return &pg.Options{
		Host: host,
		Port: port,
		User:     user,
		Password: password,
		Database: database,

		DialTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        10 * time.Second,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
}

func CaculateDbAndTableIndex(userId int) (int, int) {
	last4digit := userId % 10000

	//for product enviroment
	//dbIndex := last4digit % Index
	//tableIndex := last4digit / Index % Index

	//for local test enviroment
	dbIndex := last4digit % 1
	tableIndex := last4digit / 1 % 2

	return dbIndex, tableIndex
}

func CaculateDbAndTable(userId int) (*pg.DB, string) {
	dbIndex, tableIndex := CaculateDbAndTableIndex(userId)

	return Dbs[dbIndex], fmt.Sprintf("%s_%d", "relation", tableIndex)
}

