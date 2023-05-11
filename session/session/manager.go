// session 管理器
package session

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// session 管理器
type Manager struct {
	cookieName  string     // private cookie name
	lock        sync.Mutex // protects session
	provider    Provider
	maxLifeTime int64
	refurbish   int64
}

// 存储引擎桶
var provides = make(map[string]Provider)

// 注册存储session
// Register通过提供的名称使会话提供可用。
// if 用相同的名称调用Register两次，or if driver is nil，
// it panic
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

// 初始化管理器
//
// @param string provider 供应者
//
// @param string cookieName cookie名称
//
// @param int64 maxLifeTime 最大寿命
//
// @param int64 refurbish 刷新时间 0 不刷新
func NewManager(provideName, cookieName string, maxLifeTime int64, refurbish int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime, refurbish: refurbish}, nil
}

// 生成全局唯一session id
func (manager *Manager) sessionId() string {
	guid := GetGUID().Hex()
	return base64.URLEncoding.EncodeToString([]byte(guid))

}

// 创建
func (manager *Manager) sessionCreate(w http.ResponseWriter) (session Session) {
	sid := manager.sessionId()
	fmt.Printf("sid: %v\n", sid)
	session, _ = manager.provider.SessionInit(sid)
	fmt.Println("session init createTime")
	session.Set("createTime", time.Now().Unix())
	cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
	http.SetCookie(w, &cookie)
	return
}

// 刷新
func (manager *Manager) sessionRefurbish(sid string, w http.ResponseWriter) (session Session) {
	newsid := manager.sessionId()
	session, _ = manager.provider.SessionRefurbish(sid, newsid)
	session.Set("createTime", time.Now().Unix())
	cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(newsid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
	http.SetCookie(w, &cookie)
	return
}

// 检测用户是否与session关联，没有则创建
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	cookie, err := r.Cookie(manager.cookieName)

	if err != nil || cookie.Value == "" { // 创建 session
		session = manager.sessionCreate(w)
	} else { // 获取关联session，更新session createTime,若createTime 失效则销毁重建
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)

		if manager.refurbish > 0 {
			// 间隔刷新新的session id 防止session劫持
			createTime := session.Get("createTime")

			if createTime == nil {
				session.Set("createTime", time.Now().Unix())
			} else if (createTime.(int64) + manager.refurbish) < (time.Now().Unix()) {
				session = manager.sessionRefurbish(sid, w)
			}
		}
	}
	return
}

// Sessiion 销毁
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		sid, _ := url.QueryUnescape(cookie.Value)
		manager.provider.SessionDestroy(sid)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

// 自动销毁过期session
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() {
		manager.GC()
	})
}
