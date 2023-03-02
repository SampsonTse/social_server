package utils

import (
	"context"
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/session"
)

func initSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "sessionId", // 客户端存储 cookie 的名字
		EnableSetCookie: true,        // 是否开启 SetCookie,omitempty 这个设置
		Gclifetime:      3600 * 24,   // 触发 GC 的时间
		Maxlifetime:     3600 * 24,   // 服务器端存储的数据的过期时间
		Secure:          false,       // 是否开启 HTTPS，在 cookie 设置的时候有 cookie.Secure 设置
		CookieLifeTime:  3600 * 24,   // 客户端存储的 cookie 的时间，默认值是 0，即浏览器生命周期
		ProviderConfig:  "./tmp",     // 配置信息，根据不同的引擎设置不同的配置信息
	}
	globalSessions, _ := session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

func CheckSess(sess session.Store, key string, account string) bool {

	value, _ := sess.Get(context.Background(), key).(string)

	if value == account {
		return true
	} else {
		fmt.Println("TOKEN not belong this user")
		logs.Info("WARN: TOKEN not belong this user:", account)
		return false
	}
}
