package log4j

import (
	"regexp"
	"runtime"
	"strings"
	"time"
)


const(
	YYYYMMDDSS00 = "2006-01-02 15:04:05.00"
)

var timex = time.Now()

// @param return 文件名 文件行号
func GetCalle() (file string, line int){
	var ok bool
	_, file, line,ok = runtime.Caller(4)
	if !ok{
		file = "???"
	}
	return
}

// @param 格式化时间 2021-07-29 23:36::11.22
func GetYYYYMMDDSS00() string {
	return timex.Format(YYYYMMDDSS00)
}

// @param 当前协程ID
func GetGoroutines() (no int){
	var b = make([]byte, 1<<5)
	b = b[:runtime.Stack(b, true)]
	reg, err := regexp.Compile(`\d`)
	if err != nil {
		no = -1
	}
	if !strings.Contains(string(b), "goroutine"){
		no = -1
	}
	return int(reg.Find(b)[0])
}