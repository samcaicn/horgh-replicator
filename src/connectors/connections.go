package connectors

import (
	"database/sql"
	"fmt"
	"github.com/siddontang/go-log/log"
	"go-binlog-replication/src/constants"
	"go-binlog-replication/src/helpers"
	"strconv"
	"time"
)

type Connection interface {
	Ping() bool
	Exec(params map[string]interface{}) bool
}

type ConnectionReplicator interface {
	Connection
	Get(map[string]interface{}) *sql.Rows
}

type ConnectionPool struct {
	master     Connection // used only for loader
	slave      Connection
	replicator ConnectionReplicator
}

var connectionPool ConnectionPool

var retryCounter = map[string]int{
	constants.DBReplicator: 0,
	constants.DBSlave:      0,
	constants.DBMaster:     0,
}

func Exec(mode string, params map[string]interface{}) bool {
	switch mode {
	case constants.DBMaster:
		connectionPool.master = GetMysqlConnection(connectionPool.master, constants.DBMaster).(Connection)
		return connectionPool.master.Exec(params)
	case constants.DBReplicator:
		connectionPool.replicator = GetMysqlConnection(connectionPool.replicator, constants.DBReplicator).(ConnectionReplicator)
		return connectionPool.replicator.Exec(params)
	// adapters for slave databases
	case "mysql":
		connectionPool.slave = GetMysqlConnection(connectionPool.slave, constants.DBSlave).(Connection)
		return connectionPool.slave.Exec(params)
	case "clickhouse":
		connectionPool.slave = GetClickhouseConnection(connectionPool.slave, constants.DBSlave).(Connection)
		return connectionPool.slave.Exec(params)
	case "postgresql":
		connectionPool.slave = GetPostgresqlConnection(connectionPool.slave, constants.DBSlave).(Connection)
		return connectionPool.slave.Exec(params)
	}

	return false
}

func Get(params map[string]interface{}) *sql.Rows {
	connectionPool.replicator = GetMysqlConnection(connectionPool.replicator, constants.DBReplicator).(ConnectionReplicator)
	return connectionPool.replicator.Get(params)
}

func Retry(storageType string, cred helpers.Credentials, connection Connection, method func(connection Connection, dbName string) interface{}) interface{} {
	if retryCounter[storageType] > cred.RetryAttempts {
		log.Fatal(fmt.Sprintf(constants.ErrorDBConnect, storageType))
	}

	log.Infof(constants.MessageRetryConnect, storageType, strconv.Itoa(cred.RetryTimeout))

	time.Sleep(time.Duration(cred.RetryTimeout) * time.Second)
	retryCounter[storageType]++

	return method(connection, storageType)
}