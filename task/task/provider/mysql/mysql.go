/*
 * @Author: GG
 * @Date: 2023-05-13 10:24:02
 * @LastEditTime: 2023-05-16 15:34:53
 * @LastEditors: GG
 * @Description:
 * @FilePath: \task\task\provider\mysql\mysql.go
 *
 */
package mysql

import (
	"database/sql"
	"fmt"
	"task/task"

	_ "github.com/go-sql-driver/mysql"
)

var dao = &TaskDao{}
var pder = &Provider{dao: dao}

type MysqlConfig struct {
	TableName string
	Username  string
	Password  string
	Host      string
	Port      uint
	Database  string
	DB        *sql.DB
}

type Provider struct {
	dao *TaskDao
}

func (pder *Provider) Add(taskS *task.TaskStore) error {
	pder.dao.Create(taskS)
	return nil
}

func (pder *Provider) Get() *task.TaskStore {
	return pder.dao.Get()
}

func (pder *Provider) Run(taskS *task.TaskStore) {
	pder.dao.UpdateRunState(taskS)
}

func (pder *Provider) Finish(taskS *task.TaskStore) {
	if taskS.Type == task.IntervalTask {
		pder.dao.UpdateIntervalFinishState(taskS)
	} else {
		pder.dao.UpdateFinishState(taskS)
	}
}

func (pder *Provider) Fail(taskS *task.TaskStore) {
	pder.dao.UpdateFailState(taskS)
}

func InitDB(config MysqlConfig) {
	var db *sql.DB
	if config.DB == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", config.Username, config.Password, config.Host, config.Port, config.Database)
		mysqldb, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		db = mysqldb
	} else {
		db = config.DB
	}

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	tablesql := `CREATE TABLE IF NOT EXISTS ` + config.TableName + ` ( ` +
		`id INT NOT NULL AUTO_INCREMENT,` +
		`title VARCHAR(255) NULL DEFAULT NULL,` +
		`group_title VARCHAR(255) NULL DEFAULT NULL,` +
		`rules VARCHAR(255) NULL DEFAULT NULL,` +
		`data VARCHAR(255) NULL DEFAULT NULL,` +
		`type int(4) null default 0,` +
		`state int(4) null default 0,` +
		`create_time bigint NULL DEFAULT 0,` +
		`update_time bigint NULL DEFAULT 0,` +
		`execution_time bigint NULL DEFAULT 0,` +
		`PRIMARY KEY (id)` +
		`)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	smt, err := db.Prepare(tablesql)
	fmt.Printf("\n %s \n", tablesql)
	fmt.Printf("err: %v\n", err)
	smt.Exec()

	dao.SetDB(db)
	dao.SetTableName(config.TableName)
}

func init() {
	task.ProviderRegister("mysql", pder)
}
