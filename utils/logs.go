package utils

import (
	"encoding/json"
	"strings"

	"github.com/beego/beego/config"
	"github.com/beego/beego/v2/core/logs"
)

/**
@LastEditDate: 2023/2/28
@description: 日志生成
**/

func initLogger() {
	conf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	logMode := conf.DefaultString("logmode", "console")
	logLevel := getLogLevel(conf.DefaultString("logLevel", "INFO"))

	if logMode == "file" {
		logFile := conf.DefaultString("logfilepath", "logs/file.log")
		logConfig := &struct {
			FileName string `json:"fileName"` // 日志文件名
			MaxLines int    `json:"maxLines"` // 每个文件保存的最大行数，默认值 1000000
			MaxSize  int    `json:"maxSize"`  // 每个文件保存的最大尺寸，默认值是 1 << 28, 256 MB
			Daily    bool   `json:"daily"`    // 是否按照每天 logRotate，默认是 true
			MaxDays  int64  `json:"maxDays"`  // 文件最多保存多少天，默认保存 7 天
			Rotate   bool   `json:"rotate"`   // 是否开启 logRotate，默认是 true
			Perm     string `json:"perm"`     // 日志文件权限
		}{
			FileName: logFile,
			MaxLines: 100000,
			MaxSize:  1024000,
			Daily:    true,
			MaxDays:  30,
			Rotate:   true,
			Perm:     "0600",
		}

		logBuf, err := json.Marshal(logConfig)
		if err != nil {
			panic(err)
		}

		if err := logs.SetLogger(logs.AdapterFile, string(logBuf)); err != nil {
			panic(err)
		}
	}

	logs.EnableFuncCallDepth(true)
	logs.SetLevel(logLevel)
}

func getLogLevel(level string) int {
	switch strings.ToLower(level) {
	case "emergency":
		return logs.LevelEmergency
	case "alert":
		return logs.LevelAlert
	case "critical":
		return logs.LevelCritical
	case "error":
		return logs.LevelError
	case "warning":
		return logs.LevelWarning
	case "notice":
		return logs.LevelNotice
	case "info":
		return logs.LevelInfo
	case "debug":
		return logs.LevelDebug
	default:
		panic("错误logLevel级别")
	}
}
