package main

import (
    "fmt"
    "net/http"
    "time"
    "sync"
)

type Manager struct {
    cookieName string
    lock sync.Mutex
    provider Provider
    maxLifeTime int64
}

type Provider interface {
    SessionInit(sid string) (Session, error)
    SessionRead(sid string) (Session, error)
    SessionDestroy(sid string) error
    SessionGC(maxlifetime int64)
}

type Session interface {
    Set(key, value interface{}) error // set session value
    Get(key interface{}) interface{}  // get session value
    Delete(key interface{}) error   // delete session value
    SessionID() string              // back current sessionID
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
    provider, ok := provides[provideName]
    if !ok {
        return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
    }
    return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxlifetime}, nil
}

var globalSessions *session.Manager

func init() {
    globalSessions, _ = NewManager("memory", "gosessionid", 3600)
}

func main() {
    // Cookie is struct
    // set cookie
    // http.SetCookie(w ResponseWriter, cookie *Cookie)
    /*
        expiration := time.Now()
        expiration = expiration.AddDate(1, 0, 0)
        cookie := http.Cookie{Name: "username", Value: "zzh", Expires: expiration}
        http.SetCookie(w, &cookie)
    */

    // read cookie
    /*
        cookie, _ := r.Cookie("username")
        fmt.Fprintf(w, cookie)
    */
    /*
        for _, cookie := range r.Cookie() {
            fmt.Fprintf(w, cookie.Name)
        }
    */
    // session创建过程
    /*
        生成全局唯一标识符
        开辟数据存储空间
        将session的全局唯一标识符发送给客户端
        // cookie和URL重写
        // 1.服务端通过设置Set-cookie头将session的标识符传送到客户端
        // 客户端每次请求都会带上这个标识符
        // 2.URL重写就是在返回给用户页面的所有URL后面追加session标识符
    */
}
