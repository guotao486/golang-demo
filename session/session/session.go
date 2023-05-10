package session

// session 存储接口
type Provider interface {
	SessionInit(sid string) (Session, error) // 初始化
	SessionRead(sid string) (Session, error) // 返回session变量，不存在则初始化并返回
	SessionDestroy(sid string) error         // 销毁
	SessionGC(maxLifeTime int64)             // 根据maxLifeTime 删除过期数据
}

// session 处理接口
type Session interface {
	Set(key string, v interface{}) error // set session value
	Get(key string) interface{}          // get session value
	Delete(key string) error             // delete session value
	SessionID() string                   // back current session
}
