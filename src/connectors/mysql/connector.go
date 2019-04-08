package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/siddontang/go-log/log"
	"go-binlog-replication/src/constants"
	"go-binlog-replication/src/helpers"
	"strconv"
)

const DSN = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

type connect struct {
	base *sql.DB
}

func (conn connect) Ping() bool {
	if conn.base.Ping() == nil {
		return true
	}

	return false
}

func (conn connect) Exec(params map[string]interface{}) bool {
	_, err := conn.base.Exec(fmt.Sprintf("%v", params["query"]), helpers.MakeSlice(params["params"])...)
	if err != nil {
		log.Warnf(constants.ErrorExecQuery, "mysql", err)
		return false
	}

	return true
}

func (conn connect) Get(params map[string]interface{}) *sql.Rows {
	rows, err := conn.base.Query(fmt.Sprintf("%v", params["query"]), helpers.MakeSlice(params["params"])...)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func GetConnection(connection helpers.Storage, storageType string) interface{} {
	if connection == nil || connection.Ping() == false {
		cred := helpers.GetCredentials(storageType).(helpers.CredentialsDB)
		conn, err := sql.Open("mysql", buildDSN(cred))
		if err != nil || conn.Ping() != nil {
			connection = helpers.Retry(storageType, cred.Credentials, connection, GetConnection).(helpers.Storage)
		} else {
			connection = connect{conn}
		}
	}

	return connection
}

func buildDSN(cred helpers.CredentialsDB) string {
	return fmt.Sprintf(DSN, cred.User, cred.Pass, cred.Host, strconv.Itoa(cred.Port), cred.DBname)
}