package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Session struct {
	ID           uint64
	SessionID    string
	AccessedTime int64 // 最后访问时间
	Value        map[string]interface{}
	CreateTime   int64
}

type SessionDao struct {
	db        *sql.DB
	tableName string
}

func (s *SessionDao) SetDB(db *sql.DB) {
	s.db = db
}

func (s *SessionDao) SetTableName(tableName string) {
	s.tableName = tableName
}

func (s *SessionDao) Create(session *SessionStore) error {
	sql := fmt.Sprintf("insert into %s(session_id, value, create_time, accessed_time) values(?,?,?,?)", s.tableName)
	valueStr, _ := json.Marshal(session.value)
	fmt.Printf("valueStr: %v\n", valueStr)
	fmt.Printf("session.value: %v\n", session.value)
	ret, err := s.db.Exec(sql, session.sid, string(valueStr), time.Now().Unix(), time.Now().Unix())
	if err != nil {
		return err
	}

	_, err = ret.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// 获取
func (s *SessionDao) Get(sid string) *Session {
	var session Session
	var valueStr string
	sql := fmt.Sprintf("select * from %s where session_id = ? limit 1", s.tableName)
	err := s.db.QueryRow(sql, sid).Scan(&session.ID, &session.SessionID, &valueStr, &session.CreateTime, &session.AccessedTime)

	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(valueStr), &session.Value)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return &session
}

// 存储/更新
func (s *SessionDao) Set(sid string, value map[string]interface{}) error {
	var valueStr []byte
	valueStr, err := json.Marshal(value)
	fmt.Printf("valueStr: %v\n", valueStr)
	fmt.Printf("value: %v\n", value)
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("update %s set value = ? where session_id = ?", s.tableName)
	ret, err := s.db.Exec(sql, string(valueStr), sid)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionDao) Delete(sid string) error {
	sql := fmt.Sprintf("delete from %s where session_id = ?", s.tableName)
	ret, err := s.db.Exec(sql, sid)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionDao) UpdateAccessedTime(sid string) error {
	sql := fmt.Sprintf("update %s set value = ? where accessed_time = ?", s.tableName)
	ret, err := s.db.Exec(sql, time.Now().Unix(), sid)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// 获取已失效session
func (s *SessionDao) GetMaxlife(maxlifeTime int64) ([]string, error) {
	sql := fmt.Sprintf("select session_id from %s where accessed_time < ? limit 10", s.tableName)

	rows, err := s.db.Query(sql, time.Now().Unix()-maxlifeTime)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var sessionIDs = make([]string, 0)

	for rows.Next() {
		var sessionID string
		err := rows.Scan(&sessionID)
		if err != nil {
			return nil, err
		}

		sessionIDs = append(sessionIDs, sessionID)
	}

	return sessionIDs, nil
}

// 批量删除session
func (s *SessionDao) DeleteBatch(sessionIDs []string) error {
	sql := fmt.Sprintf("delete from %s where session_id in (?)", s.tableName)
	sessionStr := strings.Join(sessionIDs, ",")
	ret, err := s.db.Exec(sql, sessionStr)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// 更改session_id
func (s *SessionDao) UpdateSessionID(sid, newSid string) error {
	sql := fmt.Sprintf("update %s set session_id = ? where session_id = ?", s.tableName)
	ret, err := s.db.Exec(sql, sid, newSid)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
