/*
 * @Author: GG
 * @Date: 2023-05-13 10:43:11
 * @LastEditTime: 2023-05-16 16:06:07
 * @LastEditors: GG
 * @Description:
 * @FilePath: \task\task\provider\mysql\dao.go
 *
 */
package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"task/task"
	"time"
)

type TaskDao struct {
	db        *sql.DB
	tableName string
}

func (s *TaskDao) SetDB(db *sql.DB) {
	s.db = db
}

func (s *TaskDao) SetTableName(tableName string) {
	s.tableName = tableName
}

func (s *TaskDao) Get() *task.TaskStore {
	taskS := &task.TaskStore{}

	sql := fmt.Sprintf("select * from %s where execution_time <= ? and state = ? limit 1", s.tableName)
	err := s.db.QueryRow(sql, time.Now().Unix(), task.PROCESSED).Scan(&taskS.ID, &taskS.Title, &taskS.Group, &taskS.Rules, &taskS.Data, &taskS.Type, &taskS.State, &taskS.CreateTime, &taskS.ExecutionTime, &taskS.UpdateTime)
	if err != nil {
		return nil
	}

	return taskS
}

func (s *TaskDao) Create(taskS *task.TaskStore) error {

	sql := fmt.Sprintf("insert into %s (title, group_title, rules, data,type,state,create_time,execution_time,update_time) values(?,?,?,?,?,?,?,?,?)", s.tableName)
	ret, err := s.db.Exec(sql, taskS.Title, taskS.Group, taskS.Rules, taskS.Data, taskS.Type, taskS.State, taskS.CreateTime, taskS.ExecutionTime, taskS.UpdateTime)
	if err != nil {
		return err
	}

	_, err = ret.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskDao) UpdateRunState(taskS *task.TaskStore) error {
	sql := fmt.Sprintf("update %s set state = ?, update_time = ? where id = ? and state = ?", s.tableName)
	ret, err := s.db.Exec(sql, task.RUNNING, time.Now().Unix(), taskS.ID, task.PROCESSED)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskDao) UpdateFinishState(taskS *task.TaskStore) error {
	sql := fmt.Sprintf("update %s set state = ?, update_time = ? where id = ? and state = ?", s.tableName)
	ret, err := s.db.Exec(sql, task.FINISH, time.Now().Unix(), taskS.ID, task.RUNNING)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskDao) UpdateIntervalFinishState(taskS *task.TaskStore) error {
	sql := fmt.Sprintf("update %s set state = ?, update_time = ?, execution_time = ? where id = ? and state = ?", s.tableName)
	rules := make(map[string]interface{})
	json.Unmarshal([]byte(taskS.Rules), &rules)
	duration, ok := rules["duration"]
	if !ok {
		return errors.New("rules duration is nil")
	}

	switch duration.(type) {
	case float64:
		duration = time.Duration(duration.(float64))
	case int64:
		duration = time.Duration(duration.(float64))
	}

	ret, err := s.db.Exec(sql, task.PROCESSED, time.Now().Unix(), time.Now().Add(duration.(time.Duration)).Unix(), taskS.ID, task.RUNNING)
	fmt.Printf("err: %v\n", err)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskDao) UpdateFailState(taskS *task.TaskStore) error {
	sql := fmt.Sprintf("update %s set (state = ?, update_time = ?) where id = ? and state = ?", s.tableName)
	ret, err := s.db.Exec(sql, task.RUNNING, time.Now().Unix(), taskS.ID, taskS.State)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
