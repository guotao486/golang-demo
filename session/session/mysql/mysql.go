package mysql

import (
	"database/sql"
	"fmt"
	"golang-demo/session/session"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	TableName string
	Username  string
	Password  string
	Host      string
	Port      uint
	Database  string
	DB        *sql.DB
}

var pder = &Provider{sessions: make([]*SessionStore, 0)}
var dao = &SessionDao{}

type SessionStore struct {
	sid          string                 // session id 唯一标识
	timeAccessed time.Time              // 最后访问时间
	value        map[string]interface{} // session存储的值
}

func (st *SessionStore) Set(key string, value interface{}) error {
	st.value[key] = value
	dao.Set(st.sid, st.value)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key string) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		if key == "createTime" {
			switch v.(type) {
			case float64:
				return int64(v.(float64))
			case int64:
				return v.(int64)
			}
		}
		return v
	} else {
		return nil
	}
}

func (st *SessionStore) Delete(key string) error {
	delete(st.value, key)
	dao.Set(st.sid, st.value)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

func (st *SessionStore) GetValue() map[string]interface{} {
	pder.SessionUpdate(st.sid)
	return st.value
}

func (st *SessionStore) SetValue(v map[string]interface{}) {
	st.value = v
	dao.Set(st.sid, st.value)
	pder.SessionUpdate(st.sid)
}

type Provider struct {
	lock     sync.Mutex // 锁
	sessions []*SessionStore
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[string]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	dao.Create(newsess)
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	session := dao.Get(sid)
	if session != nil {
		element := &SessionStore{
			sid:          session.SessionID,
			value:        session.Value,
			timeAccessed: time.Unix(session.AccessedTime, 0),
		}
		return element, nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
}

func (pder *Provider) SessionDestroy(sid string) error {
	session := dao.Get(sid)
	if session != nil {
		dao.Delete(sid)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGC(maxLifeTime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		elements, _ := dao.GetMaxlife(maxLifeTime)

		if elements == nil {
			break
		}

		dao.DeleteBatch(elements)
		break
	}
}

func (pder *Provider) SessionRefurbish(sid string, newsid string) (session.Session, error) {
	pder.SessionUpdate(sid)
	pder.lock.Lock()
	e := dao.UpdateSessionID(sid, newsid)
	pder.lock.Unlock()
	if e == nil {
		return pder.SessionRead(newsid)
	}
	return nil, nil
}

// 更新最后访问时间
func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	session := dao.Get(sid)
	if session != nil {
		dao.UpdateAccessedTime(sid)
		return nil
	}
	return nil
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
		`id INT(10) NOT NULL AUTO_INCREMENT,` +
		`session_id VARCHAR(255) NULL DEFAULT NULL,` +
		`value VARCHAR(255) NULL DEFAULT NULL,` +
		`create_time int NULL DEFAULT 0,` +
		`accessed_time int NULL DEFAULT 0,` +
		`PRIMARY KEY (id),` +
		`UNIQUE INDEX(session_id)` +
		`)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	smt, err := db.Prepare(tablesql)
	fmt.Printf("\n %s \n", tablesql)
	smt.Exec()

	dao.SetDB(db)
	dao.SetTableName(config.TableName)
}

func init() {
	session.Register("mysql", pder)
}
