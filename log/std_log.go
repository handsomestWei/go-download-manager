package log

import (
	"fmt"
	"time"
)

var appName = "upgrade"
// 日志格式：[应用名] [日志级别] [时间] [日志内容]
var stdFormat = "[%s] [%s] [%s] %s"

func Infof(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(stdFormat, appName, "INFO", now(), printf(format, a, len(a))))
}

func Errorf(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(stdFormat, appName, "ERROR", now(), printf(format, a, len(a))))
}

func Error(a interface{}) {
	fmt.Println(fmt.Sprintf(stdFormat, appName, "ERROR", now(), a))
}

func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TODO
func printf(format string, a []interface{}, len int) string {
	switch len {
	case 1:
		return fmt.Sprintf(format, a[0])
	case 2:
		return fmt.Sprintf(format, a[0], a[1])
	case 3:
		return fmt.Sprintf(format, a[0], a[1], a[2])
	case 4:
		return fmt.Sprintf(format, a[0], a[1], a[2], a[3])
	default:
		return fmt.Sprintf(format, a)
	}
}
