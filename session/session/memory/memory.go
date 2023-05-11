/*
 * @Author: GG
 * @Date: 2023-05-06 10:26:42
 * @LastEditTime: 2023-05-11 10:33:20
 * @LastEditors: GG
 * @Description:
 * @FilePath: \session\session\memory\memory.go
 *
 */
package memory

import (
	"container/list"
	"golang-demo/session/session"
	"sync"
	"time"
)

var pder = &Provider{list: list.New()}

type SessionStore struct {
	sid          string                 // session id 唯一标识
	timeAccessed time.Time              // 最后访问时间
	value        map[string]interface{} // session存储的值
}

func (st *SessionStore) Set(key string, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key string) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (st *SessionStore) Delete(key string) error {
	delete(st.value, key)
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
	pder.SessionUpdate(st.sid)
}

type Provider struct {
	lock     sync.Mutex               // 锁
	sessions map[string]*list.Element // 存储在内存
	list     *list.List               // 用来做GC
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[string]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGC(maxLifeTime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}

}

func (pder *Provider) SessionRefurbish(sid string, newsid string) (session.Session, error) {
	olse, _ := pder.SessionRead(sid)
	nese, _ := pder.SessionInit(newsid)
	nese.SetValue(olse.GetValue())
	pder.SessionDestroy(sid)
	pder.SessionUpdate(newsid)
	return nese, nil
}

// 更新最后访问时间
func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

// 初始化，注册session管理器
func init() {
	pder.sessions = make(map[string]*list.Element)
	session.Register("memory", pder)
}
