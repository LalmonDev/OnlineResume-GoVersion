/**
 * @Author: lalmon
 * @Description: 日志初始化
 * @File:  log.go
 * @Date: 2021/8/2 5:50 下午
 */

package init_conf

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

const (
	defaultLogPath  = "/opt/logs"
	defaultAppName  = "OnlineResume-GoVersion"
	defaultLogDay   = 30
	defaultLogLevel = 7
)

func LogInit() bool {
	path := beego.AppConfig.DefaultString("LogPath", defaultLogPath)
	appName := beego.AppConfig.DefaultString("appname", defaultAppName)
	logName := fmt.Sprintf("%s/%s.log", path, appName)
	logFormat := fmt.Sprintf("{\"filename\":\"%s\",\"daily\":true,\"maxdays\":%d,\"color\":true}", logName, defaultLogDay)
	logLeavel := beego.AppConfig.DefaultInt("DebugLevel", defaultLogLevel)

	// 设置日志输出方式和位置
	err := logs.SetLogger(logs.AdapterFile, logFormat)
	if nil != err {
		logs.Error(err)
		return false
	}

	// 设置日志输出级别
	logs.SetLevel(logLeavel)

	//输出文件名和文件行号，beego默认不输出
	logs.EnableFuncCallDepth(true)

	return true
}
